<template>
  <a-layout>
    <a-layout-content>
      <div style="padding:50px 50px">
        <a-card :bordered="false" style="width: 100%; text-align: start;" :bodyStyle="style">
          <template #title>
            <a-input-search v-model:value="searchValue" placeholder="搜索关键字" style="width: 300px" @keyup.enter="onSearch"
              @search="onSearch" />
          </template>
          <template #extra>
            <a-button type="primary" @click="showModal({
              type: 0,
              question: '',
              options: '[]',
              answer: '[]'
            }, 2)">
            <FormOutlined />
            添加
            </a-button>
          </template>
          <a-table :columns="columns" :data-source="data" :pagination="false" :row-key="record => record.id">
            <template #type="{ record }">
              <a-tag :color="record.type === 0 ? 'green' : record.type === 1 ? 'cyan' : 'orange'">
                {{ quesType[record.type] }}
              </a-tag>
            </template>
            <template #answer="{ record }">
              <div >
                <a-tag v-for ="(value,index) in JSON.parse(record.answer)" color="blue" :key="index" >{{value}}</a-tag>
              </div>
            </template>
            <template #action="{ record }">
              <div style="display:flex;justify-content: space-between">
                <a-button type="primary" ghost @click="showModal(record, 1)">编辑</a-button>
                <a-button type="primary" ghost danger @click="deleteQuestion(record.id)">删除</a-button>
              </div>
            </template>
          </a-table>
          <a-pagination @change="onChange" :current="page.pageNo" :total="page.total" show-size-changer
            :page-size-options="page.pageSizeOptions" :page-size="page.pageSize" @showSizeChange="onShowSizeChange"
            style="margin-top: 20px; text-align: right;" />
        </a-card>

      </div>
      <a-modal :visible="visible" title="详情内容" @ok="confirm" @cancel="visible = false" v-if="visible">
        <a-form :model="formData">
          <a-form-item label="类型">
            <a-select v-model:value="formData.type">
              <a-select-option v-for="(value, key) in quesType" :key="key" :value="parseInt(key)">
                {{ value }}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="问题">
            <a-textarea v-model:value="formData.question" />
          </a-form-item>
          <a-form-item label="选项" v-if="formData.type == 0 || formData.type == 1 || formData.type == 3">
            <OptionBox :options="formData.options" :type="formData.type" :answer="formData.answer" ref="optionBox" />
          </a-form-item>
          <a-form-item label="答案" v-else>
            <AnswerBox :answer="formData.answer" ref="answerBox" />
          </a-form-item>
        </a-form>
      </a-modal>
    </a-layout-content>
  </a-layout>
</template>

<script>
import { defineComponent, ref } from 'vue'
import {
  getQuestions,
  updateQuestions,
  delQuestions,
  createQuestions
} from '@/api/api'
import {
  FormOutlined
} from '@ant-design/icons-vue';
import OptionBox from './OptionBox.vue'
import AnswerBox from './AnswerBox.vue'
const style = {
  padding: "0 0 24px"
}
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
    slots: {
      customRender: 'answer',
    },
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

const formData = ref({})
const visible = ref(false)
const action = ref(2) // 1是编辑 2是添加
const searchValue = ref('')
const quesType = ref({
  '0': '单选题',
  '1': '多选题',
  '2': '填空题',
  '3': '判断题',
  '4': '简答题',
})
export default defineComponent({
  setup() {
    return {
      page,
      data,
      columns,
      quesType,
      visible,
      formData,
      style,
      searchValue
    }
  },
  components: {
    OptionBox,
    AnswerBox,
    FormOutlined
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    async fetchData(question) {
      const res = await getQuestions({
        pageSize: page.value.pageSize,
        pageNo: page.value.pageNo - 1,
        question: question || ''
      })
      data.value = res.items
      this.page.total = res.total
    },
    onSearch(value) {
      this.fetchData(value)
    },
    onShowSizeChange(current, pageSize) {
      this.page.pageSize = pageSize
      this.page.pageNo = 1
      this.fetchData()
    },

    onChange(page) {
      this.page.pageNo = page
      this.fetchData()
    },

    showModal(data, act) {
      action.value = act
      formData.value = data
      visible.value = true
    },

    // showAnswer(item){
    //   if(item.length)
    // }

    async deleteQuestion(id) {
      await delQuestions(id)
      visible.value = false
      await this.fetchData()
    },

    async confirm() {
      if (formData.value.type == 0 || formData.value.type == 1 || formData.value.type == 3) {
        const childComponentData = this.$refs.optionBox.getData();
        formData.value.options = childComponentData.options
        formData.value.answer = childComponentData.answer
      } else {
        formData.value.answer = this.$refs.answerBox.getData();
        formData.value.options = "[]"
      }
      formData.value.type = parseInt(formData.value.type)
      if (action.value === 1) {
        await updateQuestions(formData.value)
      } else if (action.value === 2) {
        await createQuestions(formData.value)
      }
      visible.value = false
      await this.fetchData()
    }
  }
})
</script>
