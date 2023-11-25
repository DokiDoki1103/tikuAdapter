<template>
  <a-table :columns="columns" :data-source="data" :pagination="false">
    <template #name="{ text }">
      <a>{{ text }}</a>
    </template>
    <template #customTitle>
      <span>
        <smile-outlined />
        Name
      </span>
    </template>
    <template #tags="{ text: tags }">
      <span>
        <a-tag
            v-for="tag in tags"
            :key="tag"
            :color="tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'"
        >
          {{ tag.toUpperCase() }}
        </a-tag>
      </span>
    </template>
    <template #action="{ record }">
      <span>
        <a-button type="primary" danger @click="deleteQuestion(record)">删除 </a-button>
      </span>
    </template>
  </a-table>
  <a-pagination
      v-model:current="page.pageNo"
      :total="page.total"
      show-size-changer
      :page-size="page.pageSize"
      @showSizeChange="onShowSizeChange"
      style="margin-top: 20px; text-align: right;"
  />
</template>
<script>
import { SmileOutlined } from '@ant-design/icons-vue';
import {defineComponent, ref} from 'vue';
import axios from "axios";

const data = ref([]); // 在 setup() 外部定义
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id'
  },
  {
    title: '类型',
    key: 'type',
    dataIndex: 'type'
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
    key: 'action',
    slots: {
      customRender: 'action',
    },
  },
];
const page = ref({
  pageSizeOptions: ['3','20','50','100','500'],
  pageNo: 0,
  pageSize: 3,
  total: 0,
});
export default defineComponent({
  setup() {
    return {
      page,
      data,
      columns,
    };
  },
  components: {
    SmileOutlined
  },
  mounted() {
   this.fetchData()
  },
  methods:{
    async fetchData(){
      const res = await axios.get(`/adapter-service/questions?pageSize=${page.value.pageSize}&pageNo=${page.value.pageNo}`)
      data.value = res.data.items
    },
    deleteQuestion(id){
      console.log(id)
    },
    onShowSizeChange(){

    }
  }
});
</script>
