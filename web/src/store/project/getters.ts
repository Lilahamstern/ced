import { GetterTree } from 'vuex';
import { ProjectState, Project } from './types';
import { RootState } from '../types';

export const getters: GetterTree<ProjectState, RootState> = {
  projectInfo (state): Project | undefined {
    const { project } = state;
    return project
  }
}