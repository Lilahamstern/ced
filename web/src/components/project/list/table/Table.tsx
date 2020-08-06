import React, { Component, FunctionComponent } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import Tooltip from "../../../tooltip/Tooltip";

import { IProject } from "../../../../redux/reducers/ProjectReducer";
import { connect } from "react-redux";
import { IAppState } from "../../../../redux/store/store";
import { getTimeSince } from "../../../../utils/utils";

interface IProps {
  projects: IProject[];
}

interface IState {}

export class Table extends Component<IProps, IState> {
  render() {
    const { projects } = this.props;

    return (
      <div className="bg-gray-800 w-full max-w-6xl rounded-lg border-2 border-gray-700">
        <table className="table-fixed border-collapse text-white w-full rounded-lg">
          <thead className="bg-gray-700 flex w-full shadow-b rounded-t-lg">
            <tr className="text-sm flex w-full justify-start text-left">
              <th className="table-head w-32 text-center">
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
          <TableBody projects={projects} />
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

// Table body
interface TableBodyProps {
  projects: IProject[];
}

const TableBody: FunctionComponent<TableBodyProps> = (
  props: TableBodyProps
) => {
  const { projects } = props;
  const tableRows: Array<any> = [];

  if (projects.length > 0) {
    projects.forEach((project, index) => {
      tableRows.push(
        <TableRow project={project} index={index} key={project.id} />
      );
    });

    return (
      <tbody className="flex flex-col items-start w-full relative">
        {tableRows}
      </tbody>
    );
  }

  return (
    <tbody className="flex flex-col  items-center w-full relative">
      No projects found
    </tbody>
  );
};

// Table row
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
      <td className="text-sm font-semibold p-2 pl-5 w-32">
        {getTimeSince(project.versions[0].updated_at)} ago
      </td>
      <td className="table-text w-1/12">{project.id}</td>
      <td className="table-text w-1/5">{project.versions[0].title}</td>
      <td className="table-text w-1/6">{project.versions[0].sector}</td>
      <td className="table-text w-1/6">{project.versions[0].client}</td>
      <td className="table-text w-1/6">{project.versions[0].manager}</td>
      <td className="table-text w-1/12">{project.versions[0].co}</td>
      <td>
        <Tooltip
          tooltipClassName="tooltip-table"
          iconClassName="tooltip-table-icon"
        >
          {project.versions[0].id}
        </Tooltip>
      </td>
    </tr>
  );
};
