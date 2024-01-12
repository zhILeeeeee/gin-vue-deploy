import axios from "axios"


const service = axios.create({
    baseURL: '',
    timeout: 180000
})

let isBlob = false
service.interceptors.request.use(
    (config) => {
        const token = getToken() || TOKEN
        if (token) {
            config.headers.Authorization = token
        }
        // return setDefaultData(config)
        isBlob = config.isBlob
        if (isBlob) {
            config.responseType = 'blob'
        }
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

service.interceptors.response.use(
    (response) => {
        const res = response.data
        if (isBlob) {
            return res
        }

        if (res.code === 7025) {
            store.dispatch('user/updateRole', false)
        } else {
            store.dispatch('user/updateRole', true)
        }

        if (
            ![
                200, 6021, 6022, 6023, 6024, 6028, 7025, 8504, 8411, 8401, 8405, 8510,
                9105
            ].includes(res.code)
        ) {
            // 不需要弹报错框
            // const arr = [7028, 7029, 7030, 7031, 7032, 7033, 7034]
            // 7028, '客户名称已存在'
            // 7029, '客户不存在'
            // 7030, '导入客户信息数据为空，请检查数据'
            // 7031, '导入客户模板有误',
            // 7032, '客户名称为空'
            // 7033, '未找到对应的用户'
            // 7034, '客户手机号不为数字'
            if (
                [
                    4050, 4010, 4004, 4003, 4013, 4014, 4015, 4007, 4032, 4021, 4103,
                    4102, 4121, 4104, 4109, 6013
                ].includes(res.code)
                // &&!arr.includes(res.code)
            )
                res.message = ERROR_MESSAGE[res.code]
            ElMessage({
                message: res.message || 'Error',
                type: 'error',
                duration: 1500
            })
            return Promise.reject(new Error(res.message || 'Error'))
        }

        return res
    },
    (error) => {
        ElMessage({
            message: error.message || 'Error',
            type: 'error',
            duration: 1000
        })
        return Promise.reject(error)
    }
)

export default service