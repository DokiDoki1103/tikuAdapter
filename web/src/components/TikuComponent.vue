<template>
  <a-layout>
    <a-layout-content>
      <div style="padding:50px 50px">
        <a-table :columns="columns" :data-source="data" :pagination="false" :scroll="{ y: 500 }">
          <template #customTitle>
      <span>
        <smile-outlined/>
        Name
      </span>
          </template>
          <template #type="{ record }">
            <a-tag
                :color="record.type === 0 ? 'green' : record.type === 1 ? 'cyan' : 'orange'"
            >
              {{ quesType[record.type] }}
            </a-tag>
          </template>
          <template #action="{ record }">
            <div style="display:flex;justify-content: space-between">
              <a-button type="primary" ghost @click="showModal(record)">操作</a-button>
              <a-button type="primary" ghost danger @click="deleteQuestion(record)">删除</a-button>
            </div>
          </template>
        </a-table>
        <a-pagination
            @change="onChange"
            :current="page.pageNo"
            :total="page.total"
            show-size-changer
            :page-size-options="page.pageSizeOptions"
            :page-size="page.pageSize"
            @showSizeChange="onShowSizeChange"
            style="margin-top: 20px; text-align: right;"
        />
      </div>
      <a-modal :visible="visible" title="Basic Modal" @ok="handleOk" @cancel="visible=false">
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
      </a-modal>

    </a-layout-content>
  </a-layout>


</template>
<script>
import { SmileOutlined } from '@ant-design/icons-vue'
import { defineComponent, ref } from 'vue'
import axios from 'axios'

const data = ref([]) // 在 setup() 外部定义
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: 100
  },
  {
    title: '类型',
    key: 'type',
    dataIndex: 'type',
    width: 100,
    slots: {
      customRender: 'type',
    },
  },
  {
    title: '问题',
    dataIndex: 'question',
    key: 'question',
  },
  {
    title: '答案',
    dataIndex: 'answer',
    key: 'answer',
  },
  {
    title: '操作',
    fixed: 'right',
    key: 'action',
    width: 180,
    slots: {
      customRender: 'action',
    },
  },
]
const page = ref({
  pageSizeOptions: ['10', '20', '50', '100', '500'],
  pageNo: 1,
  pageSize: 10,
  total: 0
})

const visible = ref(false)
const quesType = ref({
  '0': '单选题',
  '1': '多选题',
  '3': '判断题',
  '4': '简答题',
})
export default defineComponent({
  setup () {
    return {
      page,
      data,
      columns,
      quesType,
      visible
    }
  },
  components: {
    SmileOutlined
  },
  mounted () {
    this.fetchData()
  },

  methods: {
    async fetchData () {
      const res = await axios.get(`/adapter-service/questions?pageSize=${page.value.pageSize}&pageNo=${page.value.pageNo - 1}`)
      data.value = res.data.items
      this.page.total = res.data.total
    },
    deleteQuestion (id) {
      console.log(id)
    },
    onChange (page) {
      this.page.pageNo = page
      this.fetchData()
    },
    showModal (e) {
      console.log(e)
      visible.value = true
      console.log(visible.value)
    },
    handleOk (e) {
      console.log(e)
      visible.value = false
    },
    onShowSizeChange (current, pageSize) {
      this.page.pageSize = pageSize
      this.page.pageNo = 1
      this.fetchData()
    }
  }
})
</script>
