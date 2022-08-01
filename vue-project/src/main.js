import { createApp } from 'vue'
import App from './App'
import router from '@/router/router'
import store from '@/store'
import 'material-design-icons-iconfont'

const app = createApp(App);

app.use(router);
app.use(store);

app.mount('#app');

