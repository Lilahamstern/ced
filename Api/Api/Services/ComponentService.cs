using Api.Contracts.V1.Requests;
using Api.Data;
using Api.Domain.Components;
using Api.Domain.Projects;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Services
{
    public class ComponentService : IComponentService
    {

        private readonly DataContext _datacontext;
        public ComponentService(DataContext dataContext)
        {
            _datacontext = dataContext;
        }

        public async Task<bool> CreateComponentsAsync(List<Component> components)
        {
            _datacontext.Components.AddRange(components);
            var created = await _datacontext.SaveChangesAsync();
            return created > 0;
        }
        public async Task<int> CreateProjectVersionAsync(int projectId, ComponentInformationRequest information)
        {

            var projectVersion = new ProjectVersion
            {
                PId = projectId,
                Version = information.Version,
                Description = information.Description,
            };
            await _datacontext.ProjectVersions.AddAsync(projectVersion);

            await _datacontext.SaveChangesAsync();

            return projectVersion.Id;
        }

        public async Task<ProjectVersion> GetProjectVersionByVersionAsync(int projectId, string version)
        {
            return await _datacontext.ProjectVersions.SingleOrDefaultAsync(x => x.PId == projectId && x.Version == version);
        }

        public async Task<List<ProjectVersion>> GetProjectVersionsByProjectAsync(int projectId)
        {
            var list = await _datacontext.ProjectVersions.ToListAsync();
            return list.Where(x => x.PId == projectId).ToList();
        }
    }
}
