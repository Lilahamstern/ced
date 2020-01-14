using Api.Contracts.V1;
using Api.Contracts.V1.Responses;
using Api.Controllers.V1.Requests;
using Api.Controllers.V1.Responses;
using Api.Domain;
using Api.Extentions;
using Api.Services;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Controllers.V1
{
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
        /// <param name="request"> Wuts this</param>
        [HttpPost(ApiRoutes.Projects.Create)]
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
            var location = baseUrl + "/" + ApiRoutes.Projects.Get.Replace("{projectId}", project.PId.ToString());

            var response = new ProjectResponse { PId = project.PId };

            return Created(location, response);
        }
        /// <summary>
        /// Get specifed project by ID.
        /// </summary>
        /// <param name="projectId">Requierd if not provided 400 will be returned</param>
        [HttpGet(ApiRoutes.Projects.Get)]
        public async Task<IActionResult> Get([FromRoute] string projectId)
        {

            if ( !String.IsNullOrEmpty(projectId))
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
        /// </summary>  s
        [HttpGet(ApiRoutes.Projects.GetAll)]
        public async Task<IActionResult> GetAll()
        {

            var projects = await _projectService.GetProjectsAsync();
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
        /// Update project endpoint
        /// </summary>
        /// <param name="projectId">Requierd to update project, if it dosent exists error will be returned</param>
        /// <param name="request">Request body check example, if field are not provided it wont be updated</param>  s
        [HttpPut(ApiRoutes.Projects.Update)]
        public async Task<IActionResult> Update([FromRoute] string projectId, [FromBody] UpdateProjectRequest request)
        {

            var project = await _projectService.GetProjectByIdAsync(projectId);
            if (project == null)
            {
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
            }

            project.UpdateAgianstUpdateRequest(request);

            var updated = await _projectService.UpdateProjectAsync(project);

            if (!updated)
                return NotFound();


            return Ok(project);
        }
        /// <summary>
        /// Deleted specifed project Auth is requierd.
        /// </summary>
        /// <param name="projectId">Is requierd if not provided 400 will be returned</param>    s
        [HttpDelete(ApiRoutes.Projects.Delete)]
        [Authorize(AuthenticationSchemes = JwtBearerDefaults.AuthenticationScheme, Roles = "Admin")]
        public async Task<IActionResult> Delete([FromRoute] string projectId)
        {
            if (!String.IsNullOrEmpty(projectId))
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
