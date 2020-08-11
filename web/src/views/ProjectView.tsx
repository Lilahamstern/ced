import React, { Component } from "react";
import { RouteComponentProps } from "react-router-dom";
import ProjectInformation from "../components/project/selected/ProjectInformation";
interface IState {}

interface IProps {
  id: string;
}
export class ProjectView extends Component<
  RouteComponentProps<IProps>,
  IState
> {
  render() {
    return (
      <div className="mt-10">
        <ProjectInformation id={this.props.match.params.id} />
      </div>
    );
  }
}

export default ProjectView;
