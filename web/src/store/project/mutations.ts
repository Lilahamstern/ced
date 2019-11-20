import { MutationTree } from 'vuex';
import { ProjectState, Project } from './types';

export const mutations: MutationTree<ProjectState> = {
  projectLoaded (state, payload: Project) {
    state.error = false;
    state.project = payload;
  },
  projectError (state) {
    state.error = true;
    state.project = undefined;
  }
}