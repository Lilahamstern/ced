import { IProject, IProjectState } from "../reducers/ProjectReducer";
import { ThunkAction } from "redux-thunk";
import { ActionCreator, Dispatch } from "redux";
import {
  FetchProjects,
  FetchSingelProjectByID,
} from "../../service/project/project";

export enum ProjectActionTypes {
  FETCH_ALL = "FETCH_ALL",
  SELECT = "SELECT_PROJECT",
}

export interface IProjectFetchAllAction {
  type: ProjectActionTypes.FETCH_ALL;
  projects: IProject[];
}

export interface IProjectSelectAction {
  type: ProjectActionTypes.SELECT;
  project: IProject;
}

export type ProjectActions = IProjectFetchAllAction | IProjectSelectAction;

export const fetchAllProjects: ActionCreator<ThunkAction<
  Promise<any>,
  IProjectState,
  null,
  IProjectFetchAllAction
>> = () => {
  return async (dispatch: Dispatch) => {
    try {
      FetchProjects().then((res) => {
        dispatch({
          type: ProjectActionTypes.FETCH_ALL,
          projects: res.data.projects,
        });
      });
    } catch (err) {
      console.log(err);
    }
  };
};

export const SelectProject: ActionCreator<ThunkAction<
  Promise<any>,
  IProjectState,
  null,
  IProjectSelectAction
>> = (id: string) => {
  return async (dispatch: Dispatch) => {
    try {
      FetchSingelProjectByID(id).then((res) => {
        dispatch({
          type: ProjectActionTypes.SELECT,
          project: res.data.project,
        });
      });
    } catch (err) {
      console.log(err);
    }
  };
};
