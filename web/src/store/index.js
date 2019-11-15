import Vue from "vue";
import Vuex from "vuex";
import projects from "./modules/projectModule"

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    projects: projects
  }
});
