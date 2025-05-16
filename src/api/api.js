import axiosInstance from "~/axios.js";


export function LoginApi(username, password) {
    // axios.post是返回一个 promise对象. 用.then(res=>{}) 处理返回值,用.catch(err=>{}) 处理错误信息
    return axiosInstance.post("/api/user/login",{
        username,
        password
    })
}