import { ElNotification, ElMessageBox } from "element-plus"


// Toast消息提示： 显示成功消息
export function showSuccessMessage(message) {
    ElNotification({
        message,
        type: "success",
        duration: 2500
    })
}

// 消息提示： 显示错误消息
export function showErrorMessage(message) {
    ElNotification({
        message,
        type: "error",
        duration: 2500
    })
}

// 消息提示框
export function showModal(content, title = "") {
    // 因为点击"确认"、"取消"会分别执行then、catch回调函数，所以要将对象return出去
    return ElMessageBox.confirm(
        content,
        title,
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
        }
    )
}

// 消息提示框 - 提交内容
export function showModalInput() {
  return ElMessageBox.prompt('请输入重修改的名字', '重命名', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
  })
}
