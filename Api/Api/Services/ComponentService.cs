using Api.Contracts.V1.Requests;
using Api.Data;
using Api.Domain.Components;
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

        public async Task<int> CreateComponentVersionAsync(int projectId, ComponentInformation information)
        {

            var component = new Component
            {
                PId = projectId,
                Version = information.Version,
                Description = information.Description,
            };
            await _datacontext.Components.AddAsync(component);

            var created = await _datacontext.SaveChangesAsync();
            Console.WriteLine(component.Id);
            return created;
        }
    }
}
