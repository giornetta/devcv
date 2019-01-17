import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/views/Home'
import Profile from '@/views/Profile'
import Edit from '@/views/Edit'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/edit',
      name: 'edit',
      component: Edit
    },
    {
      path: '/:username',
      name: 'profile',
      component: Profile
    },
  ],
  scrollBehavior () {
    return { x: 0, y: 0 }
  }
})
