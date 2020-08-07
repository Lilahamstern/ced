import React, { Component, FunctionComponent } from "react";

import { ProjectViewMode } from "../../../../enums";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IconName } from "@fortawesome/fontawesome-svg-core";

interface IProps {
  view: ProjectViewMode;
  clickHandler: (
    event: React.MouseEvent<HTMLButtonElement>,
    view: ProjectViewMode
  ) => void;
}

interface IState {}

export class Controller extends Component<IProps, IState> {
  render() {
    const { view, clickHandler } = this.props;
    return (
      <div className="bg-gray-800 text-sm text-gray-500 leading-none border-2 border-gray-800 rounded-full inline-flex">
        <Button
          onClick={clickHandler}
          view={ProjectViewMode.CARD}
          iconName="th-large"
          classes={`btn rounded-l-full focus:outline-none hover:text-gray-100 focus:text-gray-200 ${
            view === ProjectViewMode.CARD ? "active" : ""
          }`}
        />
        <Button
          onClick={clickHandler}
          view={ProjectViewMode.LIST}
          iconName="th-list"
          classes={`btn rounded-r-full focus:outline-none hover:text-white focus:text-gray-200 ${
            view === ProjectViewMode.LIST ? "active" : ""
          }`}
        />
      </div>
    );
  }
}

export default Controller;

interface ButtonProps {
  classes: string;
  view: ProjectViewMode;
  iconName: IconName;
  onClick: (
    event: React.MouseEvent<HTMLButtonElement>,
    view: ProjectViewMode
  ) => void;
}

const Button: FunctionComponent<ButtonProps> = (props: ButtonProps) => {
  const { view, classes, iconName, onClick } = props;
  return (
    <button className={classes} onClick={(e) => onClick(e, view)}>
      <FontAwesomeIcon
        icon={["fas", iconName]}
        className="w-4 h-4 mr-2 fill-current"
      />
      <span className="capitalize=">{view}</span>
    </button>
  );
};
