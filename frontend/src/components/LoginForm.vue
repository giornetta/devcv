<template>
  <div class="section">
    <div class="container">
      <div>
        <div class="field">
          <input class="input" type="text" placeholder="Username" v-model="username">
            <p class="help is-danger">{{ errors.username }}</p>
        </div>

        <div v-if="register">

          <div class="field is-grouped is-grouped-centered">
            <div class="control is-expanded">
              <input type="text" class="input" placeholder="First Name" v-model="firstName">
              <p class="help is-danger">{{ errors.firstName }}</p>
            </div>

            <div class="control is-expanded">
              <input type="text" class="input" placeholder="Last Name" v-model="lastName">
              <p class="help is-danger">{{ errors.lastName }}</p>
            </div>
          </div>
        </div>

        <div class="field">
          <input type="password" class="input" placeholder="Password" v-model="password">
          <p class="help is-danger">{{ errors.password }}</p>
        </div>
    
        <div class="field is-grouped is-grouped-centered">
          <div class="control">
            <button class="button is-link" @click="register ? submitRegister() : submitLogin()">
              {{register ? 'Register' : 'Login'}}
            </button>
          </div>
          <div class="control">
            <button class="button is-text" @click="register = !register">
              {{register ? 'Login' : 'Register'}}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data: () => ({
    register: false,
    errors: {},
    username: '',
    email: '',
    firstName: '',
    lastName: '',
    password: ''
  }),
  methods: {
    async submitLogin() {
      try {
        await this.$http.login(this.username, this.password)

        this.$toasted.show('Welcome back!', {
          duration: 1000,
          position: 'top-center',
          type: 'success',
          singleton: true
        })

        this.$router.push(`/${this.username}`)
      } catch(e) {
        this.$toasted.show('Invalid Username and Password combination', {
          duration: 1000,
          position: 'top-center',
          type: 'error',
          singleton: true
        })
      }
    },
    async submitRegister() {
      if(!this.validate()) {
        return
      }

      try {
        await this.$http.register(this.username, this.firstName, this.lastName, this.password)

        this.$toasted.show("Welcome! Let's start editing your CV!", {
            duration: 1000,
            position: 'top-center',
            type: 'error',
            singleton: true
        })
        this.$router.push('/edit')
      } catch(e) {
        this.$toasted.show('Invalid credentials.', {
          duration: 1000,
          position: 'top-center',
          type: 'error',
          singleton: true
        })
      }
    },
    validate() {
      let valid = true
      this.errors = {}

      if (this.username.length < 6) {
        this.errors.username = 'Username is required (At least 5 characters)'
        valid = false
      }

      if (this.firstName.length <= 0) {
        this.errors.firstName = 'First Name is required'
        valid = false
      }

      if (this.lastName.length <= 0) {
        this.errors.lastName = 'Last Name is required'
        valid = false
      }

      if (this.password.length < 8) {
        this.errors.password = 'Password is required (At least 8 characters)'
        valid = false
      }
      return valid
    }
  }
}
</script>
