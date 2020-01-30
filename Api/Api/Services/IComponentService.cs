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
        Task<bool> CreateComponentsAsync(List<ComponentData> components);

        Task<List<ComponentInformation>> GetComponentInformationsAsync(int projectId);
        Task<ComponentInformation> GetComponentInformationByVersionAsync(int projectId, string version);
        Task<int> CreateComponentVersionAsync(int projectId, ComponentInformationRequest information);
    }
}
