import React, { Component } from "react";
import Card from "./card/Card";
import { IProject } from "../../../../redux/reducers/ProjectReducer";

interface IProps {
  projects: IProject[];
}

interface IState {}

export class Cards extends Component<IProps, IState> {
  render() {
    const { projects } = this.props;
    if (projects.length > 0) {
      return (
        <div className="w-full">
          {projects.map((project) => {
            return (
              <Card
                project={project}
                className={"flex mt-5 w-full justify-center"}
                key={project.id}
              />
            );
          })}
        </div>
      );
    }

    return <div>No projects found</div>;
  }
}

export default Cards;
