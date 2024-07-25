<template>
  <a-card style="width: 100%">
    <a-table bordered :data-source="dataSource" :columns="columns" :loading="tabLoading">
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
      </template>
    </a-table>
  </a-card>
</template>

<script>
import { getLogList } from '@/api/TikuComponentApi';
import { defineComponent, ref } from 'vue';
const dataSource = ref([])
const tabLoading = ref(true)
export default defineComponent({
  name: 'UserList',
  setup() {
    return {
      dataSource,
      tabLoading,
      columns: [
        {
          title: '被修改的qid',
          dataIndex: 'qid',
          key: 'qid',
        },
        {
          title: '修改时间',
          dataIndex: 'create_time',
          key: 'create_time',
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
      getLogList().then(res => {
        if (res?.status == 200) {
          dataSource.value = res?.data
          tabLoading.value = false
        }
      })
    },
  }
});

</script>

<style scoped></style>
