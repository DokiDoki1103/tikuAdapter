<template>
  <a-layout>
    <a-layout-content>

      <div style="padding:50px 50px">
        <a-button type="primary" style="width: 500px;" ghost @click="showModal({},2)">添加</a-button>

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
              <a-button type="primary" ghost @click="showModal(record,1)">操作</a-button>
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

      <a-modal :visible="visible" title="详情内容" @ok="check"
               @cancel="visible=false">
        <a-form :model="formState" v-if="num">
          <a-form-item label="类型">
            <a-input  v-model:value="formState.type"></a-input>
          </a-form-item>
          <a-form-item label="问题">
            <a-textarea
                v-model:value="formState.question"
                :auto-size="{ minRows: 2, maxRows: 5 }"
            />
          </a-form-item>
          <a-form-item label="选项">
            <a-textarea
                v-model:value="formState.options"
               :auto-size="{ minRows: 2, maxRows: 5 }"
            />
          </a-form-item>
          <a-form-item label="答案">
            <a-textarea
                v-model:value="formState.answer"
                :auto-size="{ minRows: 2, maxRows: 5 }"
            />
          </a-form-item>
        </a-form>
        <p v-else>删除将无法恢复，确定删除吗</p>
      </a-modal>
    </a-layout-content>
  </a-layout>


</template>
<script>
import { SmileOutlined } from '@ant-design/icons-vue'
import { defineComponent, ref } from 'vue'
import { getQuestions, updateQuestions, delQuestions, addQuestions } from '@/api/api'

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


const formState = ref({

})
const visible = ref(false)
const num = ref(null)

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
      num,
      columns,
      quesType,
      visible,
      formState,
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
      const res = await getQuestions({
        params: {
          pageSize: page.value.pageSize,
          pageNo: page.value.pageNo - 1
        }
      })
      data.value = res.data.items
      this.page.total = res.data.total
    },

    async delOk () {
      await delQuestions(formState.value.id, formState.value)
      visible.value = false
      this.fetchData()
    },

    async deleteQuestion (record, n) {
      num.value = n
      formState.value = record
      visible.value = true
    },

    onChange (page) {
      this.page.pageNo = page
      this.fetchData()
    },

    showModal (e, n) {
      num.value = n
      formState.value = e
      visible.value = true
    },

    check () {
      formState.value.type = parseInt(formState.value.type)
      if (num.value === 1) {
        this.handleOk()
      } else if (num.value === 2) {
        this.addQues()
      } else {
        this.delOk()
      }
    },

    async addQues () {
      await addQuestions(formState.value)
      visible.value = true
      this.fetchData()
    },

    async handleOk () {
      await updateQuestions(formState.value.id, formState.value)
      console.log(formState.value.id, formState.value)
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
