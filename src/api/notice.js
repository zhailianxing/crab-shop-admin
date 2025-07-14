import axiosInstance from "~/axios.js";



// 获取公告列表
export function getNoticeList(page) {
    return axiosInstance.get(`admin/notice/${page}`)
}

// 增加公告
export function addNotice(data = {}) {
    return axiosInstance.post("/admin/notice", data)
}

// 增加公告
export function modifyNotice(id, data = {}) {
    return axiosInstance.post(`admin/notice/${id}`, data)
}

// 删除公告
export function deleteNotice(id) {
    return axiosInstance.post(`admin/notice/${id}/delete`)
}

