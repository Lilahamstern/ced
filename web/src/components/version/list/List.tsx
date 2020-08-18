import React, { Component } from "react";
import { IAppState } from "../../../redux/store/store";
import { connect } from "react-redux";
import { IVersion } from "../../../redux/reducers/VersionReducer";

interface IProps {
  versions: IVersion[];
}

interface IState {}

export class List extends Component<IProps, IState> {
  render() {
    const { versions } = this.props;
    if (!versions) {
      return <div>404</div>;
    }
    return <div>Hey</div>;
  }
}

const mapStateToProps = (store: IAppState) => {
  console.log(store.projectState);
  return {
    versions: store.projectState.project.versions,
  };
};

export default connect(mapStateToProps)(List);
