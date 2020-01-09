using Api.Domain;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IProjectService
    {


        Task<bool> CreateProjectAsync(Project project);
        Task<List<Project>> GetProjectsAsync();

        Task<Project> GetProjectByIdAsync(Guid projectId);

        Task<bool> UpdateProjectAsync(Project projectToUpdate);

        Task<bool> DeleteProjectAsync(Guid projectId);

        Task<bool> UserOwnProjectAsync(Guid projectId, string userId);
    }
}
