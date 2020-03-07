using DataAccessLayer.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IProjectService
    {
        Task<bool> CreateProjectAsync(ProjectInformation project);
        Task<List<ProjectInformation>> GetProjectsAsync(int limit);

        Task<List<ProjectInformation>> GetProjectsAsync(int limit, string search);

        Task<ProjectInformation> GetProjectByIdAsync(int projectId);

        Task<bool> UpdateProjectAsync(int projectId, ProjectInformation projectToUpdate);

        Task<bool> DeleteProjectAsync(int projectId);

    }
}
