import axios from 'axios'
import qs from 'qs';

const xmig6Instance = axios.create({
    baseURL: '/sqp/api',
    timeout: 60 * 1000,
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
xmig6Instance.interceptors.response.use(responseInterceptor, errorInterceptor);

export async function parseFile(data) {
    return await xmig6Instance.get(`/parse`, {params: data})
}

export async function reParseFile(data) {
    return await xmig6Instance.post(`/parse`, qs.stringify({html: data.html}))
}


