using Api.Contracts.V1.Requests;
using Api.Data;
using Api.Domain.Components;
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

        public async Task<bool> CreateComponentsAsync(List<ComponentData> components)
        {
            _datacontext.ComponentDatas.AddRange(components);
            var created = await _datacontext.SaveChangesAsync();
            return created > 0;
        }
        public async Task<int> CreateComponentVersionAsync(int projectId, ComponentInformationRequest information)
        {

            var component = new ComponentInformation
            {
                PId = projectId,
                Version = information.Version,
                Description = information.Description,
            };
            await _datacontext.Components.AddAsync(component);

            await _datacontext.SaveChangesAsync();

            return component.Id;
        }

        public async Task<ComponentInformation> GetComponentInformationByVersionAsync(int projectId, string version)
        {
            return await _datacontext.Components.SingleOrDefaultAsync(x => x.PId == projectId && x.Version == version);
        }

        public async Task<List<ComponentInformation>> GetComponentInformationsAsync(int projectId)
        {
            var list = await _datacontext.Components.ToListAsync();
            return list.Where(x => x.PId == projectId).ToList();
        }
    }
}
