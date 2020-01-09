using Api.Contracts.V1;
using Api.Controllers.V1.Requests;
using Api.Controllers.V1.Responses;
using Api.Domain;
using Api.Extentions;
using Api.Services;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Threading.Tasks;

namespace Api.Controllers.V1
{
    [Authorize(AuthenticationSchemes = JwtBearerDefaults.AuthenticationScheme)]
    public class ProjectController : Controller
    {

        private readonly IProjectService _projectService;

        public ProjectController(IProjectService projectService)
        {
            _projectService = projectService;
        }

        [HttpPost(ApiRoutes.Projects.Create)]
        public async Task<IActionResult> Create([FromBody] CreateProjectRequest request)
        {
            var project = new Project { 
                Name = request.Name, 
                UserId = HttpContext.GetUserId()
            };

            await _projectService.CreateProjectAsync(project);

            var baseUrl = $"{HttpContext.Request.Scheme}://{HttpContext.Request.Host.ToUriComponent()}";
            var location = baseUrl + "/" + ApiRoutes.Projects.Get.Replace("{projectId}", project.Id.ToString());

            var response = new ProjectResponse { Id = project.Id };

            return Created(location, project);
        }

        [HttpGet(ApiRoutes.Projects.Get)]
        public async Task<IActionResult> Get([FromRoute] Guid projectId)
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
        public async Task<IActionResult> Update([FromRoute] Guid projectId, [FromBody] UpdateProjectRequest request)
        {

            var userOwn = await _projectService.UserOwnProjectAsync(projectId, HttpContext.GetUserId());

            if (!userOwn)
            {
                return BadRequest("You do not own this project");
            }

            var project = await _projectService.GetProjectByIdAsync(projectId);
            project.Name = request.Name;

            var updated = await _projectService.UpdateProjectAsync(project);

            if (updated)
                return Ok(project);

            return NotFound();
        }

        [HttpDelete(ApiRoutes.Projects.Delete)]
        public async Task<IActionResult> Delete([FromRoute] Guid projectId)
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
