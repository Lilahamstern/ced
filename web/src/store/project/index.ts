import { ProjectState } from './types';
import { Module } from 'vuex';
import { RootState } from '../types';
import { actions } from './actions'
import { getters } from './getters'
import { mutations } from './mutations'

export const state: ProjectState = {
  project: undefined,
  projects: [{ id: 1234, "client": "Volvo", co2: 2312, name: "Evakuerings byggnad", sector: "Hälso och Sjukvård", state: "Complete" }, { id: 1211, "client": "Volvos", co2: 2312, name: "Evakuerings byggnad", sector: "Hälso och Sjukvård", state: "In Progress" }],
  error: false,
}

const namespaced: boolean = true;

export const project: Module<ProjectState, RootState> = {
  namespaced,
  state,
  actions,
  mutations,
  getters
}