using Api.Contracts.V1.Requests.Project;
using Api.Contracts.V1.Responses.General;
using Api.Contracts.V1.Responses.Project;
using Api.Contracts.V1.Responses.Version;
using Api.Services;
using BusinessLayer.Models;
using Microsoft.AspNetCore.Mvc;
using Swashbuckle.AspNetCore.Annotations;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Contracts.V1
{
    [SwaggerTag("Project requests")]
    [Produces("application/json")]
    public class ProjectController : Controller
    {

        private readonly IProjectService _projectService;
        private readonly IVersionService _versionService;

        public ProjectController(IProjectService projectService, IVersionService versionService)
        {
            _projectService = projectService;
            _versionService = versionService;
        }

        #region Project endpoints


        /// <summary>
        /// Create project endpoint
        /// </summary>
        /// <param name="request"> Example request body to create project</param>
        [HttpPost(ApiRoutes.Project.CreateProject)]
        [ProducesResponseType(typeof(ProjectResponse), 201)]
        public async Task<IActionResult> CreateProject([FromBody] CreateProjectRequest request)
        {
            var project = new Project
            {
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
                return BadRequest(new ErrorResponse(
                    new ErrorModel("projectId", "Project ID already exists")
                    ));
            }

            await _projectService.CreateProjectAsync(project);


            var baseUrl = $"{HttpContext.Request.Scheme}://{HttpContext.Request.Host.ToUriComponent()}";
            var location = baseUrl + "/" + ApiRoutes.Project.GetProject.Replace("{projectId}", project.PId.ToString());

            var response = new ProjectResponse(project.PId);

            return Created(location, response);
        }
        /// <summary>
        /// Get specifed project by ID.
        /// </summary>
        /// <param name="projectId">ProjectId to specifed project</param>
        [HttpGet(ApiRoutes.Project.GetProject)]
        public async Task<IActionResult> GetProject([FromRoute] int projectId)
        {
            await _versionService.GetNewestVersionCo(projectId);
            if (projectId == 0)
            {
                return BadRequest(new ErrorResponse(
                    new ErrorModel("projectId", "Project ID is requierd")
                    ));
            }

            var project = await _projectService.GetProjectByIdAsync(projectId);

            if (project == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));
            }

            return Ok(new ProjectResponse(project));
        }

        /// <summary>
        /// Returns all projects if no project is found 404 will be retunred s
        /// </summary>
        /// <param name="search">Provide search query of ProjectID, Name, OrderID, Client and Manager</param>
        /// <param name="limit">Set a limit of results you want</param> 
        [HttpGet(ApiRoutes.Project.GetProjects)]
        public async Task<IActionResult> GetProjects([FromQuery] string search, [FromQuery] int limit)
        {
            List<Project> projects;
            List<ProjectResponse> response = new List<ProjectResponse>();
            if (limit <= 0)
                limit = 10000;

            if (!string.IsNullOrEmpty(search))
            {
                projects = await _projectService.GetProjectsAsync(limit, search);
            }
            else
            {
                projects = await _projectService.GetProjectsAsync(limit);
            }

            if (projects.Count == 0)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel(null, "Projects not found")
                    ));
            }

            projects.ForEach(project =>
            {
                response.Add(new ProjectResponse(project));
            });
            return Ok(response);
        }

        /// <summary>
        /// Update specified project
        /// </summary>
        /// <param name="projectId">ProjectId to project</param>
        /// <param name="request">Example body to update project</param>
        [HttpPut(ApiRoutes.Project.UpdateProject)]
        public async Task<IActionResult> UpdateProject([FromRoute] int projectId, [FromBody] UpdateProjectRequest request)
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
                return NotFound(new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));

            return Ok(new ProjectResponse(project.PId));
        }
        /// <summary>
        /// Deleted specifed project.
        /// </summary>
        /// <param name="projectId">ProjectId to project</param>
        [HttpDelete(ApiRoutes.Project.DeleteProject)]
        //[Authorize(/*AuthenticationSchemes = JwtBearerDefaults.AuthenticationScheme,*/ Roles = "Admin")]
        public async Task<IActionResult> DeleteProject([FromRoute] int projectId)
        {


            var deleted = await _projectService.DeleteProjectAsync(projectId);

            if (deleted)
                return NoContent();

            return NotFound(
                new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));
        }

        #endregion

        #region Project Version endpoints

        /// <summary>
        /// Create a version.
        /// </summary>
        /// <param name="projectId">Project id</param>
        /// <param name="request">Request body</param>
        [HttpPost(ApiRoutes.Project.CreateVersion)]
        public async Task<IActionResult> CreateVersion([FromRoute] int projectId, [FromBody] CreateVersionRequest request)
        {
            var projectExists = await _projectService.GetProjectByIdAsync(projectId);
            if (projectExists == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));
            }

            var versionExists = await _versionService.GetVersionByTitleAsync(projectId, request.Title);
            if (versionExists != null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("title", "Version title already exists for this project")
                    ));
            };

            var version = new Version
            {
                PId = projectId,
                Title = request.Title,
                Description = request.Description
            };

            await _versionService.CreateVersionAsync(version);

            var baseUrl = $"{HttpContext.Request.Scheme}://{HttpContext.Request.Host.ToUriComponent()}";
            var location = baseUrl + "/" + ApiRoutes.Project.GetVersion.Replace("{versionId}", version.Id.ToString()).Replace("{projectId}", projectId.ToString());

            var response = new VersionResponse(version.Id);

            return Created(location, response);
        }

        /// <summary>
        /// Get version by id
        /// </summary>
        /// <param name="versionId">Version id</param>
        /// <param name="projectId">Project id</param>
        [HttpGet(ApiRoutes.Project.GetVersion)]
        public async Task<IActionResult> GetVersion([FromRoute] int projectId, [FromRoute] int versionId)
        {
            var projectExists = await _projectService.GetProjectByIdAsync(projectId);
            if (projectExists == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));
            }

            var version = await _versionService.GetVersionByIdAsync(versionId, projectId);
            if (version == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("versionId", "Version not found")
                    ));
            }

            return Ok(new VersionResponse(version));
        }

        /// <summary>
        /// Get versions on specifed project
        /// </summary>
        /// <param name="projectId">Project id</param>
        [HttpGet(ApiRoutes.Project.GetVersions)]
        public async Task<IActionResult> GetVersions([FromRoute] int projectId)
        {
            List<VersionResponse> response = new List<VersionResponse>();
            var projectExists = await _projectService.GetProjectByIdAsync(projectId);
            if (projectExists == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));
            }

            var versions = await _versionService.GetVersionsAsync(projectId);
            if (versions.Count <= 0)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel(null, "Versions not found")
                    ));
            }

            versions.ForEach(version =>
            {
                response.Add(new VersionResponse(version));
            });
            return Ok(response);
        }

        /// <summary>
        /// Update specifed version
        /// </summary>
        /// <param name="projectId">Project Id</param>
        /// <param name="versionId">Version Id</param>
        /// <param name="request">Request body</param>
        [HttpPut(ApiRoutes.Project.UpdateVersion)]
        public async Task<IActionResult> UpdateVersion([FromRoute] int projectId, [FromRoute] int versionId, [FromBody] UpdateVersionRequest request)
        {
            var version = await _versionService.GetVersionByIdAsync(versionId, projectId);
            if (version == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("versionID", "Version not found")
                    ));
            }

            var project = await _projectService.GetProjectByIdAsync(projectId);
            if (project == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));
            }

            version.Title = request.Title;
            version.Description = request.Description;

            await _versionService.UpdateVersionAsync(versionId, version);

            return Ok(new VersionResponse(version.Id));
        }

        [HttpDelete(ApiRoutes.Project.DeleteVersion)]
        public async Task<IActionResult> DeleteVersion([FromRoute] int versionId, int projectId)
        {
            var projectExists = await _projectService.GetProjectByIdAsync(projectId);
            if (projectExists == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("projectId", "Project not found")
                    ));
            }

            var version = await _versionService.GetVersionByIdAsync(versionId, projectId);
            if (version == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("versionId", "Version not found")
                    ));
            }
            await _versionService.DeleteVersionAsync(projectId, versionId);
            return NoContent();
        }

        #endregion
    }
}
