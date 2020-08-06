import { ProjectViewStatus } from "../enums/index";
import moment from "moment";

export function widthLessThen(width: number, minWidth: number): boolean {
  return width < minWidth;
}

export function viewCards(status: ProjectViewStatus): boolean {
  return status === ProjectViewStatus.CARD;
}

export function getTimeSince(date: string): string {
  let unix = Date.parse(date);
  return moment.utc(unix).fromNow(true);
}
