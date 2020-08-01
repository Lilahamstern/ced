import React, { Component } from "react";
import { Project } from "../../../../../stores/types";
import Tooltip from "../../../../tooltip/Tooltip";

interface IProps {
  project: Project;
  className?: string;
}

interface IState {}

export class Card extends Component<IProps, IState> {
  render() {
    const { project, className } = this.props;
    return (
      <div className={className}>
        <div className="bg-gray-800 shadow shadow-t rounded max-w-md md:max-w-2xl lg:max-w-4xl xl:max-w-6xl w-full">
          <div className="flex pl-4 items-center h-32 md:h-24">
            <div className="w-full text-gray-600 relative">
              <div className="card-row">
                <h2 className="text-gray-400 text-md font-semibold truncate -mt-1">
                  Project:{" "}
                  <span className="text-gray-500">{project.OrderID}</span>
                </h2>
              </div>
              <div className="card-row mt-10">
                <p className="info">
                  <span className="label">Title:</span> {project.Title}
                </p>
                <p className="info">
                  <span className="label">Sector:</span> {project.Sector}
                </p>
                <p className="info">
                  <span className="label">Manager:</span> {project.Manager}
                </p>
                <p className="info">
                  <span className="label">Client:</span> {project.Client}
                </p>
                <p className="info">
                  <span className="label">Co2-ekv/m2:</span> {project.Co}
                </p>
              </div>
              <div className="tooltip-card">
                <Tooltip projectID={project.ID} />
              </div>
              <div className="absolute right-0 top-0">
                <small className="text-sm">
                  <span className="pr-1">
                    <i className="far fa-clock"></i>
                  </span>
                  {project.UpdatedAt}d ago
                </small>
              </div>
            </div>
            <div className="flex items-center text-lg text-gray-500 hover:text-gray-600 hover-slide rounded-r h-full pl-1 pr-2 ml-1 cursor-pointer mx-auto">
              <i className="fas fa-arrow-right"></i>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default Card;
