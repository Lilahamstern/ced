import React, { Component } from "react";

import { ProjectViewStatus } from "../../../../enums";

interface IProps {
  view: ProjectViewStatus;
}

interface IState {}

export class Controller extends Component<IProps, IState> {
  render() {
    const { view } = this.props;

    return (
      <div>
        <p>{view}</p>
      </div>
    );
  }
}

export default Controller;
