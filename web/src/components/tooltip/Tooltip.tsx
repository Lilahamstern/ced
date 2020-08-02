import React, { Component } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

interface IProps {
  children: React.ReactNode;
  tooltipClassName: string;
  iconClassName: string;
}

interface IState {
  showTooltip: boolean;
}

export class Tooltip extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      showTooltip: false,
    };
  }

  toggleTooltip = (
    e: React.MouseEvent<SVGSVGElement, MouseEvent>,
    show: boolean
  ) => this.setState({ showTooltip: show });
  render() {
    const { showTooltip } = this.state;
    const { children, tooltipClassName, iconClassName } = this.props;
    return (
      <div>
        {showTooltip && (
          <div className={tooltipClassName}>
            <div className="bg-gray-900 opacity-75 text-white text-xs rounded py-1 pl-2 right-0 bottom-full truncate">
              {children}
            </div>
          </div>
        )}
        <div className={`absolute cursor-pointer ${iconClassName}`}>
          <p className="text-base mr-3">
            <FontAwesomeIcon
              onMouseEnter={(e) => this.toggleTooltip(e, true)}
              onMouseLeave={(e) => this.toggleTooltip(e, false)}
              icon={["fas", "info-circle"]}
            />
          </p>
        </div>
      </div>
    );
  }
}

export default Tooltip;
