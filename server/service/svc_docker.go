package service

import (
	"context"
	"errors"
	"fmt"
	"gin-vue-deploy/server/model/request"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/gin-contrib/cache/persistence"
	"io"
	"os"
	"strings"
	"time"
)

type DockerService struct{}

var cache_store = persistence.NewInMemoryStore(time.Minute * 15)

func dockerCli() *client.Client {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return dockerCli
}

func dockerSaveImageAsTar(imageTag string) {
	dockerCli := dockerCli()

	// 执行保存镜像
	saveResponse, err := dockerCli.ImageSave(context.Background(), []string{imageTag})
	if err != nil {
		panic("保存镜像失败")
	}

	// 将保存镜像的响应写入文件
	imageTarPath := fmt.Sprintf("deployPackage/%s.tar", strings.Replace(imageTag, ":", "-", 1))
	file, err := os.Create(imageTarPath)
	_, err = io.Copy(file, saveResponse)
	if err != nil {
		panic(fmt.Sprintf("写入镜像文件失败：%s\n", err))

	}
	saveResponse.Close()
	fmt.Sprintln("镜像%s保存成功！", imageTag)
}

func (dockerService *DockerService) SearchDockerImages(api request.SearchDockerApiModel) (list interface{}, total int64, err error) {

	ctx := context.Background()
	dockerCli := dockerCli()
	var results []registry.SearchResult
	imageTag := fmt.Sprintf("%s:%s", api.Name, api.Tag)

	results, err = dockerCli.ImageSearch(ctx, imageTag, types.ImageSearchOptions{Limit: 15})
	if err != nil {
		panic("镜像搜索失败：%s\n")
	}

	total = int64(len(results))

	return results, total, nil
}

func (dockerService *DockerService) PullDockerImages(api request.PullDockerApiModel) (io.ReadCloser, error) {

	image_name := fmt.Sprintf("%s:%s", api.Name, api.Tag)
	ctx := context.Background()
	dockerCli := dockerCli()
	var results io.ReadCloser
	results, err := dockerCli.ImagePull(ctx, image_name, types.ImagePullOptions{})
	if err != nil {
		panic(fmt.Sprintf("搜索镜像 %s 失败\n", image_name))
	}
	// 先使用内存缓存引擎，设置默认过期时间为 5 分钟(0表示永久)
	// todo: 默认操作是15分钟,需要使用更专业的缓存
	// 先拿取再拼接
	var imageListValue string
	err = cache_store.Get("image_list", &imageListValue)
	if err == persistence.ErrCacheMiss {
		//err = cache_store.Set("image_list", image_name, time.Minute*30)
		err = cache_store.Set("image_list", image_name, 0)
		imageListValue = image_name
		if err != nil {
			panic(fmt.Sprintf("保存镜像 %s 失败\n", image_name))
		}
		return nil, err
	}
	newImageListValue := fmt.Sprintf("%s,%s", imageListValue, image_name)
	err = cache_store.Set("image_list", newImageListValue, 0)

	return results, nil
}

func (dockerService *DockerService) PackDockerImages() error {

	// 查询缓存中的镜像名称
	var imageListValue string

	// test
	err := cache_store.Set("image_list", "mysql:latest,redis:latest", 0)

	err = cache_store.Get("image_list", &imageListValue)
	if err != nil {
		return errors.New(err.Error())
	}
	if imageListValue == "" {
		return errors.New("需要打包的镜像信息不存在")
	}

	imageTagList := strings.Split(imageListValue, ",")
	for _, imageTag := range imageTagList {
		go dockerSaveImageAsTar(imageTag)
	}

	return nil
}
