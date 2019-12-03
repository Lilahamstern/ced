import {MutationTree} from 'vuex';
import {Project, ProjectState} from './types';

export const mutations: MutationTree<ProjectState> = {
    projectsLoaded(state, payload: Project[]) {
    state.error = false;
    state.projects = payload;
  },
    projectsError(state) {
    state.error = true;
        state.projects = [];
  },
  selectProject(state, payload: Project) {
    state.error = false;
    state.project = payload;
  }
}