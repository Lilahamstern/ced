using Api.Contracts.V1.Requests;
using Api.Domain.Components;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IComponentService
    {
        Task<bool> CreateComponentsAsync(List<Component> components);

        Task<int> CreateComponentVersionAsync(int projectId, ComponentInformation information);
    }
}
