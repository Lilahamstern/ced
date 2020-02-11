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

        public async Task<bool> AddComponentsAsync(List<Component> components)
        {
            _datacontext.Components.AddRange(components);
            var created = await _datacontext.SaveChangesAsync();
            return created > 0;
        }
    }
}
