import { IProject, IProjectState } from "../reducers/ProjectReducer";
import { ThunkAction } from "redux-thunk";
import { ActionCreator, Dispatch } from "redux";

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
      const res: IProject[] = [
        {
          ID: 2919023,
          versions: [
            {
              ID: "324df635-ab3b-4601-97fd-d404aefdfas3",
              OrderID: 29192342,
              Description: "Health care hospital",
              Title: "Hospital building",
              Manager: "Sara DOe",
              Client: "Vastra Gotland",
              Sector: "Health care",
              Co: 401829,
              CreatedAt: "5",
              UpdatedAt: "5",
            },
          ],
          CreatedAt: "10d",
          UpdatedAt: "10d",
        },

        {
          ID: 2919,
          versions: [
            {
              ID: "324df635-ab3b-4601-97fd-d404aefdfa43",
              OrderID: 29192378,
              Description: "Something for someone",
              Title: "Equlation building",
              Manager: "John DOe",
              Client: "Klara Doe",
              Sector: "Health",
              Co: 20192,
              CreatedAt: "3",
              UpdatedAt: "3",
            },
          ],
          CreatedAt: "10d",
          UpdatedAt: "10d",
        },
      ];

      dispatch({
        type: ProjectActionTypes.FETCH_ALL,
        projects: res,
      });
    } catch (err) {
      console.log(err);
    }
  };
};
