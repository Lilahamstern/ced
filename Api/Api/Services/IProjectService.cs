using Api.Contracts.V1.Responses;
using Api.Domain;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IProjectService
    {


        Task<bool> CreateProjectAsync(Project project);
        Task<List<Project>> GetProjectsAsync(int limit);

        Task<List<Project>> GetProjectsAsync(int limit, string search);

        Task<Project> GetProjectByIdAsync(string projectId);

        Task<bool> UpdateProjectAsync(Project projectToUpdate);

        Task<bool> DeleteProjectAsync(string projectId);

    }
}
