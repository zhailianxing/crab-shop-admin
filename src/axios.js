

import axios from 'axios'
import { showErrorMessage } from '~/common/util.js'
import { getToken } from '~/common/cookie.js'

//axios() 是直接使用默认的全局配置进行请求
//axios.create() 是创建一个基于默认配置的新的自定义 axios 实例。 如果有不同的api请求，就需要用axios.create()
const instance = axios.create({
    // baseURL: 'http://127.0.0.1:8080',
    baseURL: '/whatever', 
    timeout: 3000
})

// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    // 在发送请求之前，将cooke中的token添加到请求头中
    const token = getToken()   
    if (token) {
        // 我自己的后端服务是用的Authorization字段； 别人用的是token字段
        config.headers.Authorization = `Bearer ${token}`
        config.headers.token = `${token}`

    }
    return config;
  }, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  });


// 添加响应拦截器
instance.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response.data;  // 返回服务器的响应数据，业务代码不再用response.data.data了
  }, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    // 对响应错误做点什么： 弹出错误信息
    showErrorMessage(error.response.data.msg || '请求失败')
    return Promise.reject(error);
  });


export default instance