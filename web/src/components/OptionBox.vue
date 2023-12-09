<template>
  <div>
    <a-row v-for="(item, index) in data" :key="index" class="rowStyle">
      <a-col :span="2">
        <div :class="item.isCheck ? 'CheckStyle Check' : 'CheckStyle CheckNo'" @click="handleCheck(item, index)">
          <a-tooltip placement="top">
            <template #title>
              {{ item.isCheck ? `取消选项` : '设为答案'  }}
            </template>
            {{ Index[index] }}
          </a-tooltip>
        </div>
      </a-col>
      <a-col :span="21">
        <a-input v-model:value="item.content" @change="(e)=>item.isCheck ? handleContent(e,index) : null" />
      </a-col>
      <a-col :span="1">
        <a-button v-if="index == 0" @click="addOption" type="link" primary>
          <template #icon>
            <PlusCircleOutlined />
          </template>
        </a-button>
        <a-button v-if="index > 0" @click="delOption(item,index)" type="link" danger>
          <template #icon>
            <MinusCircleOutlined />
          </template>
        </a-button>
      </a-col>
    </a-row>
  </div>
</template>
  
<script >
// 创建选项数组
import { ref } from 'vue'
import {
  PlusCircleOutlined,
  MinusCircleOutlined
} from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
export default ({
  name: 'OptionBox',
  components: {
    PlusCircleOutlined,
    MinusCircleOutlined
  },
  props: {
    options: {
      type: String,
      default: ''
    },
    type: {
      type: Number,
      default: 0
    },
    answer: {
      type: String,
      default: ''
    }
  },
  setup(props) {
    const Index = ref({
      0: 'A',
      1: 'B',
      2: 'C',
      3: 'D',
      4: 'E',
      5: 'F',
      6: 'G',
      7: 'H',
      8: 'I',
      9: 'J',
      10: 'K',
      11: 'L',
    })
    const data = ref([])
    const optionsArr = ref([])
    const answerArr = ref([])
    answerArr.value = JSON.parse(props.answer)
    if (props.options.length > 0 && props.options != '[]') {
      optionsArr.value = JSON.parse(props.options)
      optionsArr.value.map((item, index) => {
        data.value.push({
          isCheck: answerArr.value.indexOf(item) > -1 ? true : false,
          index: index,
          content: item
        })
      })
    } else {
      data.value.push({
        isCheck: true,
        index: 0,
        content: ''
      })
    }
    return {
      data,
      Index,
      answerArr
    }
  },

  methods: {
    addOption() {
      if (this.data.length >= 12) {
        message.info('添加选项已达上限');
        return
      }
      this.data.push({ isCheck: false, option: '', content: '' })
    },
    delOption(item, index) {
      this.data.splice(index, 1)
      if (item.isCheck) {
        this.answerArr.splice(this.answerArr.indexOf(item.content), 1)
      }
    },
    handleCheck(item, index) {
      if(item.content == ''){
        message.info('请先填写选项内容');
        return
      }
      this.data[index].isCheck = !item.isCheck
      this.answerArr.push(item.content)
    },
    handleContent(e, index) {
      this.data[index].content = e.target.value
      if (this.data[index].isCheck) {
        this.answerArr.splice(this.answerArr.indexOf(e.target.value), 1, e.target.value)
      }
    },
    getData() {
      let options = []
      this.data.map(item => {
        options.push(item.content)
      })
      return {
        options: JSON.stringify(options),
        answer: JSON.stringify(this.answerArr)
      }
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
</style>