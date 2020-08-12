import React, { Component } from "react";
import { RouteComponentProps } from "react-router-dom";
import ProjectInformation from "../components/project/selected/ProjectInformation";
import VersionList from "../components/version/list/List";
import { ThunkDispatch } from "redux-thunk";
import { IProjectState } from "../redux/reducers/ProjectReducer";
import {
  IProjectSelectAction,
  SelectProject,
} from "../redux/actions/ProjectActions";
import { connect } from "react-redux";
interface IState {}

interface IRouterProps {
  id: string;
}

interface IProps {
  fetchProject: (id: string) => Promise<void>;
}
export default class ProjectView extends Component<
  RouteComponentProps<IRouterProps> & IProps,
  IState
> {
  UNSAFE_componentWillMount() {
    this.props.fetchProject(this.props.match.params.id);
  }
  render() {
    return (
      <div className="mt-10">
        <ProjectInformation />
        <VersionList />
      </div>
    );
  }
}

const mapDispatchToProps = (
  dispatch: ThunkDispatch<IProjectState, null, IProjectSelectAction>
) => {
  console.log("Fetching");
  return {
    fetchProject: (id: string) => dispatch(SelectProject(id)),
  };
};

export const ProjectViewContainer = connect(
  null,
  mapDispatchToProps
)(ProjectView);
