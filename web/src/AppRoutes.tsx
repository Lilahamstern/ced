import React, { FunctionComponent } from "react";
import { BrowserRouter, Route } from "react-router-dom";

import Navbar from "./components/navbar/Navbar";
import HomeView from "./views/HomeView";
import { ProjectViewContainer } from "./views/ProjectView";

const AppRoutes: FunctionComponent = () => {
  return (
    <BrowserRouter>
      <div className="bg-gray-900 w-full h-screen">
        <Navbar />
        <Routes />
      </div>
    </BrowserRouter>
  );
};

export default AppRoutes;

const Routes: FunctionComponent = () => {
  return (
    <React.Fragment>
      <Route path="/" exact component={HomeView} />
      <Route path="/project/:id" component={ProjectViewContainer} />
    </React.Fragment>
  );
};
