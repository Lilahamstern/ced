using Api.Contracts.V1.Requests;
using Api.Domain.Components;
using Api.Domain.Projects;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IComponentService
    {
        Task<bool> CreateComponentsAsync(List<Component> components);
        Task<List<ProjectVersion>> GetProjectVersionsByProjectAsync(int projectId);
        Task<ProjectVersion> GetProjectVersionByVersionAsync(int projectId, string version);
        Task<int> CreateProjectVersionAsync(int projectId, ComponentInformationRequest information);
    }
}
