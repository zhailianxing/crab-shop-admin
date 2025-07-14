import axiosInstance from "~/axios.js";


// 获取管理列表
export function getManagerList(page, queryData = {limit: 10, keyword: ""}) {
    const encoded = new URLSearchParams(queryData).toString();
    return axiosInstance.get(`/admin/manager/${page}` + "?" + encoded)
}
