import React, { Component, FunctionComponent, MouseEvent } from "react";

import { ProjectViewMode } from "../enums";
import { viewCards, widthLessThen } from "../utils/utils";
import ProjectSearch from "../components/project/search/Search";
import Controller from "../components/project/list/controller/Controller";
import Cards from "../components/project/list/cards/Cards";
import Table from "../components/project/list/table/Table";
import { IProject, IProjectState } from "../redux/reducers/ProjectReducer";
import { connect } from "react-redux";
import { IAppState } from "../redux/store/store";
import { ThunkDispatch } from "redux-thunk";
import {
  IProjectSelectAction,
  SelectProject,
} from "../redux/actions/ProjectActions";
import { useHistory } from "react-router-dom";

interface IProps {}

interface IState {
  viewMode: ProjectViewMode;
  viewport: {
    width: number;
  };
}

export class HomeView extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      viewMode: ProjectViewMode.CARD,
      viewport: {
        width: window.innerWidth,
      },
    };
  }

  handleControllerClick = (
    e: React.MouseEvent<HTMLButtonElement>,
    mode: ProjectViewMode
  ) => {
    e.preventDefault();
    this.setState({ viewMode: mode });
  };

  updateDimension = () => {
    this.setState({ viewport: { width: window.innerWidth } });
  };

  componentDidMount() {
    window.addEventListener("resize", this.updateDimension);
  }
  componentWillUnmount() {
    window.removeEventListener("resize", this.updateDimension);
  }

  render() {
    return (
      <div>
        <div className="flex justify-center mt-12">
          <ProjectSearch />
        </div>
        <div className="mt-6 hidden lg:flex lg:justify-center">
          <Controller
            clickHandler={this.handleControllerClick}
            view={this.state.viewMode}
          />
        </div>
        <div className="mt-6">
          <ProjectsViewContainer
            width={this.state.viewport.width}
            view={this.state.viewMode}
          />
        </div>
      </div>
    );
  }
}

export default HomeView;

interface IProjectsViewProps {
  width: number;
  view: ProjectViewMode;
  projects: IProject[];
  selectProject: (id: string) => Promise<void>;
}
const ProjectsView: FunctionComponent<IProjectsViewProps> = (
  props: IProjectsViewProps
) => {
  const { width, view, projects, selectProject } = props;
  const history = useHistory();

  const handleSelectProjectClick = (e: MouseEvent, id: string) => {
    e.preventDefault();
    selectProject(id);
    history.push(`/project/${id}`);
  };

  if (!viewCards(view) && !widthLessThen(width, 1024)) {
    return (
      <div className="justify-center rounded-lg hidden lg:flex">
        <Table projects={projects} onRowClick={handleSelectProjectClick} />
      </div>
    );
  }
  return <Cards projects={projects} onCardClick={handleSelectProjectClick} />;
};

const mapStateToProps = (store: IAppState) => {
  return {
    projects: store.projectState.projects,
  };
};

const mapDispatchToProps = (
  dispatch: ThunkDispatch<IProjectState, null, IProjectSelectAction>
) => {
  return {
    selectProject: (id: string) => dispatch(SelectProject(id)),
  };
};

const ProjectsViewContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ProjectsView);
