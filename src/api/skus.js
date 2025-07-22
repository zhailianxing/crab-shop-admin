import axiosInstance from "~/axios.js";


// 获取规格列表
export function getSkusList(page) {
    return axiosInstance.get(`admin/skus/${page}`)
}

// 新增skus基本信息
export function addSkus(data) {
    return axiosInstance.post(`/admin/skus`, data)
}

// 修改skus基本信息
export function modifySkus(id,data) {
    return axiosInstance.post(`/admin/skus/${id}`, data)
}

// (批量)删除某个规格
export function deleteSkus(ids) {
    if (Number.isInteger(ids))  {
        ids = [ids]
    }
    return axiosInstance.post(`/admin/skus/delete_all`, {ids})}

// 更改某规格状态 (启用/禁用)
export function changeSkusStatus(id, status) {
    return axiosInstance.post(`/admin/skus/${id}/update_status`, {status})
}

