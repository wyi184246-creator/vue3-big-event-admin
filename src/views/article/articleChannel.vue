<script setup>
import { ref } from 'vue'
import { artGetChannelsService, artDelChannelsService } from '@/api/article'
import { Edit, Delete } from '@element-plus/icons-vue'
import ChannelEdit from './components/ChannelEdit.vue'
const channelList = ref([])
const loading = ref(false)
const dialog = ref(null)
const mockChannelList = [
  {
    id: 1,
    cate_name: 'Frontend',
    cate_alias: 'frontend'
  },
  {
    id: 2,
    cate_name: 'Backend',
    cate_alias: 'backend'
  },
  {
    id: 3,
    cate_name: 'DevOps',
    cate_alias: 'devops'
  }
]

const onEditChannel = async (row) => {
  dialog.value.open(row)
}

const onDelChannel = async (row) => {
  await ElMessageBox.confirm('确定要删除吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  await artDelChannelsService(row.id)
  ElMessage.success('分类删除成功')
  getChannelList()
}

const getChannelList = async () => {
  try {
    loading.value = true
    const res = await artGetChannelsService()

    const list = res?.data?.data || []
    channelList.value = list.length ? list : mockChannelList

    loading.value = false
  } catch {
    channelList.value = mockChannelList
  }
}
const onAddChannel = () => {
  dialog.value.open({})
}
getChannelList()
const onSuccess = () => {
  getChannelList()
}
</script>
<template>
  <page-container title="文章分类">
    <template #extra>
      <el-button @click="onAddChannel" type="primary">添加分类</el-button>
    </template>
    <el-table v-loading="loading" :data="channelList" style="width: 100%">
      <el-table-column label="序号" width="100" type="index"> </el-table-column>
      <el-table-column label="分类名称" prop="cate_name"></el-table-column>
      <el-table-column label="分类别名" prop="cate_alias"></el-table-column>
      <el-table-column label="操作" width="100">
        <template #default="{ row }">
          <el-button
            :icon="Edit"
            circle
            plain
            type="primary"
            @click="onEditChannel(row)"
          ></el-button>
          <el-button
            :icon="Delete"
            circle
            plain
            type="danger"
            @click="onDelChannel(row)"
          ></el-button>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty description="没有数据" />
      </template>
    </el-table>
    <channel-edit ref="dialog" @success="onSuccess"></channel-edit>
  </page-container>
</template>
<style lang="scss" scoped></style>
