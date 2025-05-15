

import axios from 'axios'


//axios() 是直接使用默认的全局配置进行请求
//axios.create() 是创建一个基于默认配置的新的自定义 axios 实例。 如果有不同的api请求，就需要用axios.create()
const instance = axios.create({
    baseURL: 'http://127.0.0.1:8080',
    timeout: 3000
})


export default instance