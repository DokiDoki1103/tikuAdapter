import axios from 'axios'

const instance = axios.create({
    baseURL: '/adapter-service',
    timeout: 5 * 1000,
    validateStatus: function (status) {
        return status >= 200 && status <= 500
    }
})

// 这里应该拦截 400 或者 500的请求，然后给出用户提示，待完善
instance.interceptors.response.use( (response)=> {
    if (response.status !== 200) {
        console.log(response.data.message) // 应该将这个msg展示给用户
    }
    return response.data
}, function (error) {
    // 给出网络错误提示或者各种请求失败的
    return Promise.reject(error)
})

export async function getQuestions(data) {
    return await instance.post('/questions/search', data)
}

export async function updateQuestions( data) {
    return await instance.put(`/questions/${data.id}`, data)
}

export async function createQuestions(data) {
    return await instance.post(`/questions`, data)
}

export async function delQuestions(id) {
    return await instance.delete(`/questions/${id}`)
}



