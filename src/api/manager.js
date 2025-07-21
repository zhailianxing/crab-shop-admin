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

/* 
    第三部分： 角色管理功能
*/
// 获取 所有的权限列表
export function getPermissionList() {
    return axiosInstance.get(`/admin/rule/1`)
}

// 获取角色列表
export function getRoleList(page, queryData = {limit: 10, keyword: ""}) {
    const encoded = new URLSearchParams(queryData).toString();
    return axiosInstance.get(`/admin/role/${page}` + "?" + encoded)
}
// 新增角色基本信息
export function addRole(data) {
    return axiosInstance.post(`/admin/role`, data)
}

// 修改角色基本信息
export function modifyRole(id,data) {
    return axiosInstance.post(`/admin/role/${id}`, data)
}

// 删除角色
export function deleteRole(id) {
    return axiosInstance.post(`/admin/role/${id}/delete`)
}

// 更改 此角色 状态 (启用/禁用)
export function changeRoleStatus(id, status) {
    return axiosInstance.post(`/admin/role/${id}/update_status`, {status})
}

// 配置角色权限
export function setRolePermission(roleId, permissionIds) {
    return axiosInstance.post(`admin/role/set_rules`, {id: roleId, rule_ids:permissionIds})
}