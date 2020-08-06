import React, { Component } from "react";
import Card from "./card/Card";
import { IProject } from "../../../../redux/reducers/ProjectReducer";
import { connect } from "react-redux";
import { IAppState } from "../../../../redux/store/store";

interface IProps {
  projects: IProject[];
}

interface IState {}

export class Cards extends Component<IProps, IState> {
  render() {
    const { projects } = this.props;
    const markup = [];

    if (projects.length > 0) {
      for (let project of projects) {
        markup.push(
          <Card
            project={project}
            className={"flex mt-5 w-full justify-center"}
            key={project.id}
          />
        );
      }
      return <div className="w-full">{markup}</div>;
    }

    return <div>No projects found</div>;
  }
}

const mapStateToProps = (store: IAppState) => {
  return {
    projects: store.projectState.projects,
  };
};

export default connect(mapStateToProps)(Cards);
