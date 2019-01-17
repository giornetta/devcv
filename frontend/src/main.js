import Vue from 'vue'
import App from './App.vue'

import router from './router'
import api from '@/api/api.js'
import Toasted from 'vue-toasted'


Vue.prototype.$http = api

Vue.use(Toasted)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
