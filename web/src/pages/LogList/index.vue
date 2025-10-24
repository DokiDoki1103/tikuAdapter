<template>
  <a-card style="width: 100%">
    <a-table
      bordered
      :data-source="dataSource"
      :columns="columns"
      :loading="tabLoading"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #createTime="{ text }">
        {{ formatDate(text) }}
      </template>
      <template #action="{ record }">
        <a-tag v-if="record.action == 0" color="green">
          新增
        </a-tag>
        <a-tag v-else-if="record.action == 1" color="blue">
          更新
        </a-tag>
        <a-tag v-else-if="record.action == 2" color="red">
          删除
        </a-tag>
        <a-tag v-else-if="record.action == 3" color="green">
          查询
        </a-tag>
      </template>
    </a-table>
  </a-card>
</template>

<script>
import { getLogList } from '@/api/TikuComponentApi';
import { defineComponent, ref, reactive } from 'vue';
const dataSource = ref([])
const tabLoading = ref(true)
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true,
  showQuickJumper: true,
  showTotal: (total) => `共 ${total} 条`,
  pageSizeOptions: ['10', '20', '50', '100']
})

export default defineComponent({
  name: 'UserList',
  setup() {
    return {
      dataSource,
      tabLoading,
      pagination,
      columns: [
        {
          title: 'qid',
          dataIndex: 'qid',
          key: 'qid',
        },
        {
          title: '修改时间',
          dataIndex: 'create_time',
          key: 'create_time',
          slots: {
            customRender: 'createTime',
          },
        },
        {
          title: '改前答案',
          dataIndex: 'old_answer',
          key: 'old_answer',
        },
        {
          title: '新答案',
          dataIndex: 'new_answer',
          key: 'new_answer',
        },
        {
          title: '修改人',
          dataIndex: 'user_id',
          key: 'user_id',
        },
        {
          title: '操作',
          dataIndex: 'action',
          key: 'action',
          slots: {
            customRender: 'action',
          },
        },
      ],
    };
  },
  mounted() {
    this.handleLogList()
  },
  methods: {
    handleLogList() {
      tabLoading.value = true
      getLogList({
        page: pagination.current,
        pageSize: pagination.pageSize
      }).then(res => {
        if (res?.status == 200) {
          dataSource.value = res?.data?.data || []
          pagination.total = res?.data?.total || 0
          pagination.current = res?.data?.page || 1
          pagination.pageSize = res?.data?.pageSize || 10
          tabLoading.value = false
        }
      })
    },
    handleTableChange(paginationData) {
      pagination.current = paginationData.current
      pagination.pageSize = paginationData.pageSize
      this.handleLogList()
    },
    formatDate(dateString) {
      if (!dateString) return '-'
      const date = new Date(dateString)
      if (isNaN(date.getTime())) return '-'

      const year = date.getFullYear()
      const month = String(date.getMonth() + 1).padStart(2, '0')
      const day = String(date.getDate()).padStart(2, '0')
      const hours = String(date.getHours()).padStart(2, '0')
      const minutes = String(date.getMinutes()).padStart(2, '0')
      const seconds = String(date.getSeconds()).padStart(2, '0')

      return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
    }
  }
});

</script>

<style scoped></style>
