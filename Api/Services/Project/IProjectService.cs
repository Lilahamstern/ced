using Api.Contracts.V1.Responses;
using Api.Contracts.V1.Responses.Project;
using Api.Domain;
using Api.Domain.Projects;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IProjectService
    {
        Task<bool> CreateProjectAsync(Project project);
        Task<List<Project>> GetProjectsAsync(int limit);

        Task<List<Project>> GetProjectsAsync(int limit, string search);

        Task<Project> GetProjectByIdAsync(int projectId);

        Task<bool> UpdateProjectAsync(int projectId, Project projectToUpdate);

        Task<bool> DeleteProjectAsync(int projectId);

    }
}
