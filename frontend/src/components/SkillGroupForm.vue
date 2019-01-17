<template>
  <div class="has-background-light form-area field">
    <div>
      <h1 class="title is-5 is-pulled-left">{{ skillGroup.title }}</h1>
      <a class="has-text-danger is-pulled-right" @click="$emit('delete-group')">
        <span class="icon is-medium">
          <i class="far fa-trash-alt"></i>                        
        </span>
      </a>
      <div class="is-clearfix"></div>
    </div>

    <div class="field">
      <label class="label">Title</label>
      <div class="control">
        <input class="input" type="text" v-model="skillGroup.title">
      </div>
    </div>

    <div class="field">
      <label class="label">Skills</label>

      <div class="field is-grouped" v-for="(skill, j) in skillGroup.skills" :key="j">
        <div class="control">
          <input type="text" class="input" v-model="skill.title">                          
        </div>
        <div class="control">
          <div class="select">
            <select v-model="skill.experience">
              <option v-for="(option, k) in skillExpOptions" :value="option.value" :key="k">
                {{ option.text }}
              </option>
            </select>
          </div>
        </div>
        <div class="control">
          <a class="has-text-danger" @click="$emit('delete-skill', j)">
            <span class="icon is-medium">
              <i class="far fa-trash-alt"></i>                        
            </span>
          </a>
        </div>
      </div>
    </div>

    <div class="field">
      <button class="button" @click="$emit('add-skill')">Add Skill</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SkillCardForm',
  data: () => ({
    skillExpOptions: [
      { text: 'Prefer to not say' },
      { text: 'One Year', value: 1 },
      { text: 'Two Years', value: 2 },
      { text: 'Three Years', value: 3 },
      { text: 'Four Years', value: 4 },
    ]
  }),
  props: {
    skillGroup: Object
  }
}
</script>

<style>

</style>
