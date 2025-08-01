<template>
    <FormDrawer ref="formDrawerRef" title="设置商品规格" size="50%" destroy-on-close @submitEmit="handleSubmit()"
        @close="resetForm()">
        <el-form :model="form" ref="formRef">
            <el-form-item label="规格类型">
                <el-radio-group v-model="form.sku_type">
                    <el-radio :value="0">单规格</el-radio>
                    <el-radio :value="1">多规格</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="市场价格">
                <el-input v-model="form.sku_value.oprice" style="width: 30%">
                    <template #append>元</template>
                </el-input>
            </el-form-item>
            <el-form-item label="销售价格">
                <el-input v-model="form.sku_value.pprice" style="width: 30%">
                    <template #append>元</template>
                </el-input>
            </el-form-item>
            <el-form-item label="成本价格">
                <el-input v-model="form.sku_value.cprice" style="width: 30%">
                    <template #append>元</template>
                </el-input>
            </el-form-item>
            <el-form-item label="商品重量">
                <el-input v-model="form.sku_value.weight" style="width: 30%">
                    <template #append>公斤</template>
                </el-input>
            </el-form-item>
            <el-form-item label="商品体积">
                <el-input v-model="form.sku_value.volume" style="width: 30%">
                    <template #append>立方米</template>
                </el-input>
            </el-form-item>
        </el-form>
    </FormDrawer>

</template>


<script setup>

import FormDrawer from '~/components/FormDrawer.vue'
import { reactive, ref } from 'vue'
import { updateGoodsSkus } from '~/api/goods.js'
import { showSuccessMessage } from '~/common/util.js'

const goodId = ref(0)
const formDrawerRef = ref(null)
const form = reactive({
    "sku_type": 0,
    "sku_value": {
        "oprice": 0,
        "pprice": 0,
        "cprice": 0,
        "weight": 0,
        "volume": 0
    },
})
const open = (id, goodItem) => {
    goodId.value = id
    form.sku_type = goodItem.sku_type
    form.sku_value = goodItem.sku_value ? goodItem.sku_value : {
        "oprice": 0,
        "pprice": 0,
        "cprice": 0,
        "weight": 0,
        "volume": 0
    }
    formDrawerRef.value.open()
}

const close = () => {
    formDrawerRef.value.close()
}

const resetForm = () => {
    form.sku_type = 0
    form.sku_value = {
        "oprice": 0,
        "pprice": 0,
        "cprice": 0,
        "weight": 0,
        "volume": 0
    }
}

defineExpose({
    open,
    close
})

const emit = defineEmits(["refreshData"])
const handleSubmit = () => {
    // 提交接口
    updateGoodsSkus(goodId.value, form).then(res => {
        showSuccessMessage("更新商品规格成功")
        close()
        // 通知父组件刷新页面
        emit("refreshData")
    })

}

</script>