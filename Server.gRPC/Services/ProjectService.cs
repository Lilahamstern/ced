﻿using Google.Protobuf.WellKnownTypes;
using Server.gRPC.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using static DataLibrary.BusinessLogic.ProjectProcessor;

namespace Server.gRPC.Services
{
    public class ProjectService : IProjectService
    {
        public async Task<ProjectCreateReply> CreateProjectAsync(ProjectCreateRequest request)
        {
            CreateProject(request.ProjectId);

            int created = CreateProjectInformation(request.ProjectId, request.OrderId, request.Name,
                       request.Description, request.Manager, request.Client, request.Sector);

            if (created != 0)
            {
                request.ProjectId = 0;
            }
            return new ProjectCreateReply { ProjectId = request.ProjectId };
        }

        public async Task<Boolean> ProjectExistsAsync(int projectId)
        {
            int total = ProjectExists(projectId);
            return total != 0;
        }

        public async Task<List<ProjectModel>> GetAllProjectsAsync()
        {
            List<ProjectModel> projects = new List<ProjectModel>();
            var data = LoadProjects();

            foreach (var item in data)
            {
                var project = new ProjectModel();
                project.ProjectId = item.Id;
                project.OrderId = item.OrderId;
                project.Name = item.Name;
                project.Description = item.Description;
                project.Manager = item.Manager;
                project.Client = item.Client;
                project.Sector = item.Sector;
                project.CreatedAt = Timestamp.FromDateTime(DateTime.SpecifyKind(item.CreatedAt, DateTimeKind.Utc));
                project.UpdatedAt = Timestamp.FromDateTime(DateTime.SpecifyKind(item.UpdatedAt, DateTimeKind.Utc));

                projects.Add(project);
            }
            return projects;
        }
    }
}