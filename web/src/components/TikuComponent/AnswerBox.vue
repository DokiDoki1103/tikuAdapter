<template>
  <div>
    <a-row v-for="(item, index) in data" :key="index" class="rowStyle">
      <a-col :span="22">
        <a-textarea v-model:value="item.value" />
      </a-col>
      <a-col :span="1">
        <a-button v-if="index === 0" @click="addOption" type="link" primary>
          <template #icon>
            <PlusCircleOutlined />
          </template>
        </a-button>
        <a-button v-if="index > 0" @click="data.splice(index, 1)" type="link" danger>
          <template #icon>
            <MinusCircleOutlined />
          </template>
        </a-button>
      </a-col>
    </a-row>
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
        : [{
            content: '',
            index: 0
          }])
    }
  },

  methods: {
    addOption() {
      this.data.push({
        value: '',
      })
    },
    getData() {
      return JSON.stringify(this.data.map((item) => {
        return item.value
      }))
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