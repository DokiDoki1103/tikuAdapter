import { createApp } from 'vue'
import App from './App.vue'
import { Alert, Button, Pagination, Table, Tag, Modal,Form ,Input} from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'

const app = createApp(App)
app.use(Button)
app.use(Alert)
app.use(Modal)
app.use(Form)
app.use(Table)
app.use(Input)
app.use(Tag)
app.use(Pagination)
app.mount('#app')
