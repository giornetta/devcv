<template>
  <div class="section nav has-background-white">
    <div class="container">
      <div class="columns">
        <div class="column is-8 is-offset-2">
          <h1 class="title is-pulled-left">{{ developer.first_name }} {{ developer.last_name }}</h1>

          <div v-if="isOwnAccount && isLoggedIn">
            <div class="menu dropdown is-right is-pulled-right" :class="{'is-active': dropdown}">
              <div class="dropdown-trigger">
                <button class="button is-white has-text-link" aria-haspopup="true" aria-controls="dropdown-menu" @click="dropdown = !dropdown">
                  <span>Menu</span>
                  <span class="icon is-small">
                    <i class="fas fa-angle-down" aria-hidden="true"></i>
                  </span>
                </button>
              </div>
              <div class="dropdown-menu" id="dropdown-menu" role="menu">
                <div class="dropdown-content">
                  <router-link class="dropdown-item" to="/edit">
                    Edit
                  </router-link>
                  <hr class="dropdown-divider">
                  <a class="dropdown-item" @click="logOut">
                    Log Out
                  </a>
                </div>
              </div>
            </div>
          </div>
          <div v-else-if="isLoggedIn">
            <button class="button is-white has-text-link is-pulled-right" aria-haspopup="true" aria-controls="dropdown-menu" @click="onProfile">
              My Profile    
            </button>
          </div>
          <div v-else>
            <button class="button is-white has-text-link is-pulled-right" aria-haspopup="true" aria-controls="dropdown-menu" @click="$emit('register')">
              Log In    
            </button>
          </div>
        </div>
      </div>
    </div>
  </div> 
</template>

<script>
export default {
  name: 'Navbar',
  data: () => ({
    dropdown: false
  }),
  props: {
    developer: {
      type: Object,
      default: () => ({
        //first_name: 'DevCV'
      })
    }
  },
  methods: {
    logOut() {
      localStorage.removeItem('username')
      localStorage.removeItem('token')

      this.dropdown = false

      this.$toasted.show('Goodbye!', {
        duration: 1000,
        position: 'top-center',
        type: 'success',
        singleton: true
      })

      this.$router.push('/')
    },
    onProfile () {
      const username = localStorage.getItem('username')
      this.$router.push(`/${username}`)
    }
  },
  computed: {
    isOwnAccount () {
      return this.developer.username == localStorage.getItem('username')
    },
    isLoggedIn () {
      return (localStorage.getItem('username') && localStorage.getItem('token'))
    }
  }
}
</script>

<style>
.nav {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 999;
  width: 100%;
  /*min-height: 5rem;*/
  max-height: 8rem;
}
</style>
