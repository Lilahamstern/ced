using Api.Contracts.V1.Requests.Component;
using Api.Domain.Components;
using Api.Domain.Versions;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IComponentService
    {
        Task<bool> CreateComponentsAsync(List<Component> components);
        Task<List<Version>> GetProjectVersionsByProjectAsync(int projectId);
        Task<Version> GetProjectVersionByVersionAsync(int projectId, string version);
    }
}
