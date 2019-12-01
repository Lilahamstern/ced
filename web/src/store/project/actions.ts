import { ActionTree } from 'vuex';
import { ProjectState, Project } from './types';
import { RootState } from '../types';
import axios from 'axios';

export const actions: ActionTree<ProjectState, RootState> = {
  fetchProjects ({ commit }): any {
    axios({
      url: "http://localhost:5050/projects"
    }).then(res => {
      const payload: Project = res && res.data.data
      console.log(res.status)
      if (res.status != 200) {
        console.log("heyyy")
      }
      commit("projectLoaded", payload)
    }).catch(error => {
      console.log(error)
      commit("projectError")
    })
  },
  selectProject ({ commit, state}, payload: any): any {
    if (!state.projects) return
    commit("selectProject", state.projects[payload.index])
  }
}