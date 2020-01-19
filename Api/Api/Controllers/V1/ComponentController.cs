using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Api.Contracts.V1;
using Api.Contracts.V1.Requests;
using Api.Contracts.V1.Responses;
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

        [HttpPost(ApiRoutes.Components.Create)]
        public async Task<IActionResult> Create([FromRoute] string projectId, [FromBody] List<CreateComponentRequest> request)
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
                            Message= "Project Id does not exist",
                        }
                    }
                });
            }

            return Ok(request);
        }
        
    }
}