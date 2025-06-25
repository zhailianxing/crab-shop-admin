import axiosInstance from "~/axios.js";


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
