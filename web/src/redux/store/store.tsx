import { IProjectState, projectReducer } from "../reducers/ProjectReducer";
import { combineReducers, Store, createStore, applyMiddleware } from "redux";
import thunk from "redux-thunk";

export interface IAppState {
  projectState: IProjectState;
}

const rootReducer = combineReducers<IAppState>({
  projectState: projectReducer,
});

export default function configureStore(): Store<IAppState, any> {
  const store = createStore(rootReducer, undefined, applyMiddleware(thunk));
  return store;
}
