﻿using Api.Contracts.V1;
using Api.Contracts.V1.Responses;
using Api.Controllers.V1.Requests.Project;
using Api.Controllers.V1.Responses;
using Api.Domain;
using Api.Extentions;
using Api.Services;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Swashbuckle.AspNetCore.Annotations;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Controllers.V1
{
    [SwaggerTag("Project requests")]
    [Produces("application/json")]
    public class ProjectController : Controller
    {

        private readonly IProjectService _projectService;

        public ProjectController(IProjectService projectService)
        {
            _projectService = projectService;
        }

        /// <summary>
        /// Create project endpoint
        /// </summary>
        /// <param name="request"> Example request body to create project</param>
        [HttpPost(ApiRoutes.Project.Create)]
        [ProducesResponseType(typeof(ProjectResponse), 201)]
        public async Task<IActionResult> Create([FromBody] CreateProjectRequest request)
        {
            var project = new Project { 
                PId = request.ProjectId,
                OId = request.OrderId,
                Name = request.Name,
                Description = request.Description,
                Client = request.Client,
                Manager = request.Manager,
                Sector = request.Sector,  
            };

            var projectCheck = await _projectService.GetProjectByIdAsync(project.PId);
            if (projectCheck != null)
            {
                return BadRequest(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            FieldName = "ProjectId",
                            Message= "Project ID already exists",
                        }
                    }
                });
            }

            await _projectService.CreateProjectAsync(project);


            var baseUrl = $"{HttpContext.Request.Scheme}://{HttpContext.Request.Host.ToUriComponent()}";
            var location = baseUrl + "/" + ApiRoutes.Project.Get.Replace("{projectId}", project.PId.ToString());

            var response = new ProjectResponse { ProjectId = project.PId };

            return Created(location, response);
        }
        /// <summary>
        /// Get specifed project by ID.
        /// </summary>
        /// <param name="projectId">ProjectId to specifed project</param>
        [HttpGet(ApiRoutes.Project.Get)]
        public async Task<IActionResult> Get([FromRoute] int projectId)
        {
            if (projectId == 0)
            {
                return BadRequest(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            FieldName = "projectId",
                            Message = "Project id is requierd",
                        }
                    }
                });
            }

            var project = await _projectService.GetProjectByIdAsync(projectId);

            if (project == null)
            {
                return NotFound(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            Message= "Project could not be found",
                        }
                    }
                });
            }

            return Ok(project);
        }

        /// <summary>
        /// Returns all projects if no project is found 404 will be retunred s
        /// </summary>
        /// <param name="search">Provide search query of ProjectID, Name, OrderID, Client and Manager</param>
        /// <param name="limit">Set a limit of results you want</param> 
        [HttpGet(ApiRoutes.Project.GetAll)]
        public async Task<IActionResult> GetAll([FromQuery] string search, [FromQuery] int limit)
        {
            List<Project> projects;
            if (limit <= 0)
                    limit = 10000;

            if (!string.IsNullOrEmpty(search))
            {
                projects = await _projectService.GetProjectsAsync(limit, search);
            } else
            {
                projects = await _projectService.GetProjectsAsync(limit);
            }

            if (projects.Count ==  0)
            {
                return NotFound(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            Message= "No projects found",
                        }
                    }
                });
            }
            return Ok(projects);
        }
        /// <summary>
        /// Update specified project
        /// </summary>
        /// <param name="projectId">ProjectId to project</param>
        /// <param name="request">Example body to update project</param>
        [HttpPut(ApiRoutes.Project.Update)]
        public async Task<IActionResult> Update([FromRoute] int projectId, [FromBody] UpdateProjectRequest request)
        {
            var project = new Project
            { 
                OId = request.OrderId,
                Name = request.Name,
                Description = request.Description,
                Manager = request.Manager,
                Client = request.Client,
                Sector = request.Sector,
            };

            var updated = await _projectService.UpdateProjectAsync(projectId, project);

            if (!updated)
                return NotFound(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            FieldName = "ProjectId",
                            Message= "Project could not be found",
                        }
                    }
                });

            return Ok(new ProjectResponse { ProjectId = projectId});
        }
        /// <summary>
        /// Deleted specifed project.
        /// </summary>
        /// <param name="projectId">ProjectId to project</param>
        [HttpDelete(ApiRoutes.Project.Delete)]
        [Authorize(/*AuthenticationSchemes = JwtBearerDefaults.AuthenticationScheme,*/ Roles = "Admin")]
        public async Task<IActionResult> Delete([FromRoute] int projectId)
        {
            if (!string.IsNullOrEmpty(projectId.ToString()))
            {
                return BadRequest(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            FieldName = "projectId",
                            Message = "Project id is requierd",
                        }
                    }
                });
            }

            var deleted = await _projectService.DeleteProjectAsync(projectId);

            if (deleted)
                return NoContent();

            return NotFound(
                new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            Message = "Project not found",
                        }
                    }
                });
        }

    }
}
