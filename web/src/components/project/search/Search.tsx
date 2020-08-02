import React, { Component } from "react";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

interface IProps {}

interface IState {
  search: string;
}
export class Search extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      search: "",
    };
  }

  handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    this.setState({ search: e.target.value });
  }
  render() {
    return (
      <div className="relative text-white max-w-xxs md:max-w-xs w-full">
        <label>
          <input
            className="search-input"
            type="search"
            placeholder="Search"
            onChange={(e) => this.handleChange(e)}
            value={this.state.search}
          />
        </label>
        <button type="submit" className="absolute right-0 top-0 mt-2 mr-6">
          <FontAwesomeIcon
            icon={["fas", "search"]}
            className="h-4 w-4 fill-current"
          />
        </button>
      </div>
    );
  }
}

export default Search;
