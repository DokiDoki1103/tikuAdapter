<template>
  <div v-if="data.length">
    <a-row v-for="(item, index) in data" :key="index" class="rowStyle">
      <a-col :span="20">
        <a-textarea v-model:value="item.value" />
      </a-col>
      <a-col :span="4">
        <a-button v-if="index >= 0" @click="addOption" type="link" primary>
          <template #icon>
            <PlusCircleOutlined />
          </template>
        </a-button>
        <a-button v-if="index >= 0" @click="data.splice(index, 1)" type="link" danger>
          <template #icon>
            <MinusCircleOutlined />
          </template>
        </a-button>
      </a-col>
    </a-row>
  </div>
  <div v-else>
    <Button class="addBut" @click="addOption" style="width: 100%">
      添加
    </Button>
  </div>
</template>

<script>
// 创建选项数组
import { ref } from 'vue'
import {
  PlusCircleOutlined,
  MinusCircleOutlined
} from '@ant-design/icons-vue';

export default ({
  name: 'OptionBox',
  components: {
    PlusCircleOutlined,
    MinusCircleOutlined
  },
  props: {
    answer: {
      type: Array,
      default: () => {
        return []
      }
    }
  },
  setup(props) {
    return {
      data: ref(props.answer.length > 0
        ? props.answer.map(i => {
          return {
            value: i
          };
        })
        : [])
    }
  },

  methods: {
    addOption() {
      this.data.push({
        value: '',
      })
    },
    pushAnswer(str){
      this.data.push({
        value: str,
      })
    },
    getData() {
      return JSON.stringify(this.data.map((item) => {
        if (item.value !== null && item.value!== undefined) {
          return item.value
        }
      }).filter(i=>i))
    }
  }
})
</script>

<style scoped>
.rowStyle {
  margin-bottom: 10px;
}

.CheckStyle {
  width: 30px;
  height: 30px;
  line-height: 28px;
  text-align: center;
  border-radius: 50%;
  cursor: pointer;
  font-size: 16px;
  box-sizing: border-box;
}

/* 选中的样式 */
.Check {
  background-color: #1890ff;
  color: #fff;
  border: 1px solid #1890ff;
}

/* 未选中的样式 */
.CheckNo {
  background-color: #fff;
  color: #1890ff;
  border: 1px solid #1890ff;
  box-sizing: border-box;
}

.addBut {
  color: rgba(0,0,0,.65);
  background-color: #fff;
  border-color: #d9d9d9;
  border-style: dashed;
  padding: 0 15px;
  font-size: 14px;
  border-radius: 4px;
  height: 32px;
}

.addBut:focus{
  color: #108ee9;
  background-color: #fff;
  border-color: #108ee9;
}

.addBut:hover{
  color: #108ee9;
  background-color: #fff;
  border-color: #108ee9;
}
</style>