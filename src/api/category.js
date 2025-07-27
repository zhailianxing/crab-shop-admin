

import axiosInstance from '~/axios.js'


// 获取分类列表
export function getCategoryList() {
    return axiosInstance.get(`admin/category`)
}

// 增加公告
export function addCategory(data = {}) {
    return axiosInstance.post("/admin/category", data)
}

// 增加公告
export function modifyCategory(id, data = {}) {
    return axiosInstance.post(`admin/category/${id}`, data)
}

// 删除公告
export function deleteCategory(id) {
    return axiosInstance.post(`admin/category/${id}/delete`)
}

// 更改优惠卷状态为失效
export function changeCategoryStatus(id) {
    return axiosInstance.post(`/admin/category/${id}/update_status`, {status:0})
}