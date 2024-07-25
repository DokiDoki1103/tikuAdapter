<template>
  <div class="container">
    <div class="heading">用户登陆</div>
    <a-form layout="vertical" :model="formState" @finish="handleFinish" @finishFailed="handleFinishFailed" class="form">
      <a-form-item>
        <a-input v-model:value="formState.username" placeholder="Username" class="input">
          <template #prefix>
            <UserOutlined style="color: rgba(0, 0, 0, 0.25)" />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-input v-model:value="formState.password" type="password" placeholder="Password" class="input">
          <template #prefix>
            <LockOutlined style="color: rgba(0, 0, 0, 0.25)" />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit" :disabled="formState.username === '' || formState.password === ''"
          class="login-button">
          登 陆
        </a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script>
import { defineComponent, reactive } from 'vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
import { login } from '../../api/TikuComponentApi.js'
import { setCookies } from '../../utils/cookies.js'
import { message } from "ant-design-vue";
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'AdapterLogin',
  setup() {
    const formState = reactive({
      username: '',
      password: '',
    });
    const router = useRouter()
    const handleFinish = () => {
      login(formState).then((res) => {  
        if(res?.status == 200){
          message.success('登陆成功')
          setCookies('token',res?.data?.jwt)
          router.push('/adapter/component')
        }else{
          message.error(`${res?.message}`)
        }
      })
    };
    const handleFinishFailed = errors => {
      console.log(errors);
    };
    return {
      formState,
      handleFinish,
      handleFinishFailed,
    };
  },
  components: {
    UserOutlined,
    LockOutlined,
  },
  methods: {

  }
})
</script>

<style scoped>
.container {
  min-width: 400px;
  background: #F8F9FD;
  background: linear-gradient(0deg, rgb(255, 255, 255) 0%, rgb(244, 247, 251) 100%);
  border-radius: 40px;
  padding: 25px 35px;
  border: 5px solid rgb(255, 255, 255);
  box-shadow: rgba(133, 189, 215, 0.8784313725) 0px 30px 30px -20px;
}

.heading {
  text-align: center;
  font-weight: 900;
  font-size: 30px;
  color: rgb(16, 137, 211);
}

.form {
  margin-top: 20px;
}

.form .input {
  width: 100%;
  background: white;
  border: none;
  padding: 15px 20px;
  border-radius: 20px;
  margin-top: 15px;
  box-shadow: #cff0ff 0px 10px 10px -5px;
  border-inline: 2px solid transparent;
}

.form .input::-moz-placeholder {
  color: rgb(170, 170, 170);
}

.form .input::placeholder {
  color: rgb(170, 170, 170);
}

.form .input:focus {
  outline: none;
  border-inline: 2px solid #12B1D1;
}

.form .forgot-password {
  display: block;
  margin-top: 10px;
  margin-left: 10px;
}

.form .forgot-password a {
  font-size: 11px;
  color: #0099ff;
  text-decoration: none;
}

.form .login-button {
  height: 50px !important;
  display: block;
  width: 100%;
  font-weight: bold;
  /* background: linear-gradient(45deg, rgb(16, 137, 211) 0%, #195ca3 100%); */
  /* color: #00000040; */
  padding-block: 15px;
  margin: 20px auto;
  border-radius: 20px;
  box-shadow: rgba(133, 189, 215, 0.8784313725) 0px 20px 10px -15px;
  border: none;
  transition: all 0.2s ease-in-out;
}

.form .login-button:hover {
  transform: scale(1.03);
  box-shadow: rgba(133, 189, 215, 0.8784313725) 0px 23px 10px -20px;
}

.form .login-button:active {
  transform: scale(0.95);
  box-shadow: rgba(133, 189, 215, 0.8784313725) 0px 15px 10px -10px;
}
</style>