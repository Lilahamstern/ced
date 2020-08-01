import React, { Component, FunctionComponent } from "react";

import Tooltip from "../../../tooltip/Tooltip";

import { Projects, Project } from "../../../../stores/types";

interface IProps {
  projects: Projects;
}

interface IState {}

export class Table extends Component<IProps, IState> {
  render() {
    const { projects } = this.props;
    const tableRows: Array<any> = [];

    projects.forEach((project, index) => {
      tableRows.push(
        <TableRow project={project} index={index} key={project.ID} />
      );
    });

    return (
      <div className="bg-gray-800 w-full max-w-6xl rounded-lg border-2 border-gray-700">
        <table className="table-fixed border-collapse text-white w-full rounded-lg">
          <thead className="bg-gray-700 flex w-full shadow-b rounded-t-lg">
            <tr className="text-sm flex w-full justify-start text-left">
              <th className="table-head w-1/12 text-center">
                <i className="far fa-clock"></i>
              </th>
              <th className="table-head w-1/12">Project</th>
              <th className="table-head w-1/5">Title</th>
              <th className="table-head w-1/6">Sector</th>
              <th className="table-head w-1/6">Client</th>
              <th className="table-head w-1/6">Manager</th>
              <th className="table-head w-1/12 truncate">Co2-ekv/m2</th>
            </tr>
          </thead>
          <tbody className="flex flex-col items-start w-full relative">
            {tableRows}
          </tbody>
        </table>
      </div>
    );
  }
}

export default Table;

interface TableRowProps {
  index: number;
  project: Project;
}

const TableRow: FunctionComponent<TableRowProps> = (props: TableRowProps) => {
  const { project, index } = props;
  return (
    <tr
      className={`flex w-full items-center cursor-pointer hover:bg-gray-900 ${
        index % 2 !== 0 ? "bg-gray-700" : ""
      }`}
    >
      <td className="text-sm font-semibold p-2 pl-5 w-1/12">
        {project.UpdatedAt}d ago
      </td>
      <td className="table-text w-1/12">{project.OrderID}</td>
      <td className="table-text w-1/5">{project.Title}</td>
      <td className="table-text w-1/6">{project.Sector}</td>
      <td className="table-text w-1/6">{project.Client}</td>
      <td className="table-text w-1/6">{project.Manager}</td>
      <td className="table-text w-1/12">{project.Co}</td>
      <td className="tooltip-table">
        <Tooltip projectID={project.ID} />
      </td>
    </tr>
  );
};
