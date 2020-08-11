import { ProjectViewMode } from "../enums/index";
import moment from "moment";

export function widthLessThen(width: number, minWidth: number): boolean {
  return width < minWidth;
}

export function viewCards(status: ProjectViewMode): boolean {
  return status === ProjectViewMode.CARD;
}

export function getTimeSince(date: string, withoutSuffix: boolean): string {
  let unix = Date.parse(date);
  return moment.utc(unix).fromNow(withoutSuffix);
}
