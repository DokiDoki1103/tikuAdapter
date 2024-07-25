<template>
  <Header/>
  <main>
    <router-view></router-view>
  </main>
  <Footer/>
</template>

<script>
import Header from '../components/PageHeader'
import Footer from '../components/PageFoot'
export default {
  name: 'LoginLayout',
  components:{
    Header,
    Footer
  }
}
</script>

<style scoped>
main{
  width: 100%;
  height: calc(100% - 120px);
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
