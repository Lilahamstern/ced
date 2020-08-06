import { IProject, IProjectState } from "../reducers/ProjectReducer";
import { ThunkAction } from "redux-thunk";
import { ActionCreator, Dispatch } from "redux";
import { FetchProjects } from "../../service/project/project";

export enum ProjectActionTypes {
  FETCH_ALL = "FETCH_ALL",
}

export interface IProjectFetchAllAction {
  type: ProjectActionTypes.FETCH_ALL;
  projects: IProject[];
}

export type ProjectActions = IProjectFetchAllAction;

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
