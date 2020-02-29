using BusinessLayer.Models.EntityFramework;
using Server.gRPC.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Server.gRPC.Services
{
    public interface IProjectService
    {
        Task<ProjectCreateReply> CreateProjectAsync(ProjectCreateRequest request);

        Task<Boolean> ProjectExistsAsync(int projectId);

        Task<List<ProjectInformation>> GetAllProjectsAsync();
    }
}
