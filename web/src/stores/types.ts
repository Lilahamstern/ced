export interface Project {
  ID: string;
  OrderID: number;
  Title: string;
  Description: string;
  Manager: string;
  Client: string;
  Sector: string;
  Co: number;
  CreatedAt: string;
  UpdatedAt: string;
}

export interface Projects extends Array<Project> {}
