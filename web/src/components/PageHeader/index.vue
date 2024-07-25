<template>
  <header>

    <span>
      <ReadOutlined />
      题库管理平台
    </span>
    <a-tooltip placement="bottom" v-if="hasLogout">
      <template #title>
        <span>退出登陆</span>
      </template>
      <span @click="logoutOut">
        <LogoutOutlined />
      </span>
    </a-tooltip>

  </header>
</template>

<script>
import {
  ReadOutlined,
  LogoutOutlined
} from '@ant-design/icons-vue';
import { removeCookies } from '../../utils/cookies.js'
import { useRouter } from 'vue-router';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'PageHeader',
  props: {
    hasLogout: {
      type: Boolean,
      default: false
    }
  },
  setup(props) {
    console.log(props, "props");
    const router = useRouter()
    const navigateToImport = () => {
      router.push('/login');
    };
    return {
      navigateToImport
    }
  },
  components: {
    ReadOutlined,
    LogoutOutlined
  },
  methods: {
    logoutOut() {
      removeCookies('token')
      this.navigateToImport()
    }
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
  text-align: left;
  padding-left: 3%;
  color: #fff;
}

span:nth-of-type(1) {
  margin-right: 6px;
}

span:nth-of-type(2) {
  margin-left: 84%;
}
</style>
