import axiosInstance from "~/axios.js";


// 获取图片分类列表
export function getImageCategoryList(page) {
    return axiosInstance.get("/admin/image_class/" + page)
}