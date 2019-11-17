const getters = {
  PROJECTS: state => {
    return state.projects;
  },
  GET_STATUS: state => {
    return state.loading;
  }
};

const actions = {
  GET_PROJECTS: context => {
    context.commit("GET_PROJECTS");
  },
  SEARCH_PROJECT: (context, payload) => {
    context.commit("GET_PROJECTS", payload);
  }
};

const mutations = {
  // GET_PROJECTS: ({ state }) => {
  //   // TODO: Fetch all projects.
  // },
  SEARCH_PROJECT: ({state}, payload) => {
    state.loading = true;
    console.log(payload)
    setTimeout(() => {
      state.loading = false;
    }, 5000)
  },
  SET_PROJECT: ({ state }, payload) => {
    state.projects = payload;
  }
};

export default {
  namespaced: true,
  state: {
    projects: [{ name: "test" }, { name: "apa" }],
    loading: false,
    error: null
  },
  getters: getters,
  mutations: actions,
  actions: mutations
};
