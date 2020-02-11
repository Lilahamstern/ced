using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Threading.Tasks;
using Api.Contracts.V1;
using Api.Contracts.V1.Requests.Component;
using Api.Contracts.V1.Responses;
using Api.Contracts.V1.Responses.General;
using Api.Domain.Components;
using Api.Services;
using Microsoft.AspNetCore.Mvc;

namespace Api.Contracts.V1
{
    [Produces("application/json")]
    public class ComponentController : Controller
    {
        private readonly IComponentService _componentService;
        private readonly IVersionService _versionService;
        public ComponentController(IComponentService componentService, IVersionService versionService)
        {
            _componentService = componentService;
            _versionService = versionService;
        }

        /// <summary>
        /// Create components and add to specified project
        /// </summary>
        /// <param name="request">A list of components see model below</param>
        [HttpPost(ApiRoutes.Component.Create)]
        public async Task<IActionResult> AddComponents([FromBody] CreateComponentRequest request)
        {

            var version = await _versionService.GetVersionByIdAsync(request.VersionId);
            if (version == null)
            {
                return NotFound(new ErrorResponse(
                    new ErrorModel("versionId", "Version could not be found")
                    )
                );
            }

            var components = new List<Component>(); 

            foreach(var c in request.Components)
            {
                components.Add(new Component
                {
                    PvId = request.VersionId,
                    CId = c.ComponentId,
                    Name = c.Name,
                    Type = c.Type,
                    Co = c.Co,
                    Level = c.Level,
                    Material = c.Material,
                    Profile = c.Profile,
                });
            }

            await _componentService.AddComponentsAsync(components);

            var baseUrl = $"{HttpContext.Request.Scheme}://{HttpContext.Request.Host.ToUriComponent()}";
            var location = baseUrl + "/" + ApiRoutes.Component.Get.Replace("{versionId}", request.VersionId.ToString());

            var response = new SuccessResponse { Message = $"Added {components.Count} to version: {request.VersionId}" };
            return Created(location, response);
        }      
    }
}