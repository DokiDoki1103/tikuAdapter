<template>
  <a-layout>
    <a-layout-content>
      <div style="padding:50px 50px">
        <a-button type="primary" style="width: 500px;" ghost @click="showModal({
          type: 0,
          question: '',
          options: '[]',
          answer: '[]'
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
              <a-button type="primary" ghost danger @click="deleteQuestion(record.id)">删除</a-button>
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
      <a-modal :visible="visible" title="详情内容" @ok="confirm" @cancel="visible=false">
        <a-form :model="formData">
          <a-form-item label="类型">
            <a-select v-model:value="formData.type">
              <a-select-option v-for="(value, key) in quesType" :key="key" :value="parseInt(key)">
                {{ value }}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="问题">
            <a-textarea v-model:value="formData.question"/>
          </a-form-item>
          <a-form-item label="选项">
            <a-textarea v-model:value="formData.options"/>
            <!--            <div class="checkbox-flex-container">-->
            <!--              <a-checkbox-->
            <!--                  v-for="(value, key) in JSON.parse(formData.options)"-->
            <!--                  :key="key"-->
            <!--                  :value="parseInt(key)"-->
            <!--                  class="checkbox-flex-item"-->
            <!--              >-->
            <!--                {{ value }}-->
            <!--              </a-checkbox>-->
            <!--            </div>-->
          </a-form-item>
          <a-form-item label="答案">
            <a-textarea v-model:value="formData.answer"/>
          </a-form-item>
        </a-form>
      </a-modal>
    </a-layout-content>
  </a-layout>
</template>

<script>
import {defineComponent, ref} from 'vue'
import {getQuestions, updateQuestions, delQuestions, createQuestions} from '@/api/api'

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


const formData = ref({})
const visible = ref(false)
const action = ref(2) // 1是编辑 2是添加

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
      columns,
      quesType,
      visible,
      formData,
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
      data.value = res.items
      this.page.total = res.total
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

    async deleteQuestion(id) {
      await delQuestions(id)
      visible.value = false
      await this.fetchData()
    },

    async confirm() {
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
