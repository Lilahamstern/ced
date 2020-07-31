import React, { Component } from 'react'

import { ProjectViewStatus } from '../enums'
import ProjectSearch from '../components/project/search/Search'
import Controller from '../components/project/list/controller/Controller'

interface IProps {

}

interface IState {
  view: ProjectViewStatus
}

export class Projects extends Component<IProps, IState> {

  constructor(props: IProps) {
    super(props);
    this.state = {
      view: ProjectViewStatus.CARD
    }
  }

  render() {
    return (
      <div>
        <div className="flex justify-center mt-12">
          <ProjectSearch />
          <Controller view={this.state.view} />
        </div>
      </div>
    )
  }
}

export default Projects
