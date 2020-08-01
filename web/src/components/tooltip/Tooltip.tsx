import React, { Component } from "react";

interface IProps {
  projectID: string;
}

interface IState {}

export class Tooltip extends Component<IProps, IState> {
  render() {
    const { projectID } = this.props;
    return (
      <div>
        <div>
          <div className="bg-gray-900 opacity-75 text-white text-xs rounded py-1 px-3 right-0 bottom-full truncate">
            {projectID}
          </div>
        </div>
        <div className="absolute right-0 -mt-6 -mr-3 cursor-pointer ">
          <p className="text-base mr-3">
            <i className="fas fa-info-circle"></i>
          </p>
        </div>
      </div>
    );
  }
}

export default Tooltip;
