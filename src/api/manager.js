import axiosInstance from "~/axios.js";


// 获取管理列表
export function getManagerList(page, queryData = {limit: 10, keyword: ""}) {
    const encoded = new URLSearchParams(queryData).toString();
    return axiosInstance.get(`/admin/manager/${page}` + "?" + encoded)
}

// 更改管理员状态
export function changeManagerStatus(id, status) {
    return axiosInstance.post(`/admin/manager/${id}/update_status`, {status})
}

// 新增管理员
export function addManager(data) {
    return axiosInstance.post(`/admin/manager`, data)
}

// 删除管理员
export function delManager(id) {
    return axiosInstance.post(`/admin/manager/${id}/delete`)
}

// 修改管理员信息
export function modifyManager(id,data) {
    return axiosInstance.post(`/admin/manager/${id}`, data)
}

/* 
    第二部分： 权限管理功能
*/
// 新增 菜单或者权限（权限对应实际路由）
export function addManagerMenu(data) {
    return axiosInstance.post(`/admin/rule`, data)
}

// 修改 菜单或者权限
export function modifyManagerMenu(id,data) {
    return axiosInstance.post(`/admin/rule/${id}`, data)
}

// 删除 菜单或者权限
export function deleteManagerMenu(id) {
    return axiosInstance.post(`/admin/rule/${id}/delete`)
}

// 更改 菜单或者权限 状态 (启用/禁用)
export function changeManagerMenuStatus(id, status) {
    return axiosInstance.post(`/admin/rule/${id}/update_status`, {status})
}
