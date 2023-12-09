<template>
    <div>
        <a-row v-for="(item, index) in data" :key="index" class="rowStyle">
            <a-col :span="22">
                <a-textarea v-model:value="item.content" />
            </a-col>
            <a-col :span="1">
                <a-button v-if="index == 0" @click="addOption" type="link" primary>
                    <template #icon>
                        <PlusCircleOutlined />
                    </template>
                </a-button>
                <a-button v-if="index > 0" @click="delOption(index)" type="link" danger>
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
        answer: {
            type: String,
            default: ''
        }
    },
    setup(props) {
        const data = ref([])
        const answerArr = ref([])
        if (props.answer != '[]' && props.answer != '') {
            answerArr.value = JSON.parse(props.answer)
            data.value = answerArr.value.map((item, index) => {
                return {
                    content: item,
                    index: index
                }
            })
        }else{
            data.value = [{
                content: '',
                index: 0
            }]
        }
        return {
            data
        }
    },

    methods: {
        addOption() {
            if (this.data.length >= 12) {
                message.info('添加已达上限');
                return
            }
            this.data.push({
                content: '',
                index: this.data.length
            })
        },
        delOption(index) {
            this.data.splice(index, 1)
        },
        getData() {
            let arr = []
            this.data.map((item) => {
                arr.push(item.content)
            })
            return JSON.stringify(arr)
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