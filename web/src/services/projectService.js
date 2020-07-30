import api from "./api"

export default {
  fetchProjects() {
    return api().get('/projects')
  }
}