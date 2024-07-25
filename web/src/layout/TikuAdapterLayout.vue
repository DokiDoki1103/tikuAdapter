<template>
  <Header :hasLogout="true"/>
  <main>
    <div class="Sider">
      <SiderMenu />
    </div>
    <div class="Content">
      <router-view></router-view>
    </div>
  </main>
</template>

<script>
import Header from '../components/PageHeader'
import SiderMenu from '../components/SiderMenu'
import { useRouter } from 'vue-router';
import { getCookies } from '../utils/cookies'

export default {
  name: 'AdapterLayout',
  setup() {
    const router = useRouter()
    const token = getCookies('token')
    if(!token){
      router.push('/login')
    }
  },
  components: {
    Header,
    SiderMenu
  }
}
</script>

<style scoped>
main {
  height: calc(100% - 60px);
  width: 100%;
  display: flex;
}

.Sider {
  width: 10%;
  background-color: #fff;
}

.Content {
  width: 90%;
  padding: 24px;
  overflow: hidden;
  overflow-y: auto;
}
</style>
