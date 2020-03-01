using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Microsoft.Extensions.Logging;
using Server.gRPC.Services;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Server.gRPC.Controllers
{
  public class ProjectController : Project.ProjectBase
  {
    private readonly ILogger<ProjectController> _logger;
    private readonly IProjectService _projectService;
    public ProjectController(IProjectService projectService, ILogger<ProjectController> logger)
    {
      _projectService = projectService;
      _logger = logger;
    }

    public override async Task<ProjectCreateReply> CreateProject(ProjectCreateRequest request, ServerCallContext context)
    {
      var exists = await _projectService.ProjectExistsAsync(request.ProjectId);
      if (exists)
      {
        throw new RpcException(new Status(StatusCode.AlreadyExists, "ProjectId already exists"));
      }

      var res = await _projectService.CreateProjectAsync(request);
      if (res.ProjectId == 0)
      {
        throw new RpcException(new Status(StatusCode.Internal, "Internal server error"));
      }
      return await Task.FromResult(res);
    }

        public override async Task<ProjectReply> GetProjects(ProjectGetRequest request, ServerCallContext context)
        {
            var projects =  await _projectService.GetAllProjectsAsync();

            return await Task.FromResult(new ProjectReply { Projects = { projects } });
        }
    }
}
