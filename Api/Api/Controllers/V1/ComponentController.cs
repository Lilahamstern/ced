using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Threading.Tasks;
using Api.Contracts.V1;
using Api.Contracts.V1.Requests;
using Api.Contracts.V1.Responses;
using Api.Domain.Components;
using Api.Services;
using Microsoft.AspNetCore.Mvc;

namespace Api.Controllers.V1
{
    [Produces("application/json")]
    public class ComponentController : Controller
    {
        private readonly IComponentService _componentService;
        private readonly IProjectService _projectService;
        public ComponentController(IComponentService componentService, IProjectService projectService)
        {
            _componentService = componentService;
            _projectService = projectService;
        }
        /// <summary>
        /// Get project versions
        /// </summary>
        /// <param name="projectId">Is requierd</param>
        [HttpGet(ApiRoutes.Components.GetVersions)]
        public async Task<IActionResult> GetVersions([FromRoute] int projectId, [FromQuery] string version)
        {
            var componentInformation = await _componentService.GetProjectVersionsByProjectAsync(projectId);
            if (componentInformation == null)
            {
                return NotFound(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            Message= "Component versions could not be found",
                        }
                    }
                });
            }

            return Ok(componentInformation);
        }

        /// <summary>
        /// Create components and add to specified project
        /// </summary>
        /// <param name="projectId">ProjectID to what project components are gonna be added too</param>
        /// <param name="request">A list of components see model below</param>
        [HttpPost(ApiRoutes.Components.Create)]
        public async Task<IActionResult> Create([FromRoute] int projectId, [FromBody] CreateComponentRequest request)
        {

            var components = new List<Component>();

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
                            Message= "Project Id does not exist",
                        }
                    }
                });
            }

            var versionInformation = await _componentService.GetProjectVersionByVersionAsync(projectId, request.Information.Version);
            if (versionInformation != null)
            {
                return NotFound(new ErrorResponse
                {
                    Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            FieldName = "Version",
                            Message= "Component version already exists",
                        }
                    }
                });
            }

            var projectVersion = await _componentService.CreateProjectVersionAsync(projectId, request.Information);
            
            foreach(var component in request.Components)
            {
                var cd = new Component
                {
                    PvId = projectVersion,
                    CId = component.ComponentId,
                    Name = component.Name,
                    Type = component.Type,
                    Co = component.Co,
                    Level = component.Level,
                    Material = component.Material,
                    Profile = component.Profile,
                };

                components.Add(cd);
            }

            var componentResult = await _componentService.CreateComponentsAsync(components);


            var baseUrl = $"{HttpContext.Request.Scheme}://{HttpContext.Request.Host.ToUriComponent()}";
            var location = baseUrl + "/" + ApiRoutes.Components.GetVersions.Replace("{projectId}", projectId.ToString());

            var response = new SuccessResponse { Message = componentResult.ToString() };
            return Created(location, response);
        }

        [HttpDelete(ApiRoutes.Components.Delete)]
        public async Task<IActionResult> DeleteVersion([FromRoute] int versionId)
        {
            return Ok();
        }
        
    }
}