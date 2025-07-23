import axiosInstance from '~/axios.js'


// 获取公告列表
export function getCouponList(page) {
    return axiosInstance.get(`admin/coupon/${page}`)
}

// 增加公告
export function addCoupon(data = {}) {
    return axiosInstance.post("/admin/coupon", data)
}

// 增加公告
export function modifyCoupon(id, data = {}) {
    return axiosInstance.post(`admin/coupon/${id}`, data)
}

// 删除公告
export function deleteCoupon(id) {
    return axiosInstance.post(`admin/coupon/${id}/delete`)
}

// 更改优惠卷状态为失效
export function changeCouponStatus(id) {
    return axiosInstance.post(`/admin/coupon/${id}/update_status`, {status:0})
}
