using System.ComponentModel;
using ProjectModel = Api.Domain.Projects.Project;

namespace Api.Contracts.V1.Responses.Project
{
    public class ProjectResponse
    {
        public ProjectResponse(ProjectModel project)
        {
            ProjectId = project.PId;
            OrderId = project.OId;
            Name = project.Name;
            Description = project.Description;
            Client = project.Client;
            Sector = project.Sector;
        }

        public ProjectResponse(int projectId)
        {
            ProjectId = projectId;
        }

        public int ProjectId { get; set; }
        public int? OrderId { get; set; } = null;
        public string Name { get; set; }
        public string Description { get; set; }
        public string Manager { get; set; }
        public string Client { get; set; }
        public string Sector { get; set; }
        public float? Co { get; set; } = null;

    }
}
