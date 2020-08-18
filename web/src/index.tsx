import React, { FC } from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { Store } from "redux";
import * as serviceWorker from "./serviceWorker";

import configureStore, { IAppState } from "./redux/store/store";

import "./styles/main.css";
import AppRoutes from "./AppRoutes";

// Register fontawesome
import "./fontawesome";
import { fetchAllProjects } from "./redux/actions/ProjectActions";

interface IProps {
  store: Store<IAppState>;
}

const Root: FC<IProps> = (props) => {
  return (
    <Provider store={props.store}>
      <AppRoutes />
    </Provider>
  );
};

const store = configureStore();
store.dispatch(fetchAllProjects());

ReactDOM.render(
  <Root store={store} />,
  document.getElementById("root") as HTMLElement
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
