<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in gardenTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="规格:" prop="spec">
          <el-input v-model="formData.spec" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="功率:" prop="power">
          <el-input v-model.number="formData.power" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Garden'
}
</script>

<script setup>
import {
  createGarden,
  updateGarden,
  findGarden
} from '@/api/garden'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const gardenTypeOptions = ref([])
const formData = ref({
            name: '',
            type: undefined,
            spec: '',
            power: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGarden({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.regarden
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    gardenTypeOptions.value = await getDictFunc('gardenType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createGarden(formData.value)
               break
             case 'update':
               res = await updateGarden(formData.value)
               break
             default:
               res = await createGarden(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
