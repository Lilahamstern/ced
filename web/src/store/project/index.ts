import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { ProjectState } from './types';
import { RootState } from '../types';
import { Module } from 'vuex';

export const state: ProjectState = {
    project: undefined,
    projects: undefined,
    error: false,
};

const namespaced: boolean = true;

export const project: Module<ProjectState, RootState> = {
    namespaced,
    state,
    getters,
    actions,
    mutations
}
