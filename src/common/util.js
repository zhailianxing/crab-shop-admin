import { ElNotification } from "element-plus"

// 显示成功消息
export function showSuccessMessage(message) {
    ElNotification({
        message,
        type: "success",
        duration: 2500
    })
}

// 显示错误消息
export function showErrorMessage(message) {
    ElNotification({
        message,
        type: "error",
        duration: 2500
    })
}