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

        [HttpGet(ApiRoutes.Components.GetVersions)]
        public async Task<IActionResult> GetVersions([FromRoute] int projectId)
        {
            var componentInformation = await _componentService.GetComponentInformationsAsync(projectId);
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

            var components = new List<ComponentInformation>();

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

            var versionInformation = await _componentService.GetComponentInformationByVersionAsync(projectId, request.Information.Version);
            Console.WriteLine(versionInformation);
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

            var componentResult = await _componentService.CreateComponentVersionAsync(projectId, request.Information);

            //var result = await _componentService.CreateComponentsAsync(components);
            
            // Needs to be checked over.
            //if (!result)
            //{
            //   return Problem("Failed to create components", "Internal server errror", (int)HttpStatusCode.InternalServerError, "Internal server error", "Testing");
            //}

            return Ok(new SuccessResponse{ Message = componentResult.ToString() });
        }
        
    }
}