<template>
  <div>
    <a-row v-for="(item, index) in data" :key="index" class="rowStyle">
      <a-col :span="2">
        <div :class="item.check ? 'CheckStyle Check' : 'CheckStyle CheckNo'"
             @click="item.check = !item.check">
          <a-tooltip placement="top">
            <template #title>
              {{ item.check ? `取消选项` : '设为答案' }}
            </template>
            {{ item.key }}
          </a-tooltip>
        </div>
      </a-col>
      <a-col :span="21">
        <a-input v-model:value="item.value"/>
      </a-col>
      <a-col :span="1">
        <a-button v-if="index === 0" @click="addOption" type="link" primary>
          <template #icon>
            <PlusCircleOutlined/>
          </template>
        </a-button>
        <a-button v-if="index > 0" @click="data.splice(index, 1)" type="link" danger>
          <template #icon>
            <MinusCircleOutlined/>
          </template>
        </a-button>
      </a-col>
    </a-row>
  </div>
</template>

<script>
// 创建选项数组
import {
  PlusCircleOutlined,
  MinusCircleOutlined
} from '@ant-design/icons-vue';
import {ref} from "vue";

export default ({
  name: 'OptionBox',
  components: {
    PlusCircleOutlined,
    MinusCircleOutlined
  },
  props: {
    options: {
      type: Array,
      default: () => {
        return []
      }
    },
    type: {
      type: Number,
      default: 0
    },
    answer: {
      type: Array,
      default: () => {
        return []
      }
    }
  },
  setup(props) {
    return {
      data: ref(props.options.map((item, index) => {
        return {
          check: props.answer.includes(item),
          key: String.fromCharCode(65 + index),
          value: item
        }
      })),
    }
  },

  methods: {
    addOption() {
      this.data.push({
        check: false,
        key: String.fromCharCode(65 + this.data.length),
        value: ''
      })
    },
    getData() {
      return {
        options: JSON.stringify(this.data.map(i => {
          return i.value
        })),
        answer: JSON.stringify(this.data.filter(i => i.check).map(i => {
          return i.value
        }))
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