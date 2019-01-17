<template>
<div>
  <Navbar :developer="developer" @register="onRegister"/>
  
  <div class="page">
    <div class="section">
      <div class="container">
        <div class="columns">
          <div class="column is-8 is-offset-2">

            <!-- INTRODUCTION -->
            <div class="profile-section">
<h2 class="title">Introduction</h2>
            <div class="tags">
              <span class="tag is-light is-uppercase mr">{{ developer.speciality }}</span>
              <span class="tag is-light is-uppercase mr">{{ developer.timezone }}</span>
              <span class="tag is-light is-uppercase mr">{{ developer.city }}</span>
              <span class="tag is-light is-uppercase">{{ developer.languages }}</span>
            </div>
            <div>{{ developer.introduction }}</div>
            </div>
            

            <!-- LINKS -->
            <div class="profile-section">
            <a class="mr" :href="l.url" v-for="(l, i) in developer.links" :key="i">{{ l.title }}</a>
            </div>
            <!-- SKILLS -->
            <div class="profile-section">
            <h2 class="title">Skills</h2>
            <div class="columns is-multiline">
              <div class="column is-half" v-for="(sg, i) in developer.skill_groups" :key="i">
                <table class="table">
                  <thead>
                    <th>{{ sg.title }}</th>
                  </thead>
                  <tbody>
                    <tr v-for="(s, i) in sg.skills" :key="i">
                      <td>{{ s.title }}</td>
                      <td class="is-pulled-left"><leds :level="s.experience"/></td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
            </div>

            <!-- PROJECTS -->
            <div class="profile-section">
            <h2 class="title">Projects</h2>
            </div>
            
            
          </div>
          <div style="height: 5000px;"></div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import Leds from '@/components/Leds'
import Navbar from '@/components/Navbar'

export default {
  name: 'Profile',
  components: {
    'leds': Leds,
    Navbar
  },
  data: () => ({
    developer: {},
    dropdownActive: false
  }),
  methods: {
    onRegister () {
      this.$router.push('/')
    },
    async getDeveloper() {
      const username = this.$route.params.username
      try {
        this.developer = await this.$http.getDeveloper(username)
      } catch(e) {
        this.$toasted.show('Developer not found', {
          duration: 1000,
          position: 'top-center',
          type: 'error',
          singleton: true
        })
        this.$router.push('/')
      }
    }
  },
  watch: {
    '$route': {
      handler: 'getDeveloper',
      immediate: true
    }
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.mr {
  margin-right: 0.5rem;
}

.tag {
  letter-spacing: 1px;
}

.profile-section {
  margin-bottom: 2rem;
}
</style>
