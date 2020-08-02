import { ProjectViewStatus } from "../enums/index";

export function widthLessThen(width: number, minWidth: number): boolean {
  return width < minWidth;
}

export function viewCards(status: ProjectViewStatus): boolean {
  return status === ProjectViewStatus.CARD;
}
