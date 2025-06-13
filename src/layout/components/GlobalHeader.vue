<template>
    <div class="global-header">
        <div class="global-header-left">
            <el-icon class="g-h-l-icon">
                <House />
            </el-icon>
            <span class="g-h-l-title">兴化大闸蟹</span>
            <el-icon class="g-h-l-icon g-h-l-fold">
                <Fold />
            </el-icon>
            <el-icon class="g-h-l-icon ">
                <Refresh />
            </el-icon>
        </div>
        <div class="global-header-right">
            <el-icon class="g-h-r-icon">
                <VideoCamera />
            </el-icon>
            <el-icon class="g-h-r-icon">
                <FullScreen />
            </el-icon>
            <el-dropdown>
                <span class="el-dropdown-link">
                    <!--  开始： 这里的内容自由替换  -->
                    <el-avatar :size="25" :src="$store.state.user.avatar" />
                    {{ $store.state.user.username }}
                    <!--  结束： 这里的内容自由替换  -->
                    <el-icon class="el-icon--right">
                        <arrow-down />
                    </el-icon>
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item>退出登录</el-dropdown-item>
                        <el-dropdown-item>修改密码</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>

        </div>
    </div>
</template>

<script setup>
//TODO: 优化，导入所有图标，不再一个一个导入
import { House, VideoCamera, FullScreen, Fold, Refresh } from '@element-plus/icons-vue'

</script>

<style lang="scss" scoped>
.global-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #2a93cf;
    height: 50px;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;

    .global-header-left {
        display: flex;
        align-items: center;
        gap: 10px; //gap: 10px 用于设置 Flex 子项之间的间距，避免手动给每个子项加 margin-right 或类似的样式，更方便、语义更清晰。

        .g-h-l-icon {
            font-size: 20px;
            margin-left: 20px;
            height: 50px;
            width: 50px;
            cursor: pointer; // 鼠标悬停时显示手型

            &:hover {
                // 鼠标悬停时，背景颜色变为深蓝色，会将width和height区域的背景颜色改掉，所以设置了height
                background-color: #3a8fc0;
            }

        }

        .g-h-l-title {
            font-size: 20px;
        }

        .g-h-l-fold {
            font-size: 20px;
            margin-left: 60px;
            ;
        }
    }

    .global-header-right {
        display: flex;
        align-items: center;
        gap: 20px;

        .g-h-r-icon {
            font-size: 20px;
            height: 50px;
            width: 50px;
            cursor: pointer; // 鼠标悬停时显示手型

            &:hover {
                // 鼠标悬停时，背景颜色变为深蓝色，会将width和height区域的背景颜色改掉，所以设置了height
                background-color: #3a8fc0;
            }
        }

        .el-dropdown-link {
            display: flex;
            align-items: center;
            gap: 5px;
        }
    }
}
</style>
