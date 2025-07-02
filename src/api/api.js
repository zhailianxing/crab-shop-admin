import axiosInstance from "~/axios.js";

/* 
    这里是请求我本地admin目录下的服务器接口

// 登录
export function LoginApi(username, password) {
    // axios.post是返回一个 promise对象. 用.then(res=>{}) 处理返回值,用.catch(err=>{}) 处理错误信息
    return axiosInstance.post("/api/user/login",{
        username,
        password
    })
}


// 获取用户信息
export function getUserInfo() {
    return axiosInstance.post("/api/user/profile")
}

// 退出登录
export function logout() {
    return axiosInstance.post("/api/user/logout")
}

// 修改密码
export function modifyPwd(data = {}) {
    return axiosInstance.post("/api/user/modifyPwd",data)
}

// 获取菜单
export function getMenus() {
    return axiosInstance.get("/api/menus")
}
    */


export function LoginApi(username, password) {
    // axios.post是返回一个 promise对象. 用.then(res=>{}) 处理返回值,用.catch(err=>{}) 处理错误信息
    return axiosInstance.post("/admin/login",{
        username,
        password
    })
}

// 退出登录
export function logout() {
    return axiosInstance.post("/admin/logout")
}

// 修改密码
export function modifyPwd(data = {}) {
    return axiosInstance.post("/admin/updatepassword",data)
}

// 获取菜单
export function getMenus() {
    return axiosInstance.post("/admin/getinfo")
}

// 获取用户信息
export function getUserInfo() {
    return axiosInstance.post("/admin/getinfo")
}

// 后台首页统计1
export function getStatistics1() {
    return axiosInstance.get("/admin/statistics1")
}