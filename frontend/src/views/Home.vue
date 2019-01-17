<template>
  <div>
    <!--
    <div class="menu">
        <a :href="username" class="has-text-link" v-if="username">My Profile</a>
        <a :href="username" class="has-text-link" v-else @click="onClick">Register</a>
    </div>
    -->

    <Navbar @register="onClick"/>

    <div class="hero is-fullheight">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title">Simple CVs! for Developers</h1>
          
          <button class="button is-large has-shadow is-link" @click="onClick" v-if="!showForm">Get started</button>
          <login-form v-else/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Login from '@/components/LoginForm.vue'
import Navbar from '@/components/Navbar'
export default {
  name: 'home',
  components: {
    'login-form': Login,
    Navbar: Navbar
  },
  data: () => ({
    showForm: false,
    username: ''
  }),
  methods: {
    onClick() {
      this.showForm = !this.showForm
    }
  },
  created() {
    this.username = localStorage.getItem("username")
  }
}
</script>

<style scoped>
  .menu {
    position: absolute;
    top: 1.5rem;
    right: 2rem;
  }
</style>

