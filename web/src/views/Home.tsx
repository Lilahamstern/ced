import React, { Component } from "react";

import { ProjectViewStatus } from "../enums";
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
          <Controller view={this.state.view} />
        </div>
        <div className="mt-6">
          <div className="flex justify-center w-full">
            <Cards projects={this.state.projects} />
          </div>
          <div className="flex justify-center rounded-lg">
            <Table projects={this.state.projects} />
          </div>
        </div>
      </div>
    );
  }
}

export default Home;
