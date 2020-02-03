using System;
using Version = Api.Domain.Versions.Version;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Api.Contracts.V1;
using Api.Contracts.V1.Requests.Version;
using Api.Contracts.V1.Responses;
using Api.Services;
using Microsoft.AspNetCore.Mvc;
using Swashbuckle.AspNetCore.Annotations;

namespace Api.Controllers.V1
{

    [SwaggerTag("Version endpoints")]
    [Produces("application/json")]
    public class VersionController : Controller
    {
        private IProjectService _projectService;
        private IVersionService _versionService;
        public VersionController(IVersionService versionService, IProjectService projectService)
        {
            _versionService = versionService;
            _projectService = projectService;
        }

        /// <summary>
        /// Create a version.
        /// </summary>
        /// <param name="projectId">Project id</param>
        /// <param name="request">Request body</param>
        [HttpPost(ApiRoutes.Version.Create)]
        public async Task<IActionResult> Create([FromRoute] int projectId, [FromBody] CreateVersionRequest request)
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
                            FieldName = "projectId",
                            Message = "Project not found.",
                        }
                    }
                });
            }

            var version = new Version
            {
                PId = projectId,
                Title = request.Title,
                Description = request.Description
            };

            await _versionService.CreateVersionAsync(version);

            var baseUrl = $"{HttpContext.Request.Scheme}://{HttpContext.Request.Host.ToUriComponent()}";
            var location = baseUrl + "/" + ApiRoutes.Version.Get.Replace("{versionId}", version.Id.ToString());

            var response = new VersionResponse { VersionId = version.Id };

            return Created(location, response);
        }

        [HttpGet(ApiRoutes.Version.Get)]
        public async Task<IActionResult> Get([FromRoute] int versionId)
        {
            var version = await _versionService.GetVersionByIdAsync(versionId);
            if (version == null)
            {
                return NotFound(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            FieldName = "versionId",
                            Message = "Version could not be found"
                        }
                    }
                });
            }

            return Ok(version);
        }

    }
}