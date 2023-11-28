<template>
  <a-layout>
    <a-layout-content>
      <div style="padding:50px 50px">
        <a-button type="primary" style="width: 500px;" ghost @click="showModal({
          type: 0,
          question: '',
          options: '[]',
          answer: ''
        },2)">添加
        </a-button>
        <a-table :columns="columns" :data-source="data" :pagination="false">
          <template #type="{ record }">
            <a-tag :color="record.type === 0 ? 'green' : record.type === 1 ? 'cyan' : 'orange'">
              {{ quesType[record.type] }}
            </a-tag>
          </template>
          <template #action="{ record }">
            <div style="display:flex;justify-content: space-between">
              <a-button type="primary" ghost @click="showModal(record,1)">编辑</a-button>
              <a-button type="primary" ghost danger @click="deleteQuestion(record,0)">删除</a-button>
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
      <a-modal :visible="visible" title="详情内容" @ok="check" @cancel="visible=false">
        <a-form :model="formState" v-if="num">
          <a-form-item label="类型">
            <a-select v-model:value="formState.type">
              <a-select-option v-for="(value, key) in quesType" :key="key" :value="parseInt(key)">
                {{ value }}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="问题">
            <a-textarea v-model:value="formState.question"/>
          </a-form-item>
          <a-form-item label="选项">
            <div class="checkbox-flex-container">
              <a-checkbox
                  v-for="(value, key) in JSON.parse(formState.options)"
                  :key="key"
                  :value="parseInt(key)"
                  class="checkbox-flex-item"
              >
                {{ value }}
              </a-checkbox>
            </div>
          </a-form-item>
          <a-form-item label="答案">
            <a-textarea v-model:value="formState.answer" :auto-size="{ minRows: 2, maxRows: 5 }"/>
          </a-form-item>
        </a-form>
        <p v-else>删除将无法恢复，确定删除吗</p>
      </a-modal>
    </a-layout-content>
  </a-layout>
</template>
<style>
.checkbox-flex-container {
  display: flex;
  flex-direction: column; /* 将元素垂直排列 */
}

.checkbox-flex-item {
  margin-bottom: 10px; /* 可以添加一些底部间距，让它们之间有一定的间隔 */
}
</style>
<script>
import {defineComponent, ref} from 'vue'
import {getQuestions, updateQuestions, delQuestions, addQuestions} from '@/api/api'

const data = ref([])
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


const formState = ref({})
const visible = ref(false)
const num = ref(null)

const quesType = ref({
  '0': '单选题',
  '1': '多选题',
  '3': '判断题',
  '4': '简答题',
})
export default defineComponent({
  setup() {
    return {
      page,
      data,
      num,
      columns,
      quesType,
      visible,
      formState,
    }
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      const res = await getQuestions({
        pageSize: page.value.pageSize,
        pageNo: page.value.pageNo - 1
      })
      data.value = res.data.items
      this.page.total = res.data.total
    },
    delOk() {
      delQuestions(formState.value.id).then(res => {
        console.log(res)
        visible.value = false
        this.fetchData()
      })
    },

    async deleteQuestion(record, n) {
      num.value = n
      formState.value = record
      visible.value = true
    },

    onChange(page) {
      this.page.pageNo = page
      this.fetchData()
    },

    showModal(e, n) {
      num.value = n
      formState.value = e
      visible.value = true
    },

    check() {
      formState.value.type = parseInt(formState.value.type)
      if (num.value === 1) {
        this.handleOk()
      } else if (num.value === 2) {
        this.addQues()
      } else {
        this.delOk()
      }
    },

    addQues() {
      addQuestions(formState.value).then(res => {
        console.log(res)
        visible.value = true
        this.fetchData()
      })
    },

    handleOk() {
      updateQuestions(formState.value.id, formState.value).then(res => {
        console.log(res)
        visible.value = false
      })
    },

    onShowSizeChange(current, pageSize) {
      this.page.pageSize = pageSize
      this.page.pageNo = 1
      this.fetchData()
    }
  }
})
</script>
