import axios from 'axios'
import qs from 'qs';

const adapterInstance = axios.create({
    baseURL: '/adapter-service',
    timeout: 5 * 1000,
    validateStatus: function (status) {
        return status >= 200 && status <= 500
    }
})
const xmig6Instance = axios.create({
    baseURL: '/sqp/api',
    timeout: 5 * 1000,
    validateStatus: function (status) {
        return status >= 200 && status <= 500
    }
})
// 这里应该拦截 400 或者 500 的请求，然后给出用户提示，待完善
const responseInterceptor = (response) => {
    if (response.status !== 200) {
        console.log(response.data.message); // 应该将这个 msg 展示给用户
    }
    return response.data;
};
const errorInterceptor = (error) => {
    // 给出网络错误提示或者各种请求失败的
    return Promise.reject(error);
};
adapterInstance.interceptors.response.use(responseInterceptor, errorInterceptor);
xmig6Instance.interceptors.response.use(responseInterceptor, errorInterceptor);


export async function getPlat() {
    return await adapterInstance.get('/plat')
}

export async function getCourses(plat) {
    return await adapterInstance.get('/courses', {
        params: {
            plat: plat
        }
    })
}

export async function getQuestions(data) {
    return await adapterInstance.post('/questions/search', data)
}

export async function updateQuestions(data) {
    return await adapterInstance.put(`/questions/${data.id}`, data)
}

export async function createQuestions(data) {
    return await adapterInstance.post(`/questions`, data)
}

export async function delQuestions(id) {
    return await adapterInstance.delete(`/questions/${id}`)
}

export async function parseFile(data) {
    return await xmig6Instance.get(`/parse`, {params: data})
}

export async function reParseFile(data) {
    return await xmig6Instance.post(`/parse`, qs.stringify({html: data.html}))
}


