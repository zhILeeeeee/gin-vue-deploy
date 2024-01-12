import request from '../utils/requests'


/**
 * 查询镜像信息
 * @param {*} data
 * @returns
 */
const queryDockerImageListAPI = (data) =>
    request({
        url: "/docker/queryImages",
        method: "post",
        data
    })