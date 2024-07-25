<template>
  <div style="width: 100%">
    <a-menu
      mode="inline"
      theme="light"
      v-model:selectedKeys="selectedKeys"
      @select="onclick"
    >
      <a-menu-item key="component">
        <template #icon>
          <PieChartOutlined />
        </template>
        <span>题库总览</span>
      </a-menu-item>
      <a-menu-item key="import">
        <template #icon>
          <DownloadOutlined />
        </template>
        <span>智能导入</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <template #icon>
          <TeamOutlined />
        </template>
        <span>用户列表</span>
      </a-menu-item>
      <a-menu-item key="loglist">
        <template #icon>
          <FileTextOutlined />
        </template>
        <span>日志列表</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script>
import { defineComponent, ref,  watch, onMounted } from 'vue';
import {useRouter, useRoute} from 'vue-router';
import {
  PieChartOutlined,
  DownloadOutlined,
  TeamOutlined,
  FileTextOutlined 
} from '@ant-design/icons-vue';
const selectedKeys = ref([])
export default defineComponent({
  name: 'SiderMenu',
  setup() {
    const router = useRouter();
    const route = useRoute();
    const selectedKeys = ref([route.name]);
    const navigateToImport = (key) => {
      router.push(`/adapter/${key}`);
    };
    watch(() => route.path, (newPath) => {
      const pathSegments = newPath.split('/');
      selectedKeys.value = [pathSegments[pathSegments.length - 1]];
    });

    onMounted(() => {
      const pathSegments = route.path.split('/');
      selectedKeys.value = [pathSegments[pathSegments.length - 1]];
    });

    return {
      selectedKeys,
      navigateToImport
    };
  },
  methods:{
    onclick (data) {
      selectedKeys.value=[data.key]
      this.navigateToImport(data.key)
    },
  },
  components: {
    PieChartOutlined,
    DownloadOutlined,
    TeamOutlined,
    FileTextOutlined
  },
});
</script>

