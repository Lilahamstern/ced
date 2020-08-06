import React, { Component } from "react";
import Tooltip from "../../../../tooltip/Tooltip";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IProject } from "../../../../../redux/reducers/ProjectReducer";
import { getTimeSince } from "../../../../../utils/utils";

interface IProps {
  project: IProject;
  className?: string;
}

interface IState {}

export class Card extends Component<IProps, IState> {
  render() {
    const { project, className } = this.props;
    return (
      <div className={className}>
        <div className="bg-gray-800 shadow-md rounded max-w-md md:max-w-2xl lg:max-w-4xl xl:max-w-6xl w-full">
          <div className="flex pl-4 items-center h-32 md:h-24">
            <div className="w-full text-gray-600 relative">
              <div className="card-row">
                <h2 className="text-gray-400 text-md font-semibold truncate -mt-1">
                  Project: <span className="text-gray-500">{project.id}</span>
                </h2>
              </div>
              <div className="card-row mt-10">
                <p className="info">
                  <span className="label">Title:</span>{" "}
                  {project.versions[0].title}
                </p>
                <p className="info">
                  <span className="label">Sector:</span>{" "}
                  {project.versions[0].sector}
                </p>
                <p className="info">
                  <span className="label">Manager:</span>{" "}
                  {project.versions[0].manager}
                </p>
                <p className="info">
                  <span className="label">Client:</span>{" "}
                  {project.versions[0].client}
                </p>
                <p className="info">
                  <span className="label">Co2-ekv/m2:</span>{" "}
                  {project.versions[0].co}
                </p>
              </div>
              <Tooltip
                tooltipClassName="tooltip-card"
                iconClassName="tooltip-card-icon"
              >
                {project.versions[0].id}
              </Tooltip>
              <div className="absolute right-0 top-0">
                <small className="text-sm">
                  <span>
                    <i className="far fa-clock"></i>
                  </span>
                  {getTimeSince(project.versions[0].updated_at)} ago
                </small>
              </div>
            </div>
            <div className="flex items-center text-lg text-gray-500 hover:text-gray-600 hover-slide rounded-r h-full pl-2 pr-2 ml-2 cursor-pointer mx-auto">
              <FontAwesomeIcon icon={["fas", "arrow-right"]} />
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default Card;
