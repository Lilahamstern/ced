import React, { Component } from "react";
import { RouteComponentProps } from "react-router-dom";

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
    console.log(props);
  }
  
  testss = () => {
    console.log(this.props);
  };
  render() {
    const { match } = this.props;
    this.testss();
    return <div>{match.params.id}</div>;
  }
}

export default ProjectView;
