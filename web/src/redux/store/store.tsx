import { IProjectState, projectReducer } from "../reducers/ProjectReducer";
import { combineReducers, Store, createStore, applyMiddleware } from "redux";
import thunk from "redux-thunk";
import { composeWithDevTools } from "remote-redux-devtools";

export interface IAppState {
  projectState: IProjectState;
}

const composeEnhancers = composeWithDevTools({
  realtime: true,
  name: "CED Redux Devtools",
  hostname: "localhost",
  port: 3000,
});

const rootReducer = combineReducers<IAppState>({
  projectState: projectReducer,
});

export default function configureStore(): Store<IAppState, any> {
  const store = createStore(
    rootReducer,
    undefined,
    composeEnhancers(applyMiddleware(thunk))
  );
  return store;
}
