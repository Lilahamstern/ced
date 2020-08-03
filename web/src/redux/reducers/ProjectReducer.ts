import { IVersion } from "./VersionReducer";
import { Reducer } from "react";
import { ProjectActionTypes, ProjectActions } from "../actions/ProjectActions";

export interface IProject {
  ID: number;
  versions: IVersion[];
  CreatedAt: string;
  UpdatedAt: string;
}

export interface IProjectState {
  projects: IProject[];
}

const initialProjectState: IProjectState = {
  projects: [],
};

export const projectReducer: Reducer<IProjectState, ProjectActions> = (
  state = initialProjectState,
  action
) => {
  switch (action.type) {
    case ProjectActionTypes.FETCH_ALL: {
      return {
        ...state,
        projects: action.projects,
      };
    }
    default:
      return state;
  }
};
