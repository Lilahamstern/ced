import React, { Component, FunctionComponent } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import Tooltip from "../../../tooltip/Tooltip";

import { IProject } from "../../../../redux/reducers/ProjectReducer";
import { connect } from "react-redux";
import { IAppState } from "../../../../redux/store/store";

interface IProps {
  projects: IProject[];
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
                <FontAwesomeIcon icon={["far", "clock"]} />
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

const mapStateToProps = (store: IAppState) => {
  return {
    projects: store.projectState.projects,
  };
};

export default connect(mapStateToProps)(Table);

interface TableRowProps {
  index: number;
  project: IProject;
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
        {project.versions[0].UpdatedAt}d ago
      </td>
      <td className="table-text w-1/12">{project.versions[0].OrderID}</td>
      <td className="table-text w-1/5">{project.versions[0].Title}</td>
      <td className="table-text w-1/6">{project.versions[0].Sector}</td>
      <td className="table-text w-1/6">{project.versions[0].Client}</td>
      <td className="table-text w-1/6">{project.versions[0].Manager}</td>
      <td className="table-text w-1/12">{project.versions[0].Co}</td>
      <td>
        <Tooltip
          tooltipClassName="tooltip-table"
          iconClassName="tooltip-table-icon"
        >
          {project.ID}
        </Tooltip>
      </td>
    </tr>
  );
};
