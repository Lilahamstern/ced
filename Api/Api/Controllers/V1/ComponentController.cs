using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Threading.Tasks;
using Api.Contracts.V1;
using Api.Contracts.V1.Requests.Component;
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
        [HttpPost(ApiRoutes.Component.Create)]
        public async Task<IActionResult> Create([FromRoute] int projectId, [FromBody] CreateComponentRequest request)
        {

            var components = new List<Component>(); 

          
            
            foreach(var component in request.Components)
            {
                var cd = new Component
                {
                    PvId = projectId,
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
            var location = baseUrl + "/" + ApiRoutes.Version.Get.Replace("{projectId}", projectId.ToString());

            var response = new SuccessResponse { Message = componentResult.ToString() };
            return Created(location, response);
        }      
    }
}