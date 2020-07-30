import VersionModel from "./version"

export default class ProjectModel {
  public ID: Number;
  public Versions: Array<VersionModel>;
  public CreatedAt: string;
  public UpdatedAt: string;
}