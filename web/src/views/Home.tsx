import React, { Component, FunctionComponent } from "react";

import { ProjectViewStatus } from "../enums";
import { viewCards, widthLessThen } from "../utils/utils";
import ProjectSearch from "../components/project/search/Search";
import Controller from "../components/project/list/controller/Controller";
import Cards from "../components/project/list/cards/Cards";
import Table from "../components/project/list/table/Table";
import { Projects } from "../stores/types";

interface IProps {}

interface IState {
  view: ProjectViewStatus;
  viewport: {
    width: number;
  };
  projects: Projects;
}

export class Home extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      view: ProjectViewStatus.CARD,
      viewport: {
        width: 0,
      },
      projects: [
        {
          ID: "324df635-ab3b-4601-97fd-d404aefdfa43",
          OrderID: 29192378,
          Description: "Nothine",
          Title: "Twitch",
          Manager: "So,eone",
          Client: "Dude",
          Sector: "Something",
          Co: 20192,
          CreatedAt: "3",
          UpdatedAt: "3",
        },
        {
          ID: "324df635-ab3b-4601-97fd-d404aefdfa44",
          OrderID: 29192378,
          Description: "Nothine",
          Title: "Twitch",
          Manager: "So,eone",
          Client: "Dude",
          Sector: "Something",
          Co: 20192,
          CreatedAt: "3",
          UpdatedAt: "3",
        },
        {
          ID: "324df635-ab3b-4601-97fd-d404aefdfa45",
          OrderID: 29192378,
          Description: "Nothine",
          Title: "Twitch",
          Manager: "So,eone",
          Client: "Dude",
          Sector: "Something",
          Co: 20192,
          CreatedAt: "3",
          UpdatedAt: "3",
        },
      ],
    };
  }

  clickHandler = (
    e: React.MouseEvent<HTMLButtonElement>,
    view: ProjectViewStatus
  ) => {
    e.preventDefault();
    this.setState({ view });
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
          <Controller clickHandler={this.clickHandler} view={this.state.view} />
        </div>
        <div className="mt-6">
          <ProjectsView
            width={this.state.viewport.width}
            view={this.state.view}
            projects={this.state.projects}
          />
        </div>
      </div>
    );
  }
}

export default Home;

interface IProjectsViewProps {
  width: number;
  view: ProjectViewStatus;
  projects: Projects;
}
const ProjectsView: FunctionComponent<IProjectsViewProps> = (
  props: IProjectsViewProps
) => {
  const { width, view, projects } = props;
  if (!viewCards(view) && !widthLessThen(width, 1024)) {
    return (
      <div className="justify-center rounded-lg hidden lg:flex">
        <Table projects={projects} />
      </div>
    );
  }
  return <Cards projects={projects} />;
};
