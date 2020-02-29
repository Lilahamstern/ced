using BusinessLayer.Models.EntityFramework;
using DataAccessLayer;
using Microsoft.EntityFrameworkCore;
using Server.gRPC.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Server.gRPC.Services
{
  public class ProjectService : IProjectService
  {
    private readonly DataContext _dataContext;
    public ProjectService(DataContext dataContext)
    {
      _dataContext = dataContext;
    }
    public async Task<ProjectCreateReply> CreateProjectAsync(ProjectCreateRequest request)
    {
      var project = new BusinessLayer.Models.EntityFramework.Project { ProjectId = request.ProjectId };
      project.ProjectInformation = new List<ProjectInformation> {
                new ProjectInformation {
                    OrderId = request.OrderId,
                    Name = request.Name,
                    Description = request.Description,
                    Manager = request.Manager,
                    Client = request.Client,
                    Sector = request.Sector,
                }
            };

      await _dataContext.AddAsync(project);
      var created = await _dataContext.SaveChangesAsync();
      if (created != 0)
      {
        return new ProjectCreateReply
        {
          ProjectId = request.ProjectId
        };
      }
      return new ProjectCreateReply { ProjectId = 0 };
    }

    public async Task<Boolean> ProjectExistsAsync(int projectId)
    {
      var project = await _dataContext.Projects.Where(x => x.ProjectId == projectId).SingleOrDefaultAsync();
      return project != null;
    }

    public async Task<List<ProjectInformation>> GetAllProjectsAsync()
    {
      return await _dataContext.ProjectInformation.ToListAsync();
    }
  }
}
