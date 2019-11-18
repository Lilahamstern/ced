export interface Project {
    id: number,
    name: string,
    client: string,
    sector: string,
    co2: number,
}

export interface ProjectState {
    project?: Project;
    error: boolean;
    projects?: [Project];
}