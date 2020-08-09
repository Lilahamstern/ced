import React, { Component, Fragment } from "react";
import { RouteComponentProps } from "react-router-dom";
import { FetchSingelProjectByID } from "../service/project/project";
import ProjectInformation from "../components/project/selected/ProjectInformation";
interface IState {}

interface IProps {
  id: string;
}
export class ProjectView extends Component<
  RouteComponentProps<IProps>,
  IState
> {
  constructor(props: RouteComponentProps<IProps>) {
    super(props);
  }

  render() {
    return (
      <Fragment>
        <ProjectInformation id={this.props.match.params.id} />
      </Fragment>
    );
  }
}

export default ProjectView;
