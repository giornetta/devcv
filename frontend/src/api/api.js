import axios from 'axios'

const baseURL = '/api'
const api = axios.create({
    baseURL
})

export default {
  async login(username, password) {
    const res = await api.post('/login', {
      username,
      password
    })

    localStorage.setItem("token", res.data.token)
    localStorage.setItem("username", username)
  },

  async register(username, firstName, lastName, password) {
    const res = await api.post('/developers', {
      username,
      first_name: firstName,
      last_name: lastName,
      password
    })

    localStorage.setItem("token", res.data.token)
    localStorage.setItem("username", username)
  },

  async getDeveloper(username) {
    const res = await api.get(`/developers/${username}`)
    let dev = res.data
    dev.skill_groups[0].skills = dev.skill_groups[0].skills || []
    return dev
  },

  async updateDeveloper(developer) {
    const token = localStorage.getItem("token")
    await api.put(`/developers/${developer.username}`, developer, {
      headers: {
        Authorization: token
      }
    })
  },
}