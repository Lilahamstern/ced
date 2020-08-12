import React, { Component } from "react";
import { IAppState } from "../../../redux/store/store";
import { IProject } from "../../../redux/reducers/ProjectReducer";
import { connect } from "react-redux";
import { getTimeSince } from "../../../utils/utils";

interface IProps {
  project: IProject;
}

interface IState {}

class ProjectInformation extends Component<IProps, IState> {
  render() {
    const { project } = this.props;
    if (!project) {
      return <div>404</div>;
    }
    return (
      <div className="">
        <div className="flex p-5 bg-gray-700 text-white">
          <ul>
            <li>Project ID: {project.id}</li>
            <li>Created at: {getTimeSince(project.created_at, false)}</li>
            <li>
              Last updated at:{" "}
              {getTimeSince(project.versions[0].updated_at, false)}
            </li>
            <li>Total projects: {project.versions.length}</li>
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

export default connect(mapStateToProps, null)(ProjectInformation);
