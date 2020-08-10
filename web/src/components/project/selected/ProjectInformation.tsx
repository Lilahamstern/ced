import React, { Component } from "react";
import { IAppState } from "../../../redux/store/store";
import {
  IProject,
  IProjectState,
} from "../../../redux/reducers/ProjectReducer";
import {
  SelectProject,
  IProjectSelectAction,
} from "../../../redux/actions/ProjectActions";
import { connect } from "react-redux";
import { ThunkDispatch } from "redux-thunk";
import { getTimeSince } from "../../../utils/utils";

interface IProps {
  id: string;
  project: IProject;
  fetchProject: (id: string) => Promise<void>;
}

interface IState {}

class ProjectInformation extends Component<IProps, IState> {
  UNSAFE_componentWillMount() {
    this.props.fetchProject(this.props.id);
  }
  render() {
    const { project } = this.props;
    if (!project) {
      return <div>404</div>;
    }
    return (
      <div className="">
        <div className="flex p-5 bg-red-200">
          <ul>
            <li>{project.id}</li>
            <li>{getTimeSince(project.updated_at)}</li>
            <li>{project.versions.length}</li>
            <li>{}</li>
            <li></li>
          </ul>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (store: IAppState) => {
  return {
    project: store.projectState.project,
  };
};

const mapDispatchToProps = (
  dispatch: ThunkDispatch<IProjectState, null, IProjectSelectAction>
) => {
  return {
    fetchProject: (id: string) => dispatch(SelectProject(id)),
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(ProjectInformation);
