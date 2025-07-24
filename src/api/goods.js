
import axiosInstance from "~/axios.js";

// 获取商品列表
export function getGoodsList(page, queryData = { tab: null, title: null, category_id: null}) {
    // 过滤出 value不为null的
    const filtered = Object.fromEntries(
        Object.entries(queryData).filter(([_, v]) => v !== null)
    );
    const encoded = new URLSearchParams(filtered).toString();
    return axiosInstance.get(`/admin/goods/${page}` + "?" + encoded)
}

// 更改商品状态(上架、下架)
export function changeGoodsStatus(ids, status) {
    if (typeof ids == 'number') {
        ids = [ids]
    }
    return axiosInstance.post(`/admin/goods/${id}/update_status`, { ids, status })
}

// 新增商品
export function addGoods(data) {
    return axiosInstance.post(`/admin/goods`, data)
}

// 删除商品
export function delGoods(ids) {
    return axiosInstance.post(`/admin/goods/delete_all`, { ids })
}

// 修改商品
export function modifyGoods(id,data) {
    return axiosInstance.post(`/admin/goods/${id}`, data)
}
