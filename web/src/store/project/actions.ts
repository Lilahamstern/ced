import {ActionTree} from 'vuex';
import {Project, ProjectState} from './types';
import {RootState} from '../types';
import axios from 'axios';

export const actions: ActionTree<ProjectState, RootState> = {
  async fetchProjects ({ commit }) {
    try {
      var res = await axios.get("http://localhost:5050/projects");
        const payload: Project[] = res && res.data.data
        commit("projectsLoaded", payload)
    } catch (err) {
        commit("projectsError")
    }
  },
    async searchProjects({commit},) {
        try {
            var res = await axios.get("http://localhost:5050/projects?search=");
            const payload: Project[] = res && res.data.data
            commit("projectsLoaded", payload)
    } catch (err) {
            // console.error("Error response:");
            // console.error(err.response.data);
            // console.error(err.response.status);
            // console.error(err.response.headers);
            commit("projectsError")
    }
  },
  selectProject ({ commit, state }, payload: any): any {
    if (!state.projects) return
    commit("selectProject", state.projects[payload.index])
  }
}