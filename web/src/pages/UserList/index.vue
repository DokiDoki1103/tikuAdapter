<template>
  <a-card style="width: 100%">
    <template #extra><a-button type="primary" @click="visible = true">添加用户</a-button></template>
    <a-table bordered :data-source="dataSource" :columns="columns" :loading="tabLoading">
      <template #operation="{ record }">
        <a-popconfirm v-if="dataSource.length" ok-text="确定" cancel-text="取消" title="确定删除该用户吗?"
          @confirm="onDelete(record.id)">
          <a style="color: red;">删除</a>
        </a-popconfirm>
      </template>
    </a-table>
  </a-card>
  <a-modal v-model:visible="visible" title="添加用户" @ok="onSubmit" :afterClose="resetFormState" okText="确定"
    cancelText="取消">
    <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="用户名">
        <a-input v-model:value="formState.username" />
      </a-form-item>
      <a-form-item label="密码">
        <a-input v-model:value="formState.password" />
      </a-form-item>
      <a-form-item label="昵称">
        <a-input v-model:value="formState.nickname" />
      </a-form-item>
      <a-form-item label="权限">
        <a-input :disabled="true" placeholder="暂未开放" />
      </a-form-item>
    </a-form>

  </a-modal>
</template>

<script>
import { getUserList, delUser, addUser } from '@/api/TikuComponentApi';
import { message } from 'ant-design-vue';
import { defineComponent, ref, toRaw } from 'vue';
const dataSource = ref([])
const visible = ref(false);
const formState = ref({
  username: '',
  password: '',
  nickname: '',
  perms: '[]',
})
const tabLoading = ref(true)
export default defineComponent({
  name: 'UserList',
  setup() {
    return {
      labelCol: { style: { width: '150px' } },
      wrapperCol: { span: 12 },
      formState,
      visible,
      dataSource,
      tabLoading,
      columns: [
        {
          title: '用户名',
          dataIndex: 'username',
          key: 'username',
        },
        {
          title: '密码',
          dataIndex: 'password',
          key: 'password',
        },
        {
          title: '昵称',
          dataIndex: 'nickname',
          key: 'nickname',
        },
        {
          title: '权限',
          dataIndex: 'perms',
          key: 'perms',
        },
        {
          title: '操作',
          dataIndex: 'operation',
          key: 'operation',
          slots: {
            customRender: 'operation',
          },
        },
      ],
    };
  },
  mounted() {
    this.handleUserList()
  },
  methods: {
    handleUserList() {
      tabLoading.value=true
      getUserList().then(res => {
        if (res?.status == 200) {
          dataSource.value = res?.data
          tabLoading.value=false
        }
      })
    },
    onDelete(val) {
      delUser({ id: val }).then(res => {
        if (res?.status == 200) {
          message.success('删除成功')
        } else {
          message.error('删除失败')
        }
      })
    },
    onSubmit() {
      addUser({ ...toRaw(formState).value }).then(res => {
        if (res?.status == 200) {
          message.success('添加用户成功')
          this.handleUserList()
          this.resetFormState()
          visible.value = false
        } else {
          message.error(`${res.message}`)
        }
      }).catch(() => {
        message.success(`添加失败`)
        this.resetFormState()
        visible.value = false
      })
    },
    resetFormState() {
      formState.value = {
        username: '',
        password: '',
        nickname: '',
        perms: '[]',
      };
    }
  }
});

</script>

<style scoped></style>
