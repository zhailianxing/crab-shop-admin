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
