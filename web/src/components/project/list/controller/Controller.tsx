import React, { Component, FunctionComponent } from "react";

import { ProjectViewStatus } from "../../../../enums";

interface IProps {
  view: ProjectViewStatus;
}

interface IState {}

export class Controller extends Component<IProps, IState> {
  render() {
    const { view } = this.props;

    console.log(view === ProjectViewStatus.CARD);
    return (
      <div className="bg-gray-800 text-sm text-gray-500 leading-none border-2 border-gray-800 rounded-full inline-flex">
        <Button
          title="Card"
          classes={`btn rounded-l-full focus:outline-none hover:text-gray-100 focus:text-gray-200 ${
            view === ProjectViewStatus.CARD ? "active" : "active"
          }`}
        />
        <Button
          title="List"
          classes={`btn rounded-r-full focus:outline-none hover:text-white focus:text-gray-200 ${
            view === ProjectViewStatus.LIST ? "active" : ""
          }`}
        />
      </div>
    );
  }
}

export default Controller;

interface ButtonProps {
  classes: string;
  title: string;
}

const Button: FunctionComponent<ButtonProps> = (props: ButtonProps) => {
  return (
    <button className="btn rounded-r-full focus:outline-none hover:text-white focus:text-gray-200">
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        className="fill-current w-4 h-4 mr-2"
      >
        <line x1="8" y1="6" x2="21" y2="6"></line>
        <line x1="8" y1="12" x2="21" y2="12"></line>
        <line x1="8" y1="18" x2="21" y2="18"></line>
        <line x1="3" y1="6" x2="3.01" y2="6"></line>
        <line x1="3" y1="12" x2="3.01" y2="12"></line>
        <line x1="3" y1="18" x2="3.01" y2="18"></line>
      </svg>
      <span>{props.title}</span>
    </button>
  );
};
