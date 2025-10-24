<template>
  <a-card>
    <a-row>
      <a-col :span="5">
        <a-form-item label="平台" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-select v-model:value="platValue" show-search placeholder="请选择" :options="platType"
                    @change="selectChange">
          </a-select>
        </a-form-item>
      </a-col>
      <a-col :span="5">
        <a-form-item label="课程" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-select v-model:value="courseValue" show-search placeholder="请选择" :options="courses"
                    @change="onFilterChange">
          </a-select>
        </a-form-item>
      </a-col>
      <a-col :span="3">
        <a-form-item label="类型" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-select v-model:value="typeValue" show-search placeholder="请选择" :options="typeList"
                    @change="onFilterChange">
          </a-select>
        </a-form-item>
      </a-col>
      <a-col :span="5">
        <a-form-item label="问题" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-input v-model:value="searchValue.question" placeholder="搜索问题"/>
        </a-form-item>
      </a-col>

      <a-col :span="3">
        <a-form-item label="无答案" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-checkbox v-model:checked="searchValue.onlyShowEmptyAnswer"></a-checkbox>
        </a-form-item>

      </a-col>

      <a-col :span="2">
        <a-form-item :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-button type="primary" @click="onSearch" @keyup.enter="onSearch">
            <SearchOutlined/>
            搜索
          </a-button>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>

  <a-layout>
    <a-layout-content>
      <div style="padding-top: 24px; text-align: center;">
        <a-card :bordered="false" style="width: 100%; text-align: start;" :bodyStyle="style">
          <template #extra>
            <a-button type="primary" @click="navigateToImport" style="margin-right: 12px;">
              <DownloadOutlined/>
              智能导入
            </a-button>
            <a-button type="primary" @click="showModal({
              type: 0,
              question: '',
              options: '[]',
              answer: '[]'
            }, 2)">
              <FormOutlined/>
              添加
            </a-button>
          </template>
          <a-table :columns="columns" :data-source="data" :pagination="false" :row-key="record => record.id"
                   :loading="tabLoading" :scroll="{ x: '2000px' }">
            <template #type="{ record }">
              <a-tag :color="record.type === 0 ? 'green' : record.type === 1 ? 'cyan' : 'orange'">
                {{ quesType[record.type] }}
              </a-tag>
            </template>
            <template #answer="{ record }">
              <div>
                <a-tag v-for="(value, index) in (record.answer ? JSON.parse(record.answer) : [])" color="blue"
                       :key="index">{{ record.source === 2 ? index : value }}
                </a-tag>
              </div>
            </template>

            <template #plat="{ record }">
              <span>
                {{getPlatStr(record.plat)}}
              </span>
            </template>


            <template #action="{ record }">
              <div style="display:flex;justify-content: space-between">
                <a-button type="primary" ghost @click="showModal(record, 1)">编辑</a-button>
                <a-button type="primary" ghost danger @click="deleteQuestion(record.id)">删除</a-button>
              </div>
            </template>
          </a-table>
          <a-pagination @change="onChange" :current="page.pageNo" :total="page.total" show-size-changer
                        :page-size-options="page.pageSizeOptions" :page-size="page.pageSize"
                        @showSizeChange="onShowSizeChange"
                        style="margin-top: 20px; text-align: right;"/>
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
          <a-form-item label="问题" class="margin_top">
            <a-textarea v-model:value="formData.question"/>
          </a-form-item>
          <div v-if="/[013]/.test(formData.type)" class="margin_top">
            <a-form-item label="选项">
              <OptionBox :options="formData.options ? JSON.parse(formData.options) : []" :type="formData.type"
                         :answer="answer" ref="optionBox"/>
            </a-form-item>
          </div>
          <div v-else>
            <a-card :bordered="false" style="width: 100%; text-align: start;" :bodyStyle="style">
              <a-tabs default-active-key="1">
                <a-tab-pane tab="默认选项" key="1">
                  <AnswerBox :answer="answer" ref="answerBox"/>
                </a-tab-pane>
                <a-tab-pane tab="提交附件" key="2">
                  <div>
                    <a-upload-dragger name="file" multiple="true" show-upload-list="true"
                                      action='/adapter-service/upload?parentDir=myword' @change="handleChange">
                      <p class="ant-upload-drag-icon">
                        <upload-outlined></upload-outlined>
                      </p>
                      <p class="ant-upload-text">点击或者拖动文件到该区域来上传</p>
                      <p class="ant-upload-hint">请不要上传敏感数据，银行卡号和密码，信用卡号有效期和安全码</p>
                    </a-upload-dragger>
                  </div>

                </a-tab-pane>
              </a-tabs>


            </a-card>
          </div>
        </a-form>
      </a-modal>
    </a-layout-content>
  </a-layout>
</template>

<script>
import {defineComponent, ref} from 'vue';
import {useRouter} from 'vue-router';
import {
  getPlat,
  getCourses,
  getQuestions,
  updateQuestions,
  delQuestions,
  createQuestions,
} from '@/api/TikuComponentApi'
import {
  FormOutlined,
  DownloadOutlined,
  SearchOutlined,
  UploadOutlined
} from '@ant-design/icons-vue';
import OptionBox from '../../components/OptionBox/index.vue'
import AnswerBox from '../../components/AnswerBox/index.vue'
import {questionType} from "@/utils/uitls";
import {message} from "ant-design-vue";

const style = {
  padding: "0 0 24px"
}
const fileList = ref([])
const tabLoading = ref(true)
const data = ref([])
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: 90
  },
  {
    title: '类型',
    key: 'type',
    dataIndex: 'type',
    width: 80,
    slots: {
      customRender: 'type',
    },
  },
  {
    title: '课程名称',
    width: 170,
    dataIndex: 'course_name',
    key: 'course_name',
  },
  {
    title: '问题',
    width: 700,
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
    title: '平台',
    dataIndex: 'plat',
    key: 'plat',
    width: 100,
    slots: {
      customRender: 'plat',
    },
  },
  {
    title: '拓展参数',
    dataIndex: 'extra',
    key: 'extra',
    width: 300,
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
  pageSizeOptions: ['10', '20', '50', '100'],
  pageNo: 1,
  pageSize: 10,
  total: 0
})

const typeValue = ref(null)
const typeList = ref([])
const platType = ref([])
const courses = ref([])
const courseValue = ref(null)
const platValue = ref(null)
const formData = ref({})
const answer = ref([])
const visible = ref(false)
const action = ref(2) // 1是编辑 2是添加
const searchValue = ref({})
const quesType = ref(questionType)
export default defineComponent({
  name: 'TikuComponent',
  setup() {
    const router = useRouter();
    const navigateToImport = () => {
      router.push('/adapter/import');
    };
    return {
      tabLoading,
      typeList,
      typeValue,
      courses,
      platType,
      courseValue,
      platValue,
      answer,
      page,
      data,
      columns,
      quesType,
      visible,
      formData,
      style,
      searchValue,
      labelCol: {span: 8},
      wrapperCol: {span: 14},
      navigateToImport,
      fileList
    }
  },
  components: {
    UploadOutlined,
    OptionBox,
    AnswerBox,
    FormOutlined,
    SearchOutlined,
    DownloadOutlined
  },
  mounted() {
    this.getType()
    this.getCourses()
    this.getPlat()
    this.fetchData()
  },
  methods: {
    async selectChange(info) {
      const list = await getCourses({plat: info})
      if (list?.data && list?.data.length) {
        await this.getCourses(list?.data)
      } else {
        courses.value = []
      }
      // Reset page number when platform changes and fetch data
      this.page.pageNo = 1
      await this.fetchData()
    },
    async onFilterChange() {
      // Reset page number when any filter changes and fetch data
      this.page.pageNo = 1
      await this.fetchData()
    },
    handleChange(info) {
      const status = info.file.status;
      if (status === 'done') {
        message.success(`${info.file.name} 上传成功`);
        fileList.value.push(info.file.response)
        this.$refs.answerBox.pushAnswer(info.file.response)
      } else if (status === 'error') {
        // 处理上传失败
        message.error(`${info.file.name} 上传失败`);
      }
    },
    typeObjects(types) {
      return Object.entries(types).map(([key, value]) => ({
        value: parseInt(key),
        label: value
      }));
    },
    getType() {
      const arr = this.typeObjects(questionType)
      typeList.value = arr
    },
    async getCourses(data) {
      if (data) {
        courses.value = data.map(i => {
          return {
            value: i,
            label: i,
          }
        })
      } else {
        const list = await getCourses()

        if (list?.data && list?.data.length) {
          courses.value = list?.data.map(i => {
            return {
              value: i,
              label: i,
            }
          })
        }
      }
    },

     getPlatStr(plat){
      for (let valueElement of platType.value) {
        if (plat === valueElement.value){
          return valueElement.label
        }
      }
    },
    async getPlat() {
      const list = await getPlat()
      if (list?.data && list?.data.length) {
        platType.value = list?.data.map(i => {
          return {
            value: parseInt(i.Value),
            label: i.Label,
          }
        })
      }
    },
    async fetchData() {
      fileList.value = []
      tabLoading.value = true
      getQuestions({
        courseName: courseValue.value ? courseValue.value : "",
        plat: !isNaN(parseInt(platValue.value)) ? parseInt(platValue.value) : -1,
        type: !isNaN(parseInt(typeValue.value)) ? parseInt(typeValue.value) : -1,
        pageSize: page.value.pageSize,
        pageNo: page.value.pageNo - 1,
        question: searchValue.value.question || '',
        source: parseInt(searchValue.value.source) || 0,
        extra: searchValue.value.extra || '',
        onlyShowEmptyAnswer: searchValue.value.onlyShowEmptyAnswer || false
      }).then(res => {
        data.value = res?.data?.items
        tabLoading.value = false
        this.page.total = res?.data?.total
      })

    },
    onSearch() {
      // Reset page number when search filters change
      this.page.pageNo = 1
      this.fetchData()
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

    showModal(record, act) {
      if (record.options === '[]' || record.options === '') {
        record.options = '[]'
      }
      if (record.answer === '[]' || record.answer === '') {
        record.answer = '[]'
      }

      formData.value = record
      answer.value = formData.value.answer ? JSON.parse(formData.value.answer) : []
      action.value = act
      fileList.value = []
      visible.value = true
    },

    async deleteQuestion(id) {
      fileList.value = []
      await delQuestions({id: id}).then(res => {
        res.status == 200 ? message.success(`删除成功`) : message.error(`删除失败`)
      })
      visible.value = false
      await this.fetchData()
    },


    async confirm() {
      formData.value.type = parseInt(formData.value.type)

      if (/[0|1|3]/.test(formData.value.type)) {
        const childComponentData = this.$refs.optionBox.getData()
        Object.assign(formData.value, childComponentData)
      } else {
        const textAns = this.$refs.answerBox.getData()
        console.log('textAns==>', textAns)
        formData.value.answer = textAns
      }
      let response;
      const successMessage = action.value === 1 ? '修改成功' : '创建成功';
      const errorMessage = action.value === 1 ? '修改失败' : '创建失败';
      if (action.value === 1) {
        response = await updateQuestions(formData.value)
      } else if (action.value === 2) {
        response = await createQuestions([formData.value])
      }
      if (response.status === 200) {
        message.success(successMessage);
      } else {
        message.error(errorMessage);
      }
      fileList.value = []
      visible.value = false
      await this.fetchData()
    },
  }
})
</script>
<style scoped>
header {
  width: 100%;
  height: 60px;
  background-color: #195ca3;
  border-bottom: 1px solid #e8e8e8;
  line-height: 60px;
  font-size: 20px;
  font-weight: 600;
  text-align: center;
  color: #fff;
}

.ant-form-item {
  margin-bottom: 0;
}

.margin_top {
  margin-top: 5%
}
</style>
