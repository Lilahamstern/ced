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
        /// Create components and add to specified project
        /// </summary>
        /// <param name="projectId">ProjectID to what project components are gonna be added too</param>
        /// <param name="request">A list of components see model below</param>
        [HttpPost(ApiRoutes.Components.Create)]
        public async Task<IActionResult> Create([FromRoute] int projectId, [FromBody] List<CreateComponentRequest> request)
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

            foreach (var component in request)
            {
                components.Add(
                    new Component
                    {
                        PId = projectId,
                        CId = component.CId,
                        Material = component.Material,
                        Name = component.Name,
                        Profile = component.Profile,
                        Type = component.Type,
                        Version = 1,
                        Co = component.Co,
                    });
            }

            var result = await _componentService.CreateComponents(components);
            
            // Needs to be checked over.
            if (!result)
            {
               return Problem("Failed to create components", "Internal server errror", (int)HttpStatusCode.InternalServerError, "Internal server error", "Testing");
            }

            return Ok(new SuccessResponse{ Message = "Created components" });
        }
        
    }
}