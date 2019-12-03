import { ActionTree } from 'vuex';
import { ProjectState, Project } from './types';
import { RootState } from '../types';
import axios from 'axios';

export const actions: ActionTree<ProjectState, RootState> = {
  async fetchProjects ({ commit }) {
    try {
      var res = await axios.get("http://localhost:5050/projects");
      const payload: Project = res && res.data.data
      console.log(res.status)
      if (res.status != 200) {
        console.log("heyyy")
      }
      commit("projectLoaded", payload)
    } catch (err) {
      console.error("Error response:");
      console.error(err.response.data);    // ***
      console.error(err.response.status);  // ***
      console.error(err.response.headers); // ***
    }
  },
  selectProject ({ commit, state }, payload: any): any {
    if (!state.projects) return
    commit("selectProject", state.projects[payload.index])
  }
}