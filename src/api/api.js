import axiosInstance from "~/axios.js";


export function LoginApi(username, password) {
    // axios.post是返回一个 promise对象. 用.then() 获取返回值,用.catch() 获取错误信息
    return axiosInstance.post("/api/login",{
        username,
        password
    })
}