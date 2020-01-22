using Api.Data;
using Api.Domain;
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
            return await _dataContext.Projects.SingleOrDefaultAsync(x => x.PId == projectId);
        }

        public async Task<List<Project>> GetProjectsAsync(int limit)
        {
            return await _dataContext.Projects.Take(limit).ToListAsync();
        }

        public async Task<List<Project>> GetProjectsAsync(int limit, string search)
        {
            var projects = await _dataContext.Projects.ToListAsync();
            var filterd = projects.Take(limit).Where(p => 
            //p.OId.Contains(search) || 
            //p.PId.Contains(search) || 
            p.Client.Contains(search) ||
            p.Name.Contains(search) ||
            p.Manager.Contains(search)).ToList();
            return filterd;
        }

        public async Task<bool> UpdateProjectAsync(Project projectToUpdate)
        {
            _dataContext.Projects.Update(projectToUpdate);
            var updated = await _dataContext.SaveChangesAsync();
            return updated > 0;
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
