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

        //[Authorize(AuthenticationSchemes = JwtBearerDefaults.AuthenticationScheme, Roles = "Admin")]
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

        [HttpGet(ApiRoutes.Projects.Get)]
        public async Task<IActionResult> Get([FromRoute] string projectId)
        {
            var project = await _projectService.GetProjectByIdAsync(projectId);

            if (project == null)
            {
                return NotFound();
            }

            return Ok(project);
        }

        [HttpGet(ApiRoutes.Projects.GetAll)]
        public async Task<IActionResult> GetAll()
        {
            return Ok(await _projectService.GetProjectsAsync());
        }

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

            project.Name = request.Name;

            var updated = await _projectService.UpdateProjectAsync(project);

            if (!updated)
                return NotFound();


            return Ok(project);
        }

        [HttpDelete(ApiRoutes.Projects.Delete)]
        [Authorize(AuthenticationSchemes = JwtBearerDefaults.AuthenticationScheme, Roles = "Admin")]
        public async Task<IActionResult> Delete([FromRoute] string projectId)
        {
            var userOwn = await _projectService.UserOwnProjectAsync(projectId, HttpContext.GetUserId());

            if (!userOwn)
            {
                return BadRequest("You do not own this project");
            }

            var deleted = await _projectService.DeleteProjectAsync(projectId);

            if (deleted)
                return NoContent();

            return NotFound();
        }

    }
}
