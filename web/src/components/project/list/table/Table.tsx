import React, { FunctionComponent, MouseEvent } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import Tooltip from "../../../tooltip/Tooltip";

import { IProject } from "../../../../redux/reducers/ProjectReducer";
import { getTimeSince } from "../../../../utils/utils";

interface IProps {
  projects: IProject[];
  onRowClick: (e: MouseEvent, id: string) => void;
}

export const Table: FunctionComponent<IProps> = (props: IProps) => {
  const { projects, onRowClick } = props;
  let tableBody: any;

  if (projects.length > 0) {
    tableBody = (
      <tbody className="flex flex-col items-start w-full relative">
        {projects.map((project, index) => {
          return (
            <TableRow
              onClick={(e) => onRowClick(e, project.id.toString())}
              project={project}
              index={index}
              key={project.id}
            />
          );
        })}
      </tbody>
    );
  } else {
    tableBody = (
      <tbody className="flex flex-col  items-center w-full relative">
        No projects found
      </tbody>
    );
  }

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
        {tableBody}
      </table>
    </div>
  );
};

export default Table;

// Table row
interface TableRowProps {
  index: number;
  project: IProject;
  onClick: (e: MouseEvent) => void;
}

const TableRow: FunctionComponent<TableRowProps> = (props: TableRowProps) => {
  const { project, index, onClick } = props;

  return (
    <tr
      className={`flex w-full items-center cursor-pointer hover:bg-gray-900 ${
        index % 2 !== 0 ? "bg-gray-700" : ""
      }`}
      onClick={(e) => onClick(e)}
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

// const mapDispatchToProps = (
//   dispatch: ThunkDispatch<IProjectState, null, IProjectSelectAction>
// ) => {
//   return {
//     selectProject: (id: string) => dispatch(SelectProject(id)),
//   };
// };
