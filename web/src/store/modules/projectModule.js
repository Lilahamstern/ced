const getters = {
  PROJECTS: state => {
    return state.projects;
  }
}
const mutations = {
  GET_PROJECTS: (state) => {
    state.loading = true;
  },
  GET_PROJECTS_SUCCESS: (state, projects) => {
    state.projects = projects;
    state.loading = false;
  },
  GET_PROJECTS_FAIL: (state, error) => {
    state.error = error;
    state.loading = false;
  }
}

export default {
  namespaced: true,
  state: {
    projects: [],
    loading: false,
    error: null,
  },
  getters: getters,
  mutations: actions,
  actions: mutations,
}