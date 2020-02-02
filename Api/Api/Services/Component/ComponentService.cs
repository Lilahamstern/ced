using Api.Contracts.V1.Requests;
using Api.Data;
using Api.Domain.Components;
using Api.Domain.Projects;
using Microsoft.EntityFrameworkCore;
using System;
using Version = Api.Domain.Versions.Version;
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

        public async Task<Version> GetProjectVersionByVersionAsync(int projectId, string version)
        {
            return await _datacontext.Versions.SingleOrDefaultAsync(x => x.PId == projectId && x.Title == version);
        }

        public async Task<List<Version>> GetProjectVersionsByProjectAsync(int projectId)
        {
            var list = await _datacontext.Versions.ToListAsync();
            return list.Where(x => x.PId == projectId).ToList();
        }
    }
}
