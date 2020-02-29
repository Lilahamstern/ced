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
        Console.WriteLine();
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
            List<ProjectModel> projects = new List<ProjectModel>();
            var dbProjects = await _projectService.GetAllProjectsAsync();

            foreach (var item in dbProjects)
            {
                var project = new ProjectModel();
                project.ProjectId = item.ProjectId;
                project.OrderId = item.OrderId;
                project.Name = item.Name;
                project.Description = item.Description;
                project.Manager = item.Manager;
                project.Client = item.Client;
                project.Sector = item.Sector;
                project.CreatedAt = Timestamp.FromDateTime(DateTime.SpecifyKind(item.CreatedAt.Value, DateTimeKind.Utc));
                project.UpdatedAt = Timestamp.FromDateTime(DateTime.SpecifyKind(item.UpdatedAt.Value, DateTimeKind.Utc));

                projects.Add(project);
            }

            var message = new ProjectReply { Projects = { projects } };

            return await Task.FromResult(message);
        }
    }
}
