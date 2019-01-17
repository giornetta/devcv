<template>
  <div>
    <Navbar :developer="developer"/>

    <div class="page">
      <div class="section">
        <div class="container">
          <div class="columns">
            <div class="column is-8 is-offset-2">

              <div class="has-background-light form-area field">

                <div class="field">
                  <label class="label">Speciality</label>
                  <div class="control">
                    <input class="input" type="text" v-model="developer.speciality">
                  </div>
                </div>

                <div class="field">
                  <label class="label">Timezone</label>
                  <div class="control">
                    <input class="input" type="text" v-model="developer.timezone">
                  </div>
                </div>

                <div class="field">
                  <label class="label">City</label>
                  <div class="control">
                    <input class="input" type="text" v-model="developer.city">
                  </div>
                </div>

                <div class="field">
                  <label class="label">Languages</label>
                  <div class="control">
                    <input class="input" type="text" v-model="developer.languages">
                  </div>
                </div>

                <div class="field">
                  <label class="label">Introduction Text</label>
                  <div class="control">
                    <textarea class="textarea has-fixed-size" type="text" v-model="developer.introduction"></textarea>
                  </div>
                </div>
              </div>

              <SkillGroupForm 
                :skillGroup="sg"
                v-for="(sg, i) in developer.skill_groups" :key="i"
                @add-skill="addSkill(i)"
                @delete-skill="deleteSkill(i, $event)"
                @delete-group="deleteGroup(i)"
              />

              <div class="field">
                <button class="button" @click="addGroup">Add Skill Card</button>
              </div> 


              <div style="height: 3rem;"></div>
              <div class="actions-bottom">
                <div class="field is-grouped">
                  <div class="control">
                    <button class="button is-link is-cta" @click="onUpdate">Update</button>
                  </div>
                  <div class="control">
                    <router-link class="button is-link is-outlined" :to="{path: `/${developer.username}`}">Cancel</router-link>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar'
import SkillGroupForm from '@/components/SkillGroupForm'

export default {
  name: 'Edit',
  components: {
    Navbar,
    SkillGroupForm
  },
  data: () => ({
    developer: {},
    dropdownActive: false,
   
  }),
  async created() {
    const username = localStorage.getItem('username')
    if (!username) {
        this.$router.push('/')
    }

    this.developer = await this.$http.getDeveloper(username)
    if (!this.developer) {
      this.$router.push('/')
    }
  },
  methods: {
    async onUpdate() {
      try {
        await this.$http.updateDeveloper(this.developer)
        this.$toasted.show('Profile Updated!', {
          duration: 1000,
          position: 'top-center',
          type: 'success',
          singleton: true
        })
        this.$router.push(`/${this.developer.username}`)
      } catch(e) {
        this.$toasted.show('Could not update :(', {
          duration: 1000,
          position: 'top-center',
          type: 'error',
          singleton: true
        })
      }
    },
    deleteSkill(i, j) {
      this.developer.skill_groups[i].skills.splice(j, 1)
    },
    addSkill(i) {
      this.developer.skill_groups[i].skills.push({title: '', experience: 0})
    },
    deleteGroup(i) {
      this.developer.skill_groups.splice(i, 1)
    },
    addGroup() {
      this.developer.skill_groups.push({
        title: `Skill Card #${this.developer.skill_groups.length + 1}`,
        skills: []
      })
    }
  }
}
</script>

<style scoped>
.form-area {
    padding: 15px 15px;
    border-radius: 8px;
}
.actions-bottom {
    position: fixed;
    width: 100%;
    bottom: 0;
    padding: 16px 0px;
    background-color: #fff;
    z-index: 10;
}
</style>
