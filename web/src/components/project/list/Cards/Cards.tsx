import React, { Component } from "react";
import { Projects } from "../../../../stores/types";
import Card from "./Card/Card";

interface IProps {
  projects: Projects;
}

interface IState {}

export class Cards extends Component<IProps, IState> {
  render() {
    const { projects } = this.props;
    const markup = [];

    for (let project of projects) {
      markup.push(
        <Card
          project={project}
          className={"flex mt-5 w-full justify-center"}
          key={project.ID}
        />
      );
    }
    return <div className="w-full">{markup}</div>;
  }
}

export default Cards;
