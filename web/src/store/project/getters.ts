import { GetterTree } from 'vuex';
import { ProjectState, Project } from './types';
import { RootState } from '../types';
import { project } from '.';

export const getters: GetterTree<ProjectState, RootState> = {
  projectInfo (state): Project | undefined {
    const { project } = state;
    return project
  },
  projects (state): Project[] {
    const { projects, error } = state;
    if (error || !projects) {
      return []
    }

    return projects
  }
}