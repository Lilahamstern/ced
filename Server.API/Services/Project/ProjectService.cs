using BusinessLayer.Models;
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

        public async Task<bool> CreateProjectAsync(Project project)
        {
            await _dataContext.Projects.AddAsync(project);
            var created = await _dataContext.SaveChangesAsync();
            return created > 0;
        }

        public async Task<Project> GetProjectByIdAsync(int projectId)
        {
            var result = await _dataContext.Projects.SingleOrDefaultAsync(x => x.PId == projectId);

            return result;

        }

        public async Task<List<Project>> GetProjectsAsync(int limit)
        {
            return await _dataContext.Projects.Take(limit).ToListAsync();
        }

        public async Task<List<Project>> GetProjectsAsync(int limit, string search)
        {
            var projects = await _dataContext.Projects.Take(limit).Where(p =>
            p.OId.ToString().Contains(search) ||
            p.PId.ToString().Contains(search) ||
            p.Client.Contains(search) ||
            p.Name.Contains(search) ||
            p.Manager.Contains(search)).ToListAsync();
            return projects;
        }

        public async Task<bool> UpdateProjectAsync(int projectId, Project projectToUpdate)
        {
            var result = await _dataContext.Projects.SingleOrDefaultAsync(p => p.PId == projectId);
            if (result != null)
            {
                result.OId = projectToUpdate.OId;
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
            _dataContext.Projects.Remove(project);
            var deleted = await _dataContext.SaveChangesAsync();
            return deleted > 0;
        }
    }
}
