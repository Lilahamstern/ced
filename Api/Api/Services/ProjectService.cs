﻿using Api.Data;
using Api.Domain;
using Microsoft.EntityFrameworkCore;
using System;
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

        public async Task<Project> GetProjectByIdAsync(Guid projectId)
        {
            return await _dataContext.Projects.SingleOrDefaultAsync(x => x.Id == projectId);
        }

        public async Task<List<Project>> GetProjectsAsync()
        {
            return await _dataContext.Projects.ToListAsync();
        }

        public async Task<bool> UpdateProjectAsync(Project projectToUpdate)
        {
            _dataContext.Projects.Update(projectToUpdate);
            var updated = await _dataContext.SaveChangesAsync();
            return updated > 0;
        }

        public async Task<bool> DeleteProjectAsync(Guid projectId)
        {
            var project = await GetProjectByIdAsync(projectId);
            _dataContext.Projects.Remove(project);
            var deleted = await _dataContext.SaveChangesAsync();
            return deleted > 0;
        }

        public async Task<bool> UserOwnProjectAsync(Guid projectId, string userId)
        {
            Project project = await _dataContext.Projects.AsNoTracking().SingleOrDefaultAsync(x => x.Id == projectId);

            if (project == null)
            {
                return false;
            }

            if (project.UserId != userId)
            {
                return false;
            };

            return true;
        }
    }
}
