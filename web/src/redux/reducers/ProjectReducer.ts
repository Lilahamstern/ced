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
  selected: string;
}

const initialProjectState: IProjectState = {
  projects: [],
  selected: "",
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
        selected: action.selected,
      };
    }
    default:
      return state;
  }
};
