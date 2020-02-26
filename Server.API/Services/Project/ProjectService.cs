using BusinessLayer.Models.EntityFramework;
using DataAccessLayer;
using Microsoft.EntityFrameworkCore;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Services
{
    public class ProjectService : IProjectService
    {

        private readonly DataContext _dataContext;

        public ProjectService(DataContext dataContext)
        {
            _dataContext = dataContext;
        }

        public async Task<bool> CreateProjectAsync(ProjectInformation project)
        {
            await _dataContext.Projects.AddAsync(new Project { ProjectId = project.ProjectId });

            await _dataContext.ProjectInformation.AddAsync(project);
            var created = await _dataContext.SaveChangesAsync();
            return created > 0;
        }

        public async Task<ProjectInformation> GetProjectByIdAsync(int projectId)
        {
            var result = await _dataContext.ProjectInformation.SingleOrDefaultAsync(x => x.ProjectId == projectId);

            return result;

        }

        public async Task<List<ProjectInformation>> GetProjectsAsync(int limit)
        {
            return await _dataContext.ProjectInformation.Take(limit).ToListAsync();
        }

        public async Task<List<ProjectInformation>> GetProjectsAsync(int limit, string search)
        {
            var projects = await _dataContext.ProjectInformation.Take(limit).Where(p =>
            p.ProjectId.ToString().Contains(search) ||
            p.OrderId.ToString().Contains(search) ||
            p.Client.Contains(search) ||
            p.Name.Contains(search) ||
            p.Manager.Contains(search)).ToListAsync();
            return projects;
        }

        public async Task<bool> UpdateProjectAsync(int projectId, ProjectInformation projectToUpdate)
        {
            var result = await _dataContext.ProjectInformation.SingleOrDefaultAsync(p => p.ProjectId == projectId);
            if (result != null)
            {
                result.OrderId = projectToUpdate.OrderId;
                result.Name = projectToUpdate.Name;
                result.Description = projectToUpdate.Description;
                result.Manager = projectToUpdate.Manager;
                result.Client = projectToUpdate.Client;
                result.Sector = projectToUpdate.Sector;
                var updated = await _dataContext.SaveChangesAsync();
                return updated > 0;
            }
            return false;
        }

        public async Task<bool> DeleteProjectAsync(int projectId)
        {
            var project = await GetProjectByIdAsync(projectId);
            _dataContext.ProjectInformation.Remove(project);
            var deleted = await _dataContext.SaveChangesAsync();
            return deleted > 0;
        }
    }
}
