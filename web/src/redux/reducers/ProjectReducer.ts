import { IVersion } from "./VersionReducer";
import { Reducer } from "react";
import { ProjectActionTypes, ProjectActions } from "../actions/ProjectActions";

export interface IProject {
  id: number;
  versions: IVersion[];
  created_at: string;
  updated_at: string;
}

export interface IProjectState {
  projects: IProject[];
  project: IProject;
}

const initialProjectState: IProjectState = {
  projects: [],
  project: null,
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
    case ProjectActionTypes.SELECT: {
      return {
        ...state,
        project: action.project,
      };
    }
    default:
      return state;
  }
};
