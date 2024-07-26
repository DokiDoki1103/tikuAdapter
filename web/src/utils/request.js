import axios from 'axios';
import { getCookies, removeCookies } from './cookies'
import { message } from 'ant-design-vue';

// 创建axios实例
const service = axios.create({
  baseURL: "/",
  timeout: 5 * 1000,
  validateStatus: function (status) {
    return status >= 200 && status <= 500
  }
});
// request拦截器
service.interceptors.request.use(
  (config) => {
    const token = getCookies('token')
    if (token) {
      config.headers['Authorization'] = token; // 让每个请求携带自定义token 请根据实际情况自行修改
    }
    return config;
  },
  (error) => {
    console.log(error, "error-request");
    Promise.reject(error);
  }
);
// 这里应该拦截 400 或者 500 的请求，然后给出用户提示，待完善
const responseInterceptor = (response) => {
  switch (response.status) {
    case 200:
      break;
    case 401:
    case 403:
      // 清除 token 并跳转到登录页
      removeCookies('token')
      message.error('当前登陆过期，请重新登陆。')
      window.location.href = '/login'
      break;
    default:
      console.log(response.data.message); // 将 msg 展示给用户
      break;
  }
  return {
    data: response.data,
    status: response.status,
    message: response.data.message
  };
};
const errorInterceptor = (error) => {
  // 给出网络错误提示或者各种请求失败的
  return Promise.reject(error);
};
service.interceptors.response.use(responseInterceptor, errorInterceptor);

export default service;
