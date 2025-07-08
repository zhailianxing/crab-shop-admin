import axiosInstance from "~/axios.js";

// 获取图片分类列表
export function getImageCategoryList(page) {
    return axiosInstance.get("/admin/image_class/" + page)
}

// 新增图片分类
export function addImageCategory(data) {
    return axiosInstance.post("/admin/image_class", data)
}

// 修改图片分类
export function updateImageCategory(id, data) {
    return axiosInstance.post("/admin/image_class/"+id, data)
}

// 删除图片分类
export function deleteImageCategory(id) {
    return axiosInstance.post("/admin/image_class/"+id +"/delete", )
}