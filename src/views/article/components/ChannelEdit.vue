<script setup>
import { ref } from 'vue'
import { artAddChannelsService, artEditChannelsService } from '@/api/article'
import { ElMessage } from 'element-plus'
const formModel = ref({
  cate_name: '',
  cate_alias: ''
})
const rules = {
  cate_name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' },
    {
      pattern: /^\S{1,10}$/,
      message: '分类名称只能包含字母、数字和中文',
      trigger: 'blur'
    }
  ],
  cate_alias: [
    { required: true, message: '请输入分类别名', trigger: 'blur' },
    {
      pattern: /^[a-zA-Z0-9]{1,15}$/,
      message: '分类别名必须是1-15个字母或数字',
      trigger: 'blur'
    }
  ]
}
const emit = defineEmits(['success'])
const onSubmit = async () => {
  formRef.value.validate()
  const isEdit = formModel.value.id
  if (isEdit) {
    // 编辑分类的逻辑
    await artEditChannelsService(formModel.value)
    ElMessage.success('分类编辑成功')
  } else {
    // 添加分类的逻辑
    await artAddChannelsService(formModel.value)

    ElMessage.success('分类添加成功')
  }
  dialogVisible.value = false
  emit('success')
}
const dialogVisible = ref(false)
const open = (row) => {
  console.log(row)
  dialogVisible.value = true
  formModel.value = {
    ...row
  }
}
const formRef = ref(null)
defineExpose({
  open
})
</script>

<template>
  <el-dialog
    v-model="dialogVisible"
    :title="formModel.id ? '编辑分类' : '添加分类'"
    width="30%"
  >
    <el-form
      ref="formRef"
      :model="formModel"
      :rules="rules"
      label-width="100px"
      style="padding-right: 30px"
    >
      <el-form-item label="分类名称" prop="cate_name">
        <el-input v-model="formModel.cate_name" placeholder="请输入分类名称" />
      </el-form-item>
      <el-form-item label="分类别名" prop="cate_alias">
        <el-input v-model="formModel.cate_alias" placeholder="请输入分类别名" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="onSubmit"> 确认 </el-button>
      </span>
    </template>
  </el-dialog>
</template>
