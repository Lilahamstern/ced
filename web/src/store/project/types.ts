export interface Project {
  id: number,
  name: string,
  client: string,
  sector: string,
  co2: number,
  state: string,
  desc: string,
}

export interface ProjectState {
  project?: Project;
  error: boolean;
  projects?: Project[];
}