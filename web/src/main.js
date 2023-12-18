import { createApp } from 'vue'
import App from './App.vue'
import {Alert, Button, Pagination, Table, Tag, Modal, Form, Input, Select, Checkbox, Layout, Row, Col, Tooltip, Card, Space, Upload, Spin} from 'ant-design-vue'
import router from './router/router'
import 'ant-design-vue/dist/antd.css'


const app = createApp(App)
app.use(Button)
app.use(Alert)
app.use(Modal)
app.use(Form)
app.use(Table)
app.use(Input)
app.use(Tag)
app.use(Select)
app.use(Pagination)
app.use(Checkbox)
app.use(Layout)
app.use(Row)
app.use(Col)
app.use(Tooltip)
app.use(Card)
app.use(Space)
app.use(Upload)
app.use(Spin)
app.use(router);
app.mount('#app')
