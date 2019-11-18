import { GetterTree } from 'vuex';
import { ProjectState } from './types';
import { RootState } from '../types';
import { Project } from '@/types';

export const getters: GetterTree<ProjectState, RootState> = {
    projectInfo(state): string {
        const { project } = state;
        const projectName = (project && project.name) || '';
        const projectClient = (project && project.client) || '';
        const projectSector = (project && project.sector) || '';

        return `${projectClient} have project: ${projectName}. In sector ${projectSector}`
    }
}