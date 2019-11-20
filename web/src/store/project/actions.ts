import { ActionTree } from 'vuex';
import { ProjectState, Project } from './types';
import { RootState } from '../types';
import axios from 'axios';

export const actions: ActionTree<ProjectState, RootState> = {
  fetchData ({ commit }): any {
    axios({
      url: "www.google.com"
    }).then((res) => {
      const payload: Project = res && res.data
      commit("projectLoaded", payload)
    }), (error: any) => {
      console.log(error)
      commit("projectError")
    }
  }
}