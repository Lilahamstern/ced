using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Api.Contracts.V1;
using Api.Contracts.V1.Requests.Version;
using Microsoft.AspNetCore.Mvc;
using Swashbuckle.AspNetCore.Annotations;

namespace Api.Controllers.V1
{

    [SwaggerTag("Version endpoints")]
    [Produces("application/json")]
    public class VersionController : Controller
    {
        [HttpPost(ApiRoutes.Version.Create)]
        public async Task<IActionResult> Create([FromRoute] int projectId, [FromBody] CreateVersionRequest request)
        {
            return Ok();
        }
    }
}