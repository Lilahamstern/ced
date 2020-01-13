using Api.Controllers.V1.Requests;
using Swashbuckle.AspNetCore.Filters;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.SwaggerExample.Request
{
    public class CreateProjectRequestExample : IExamplesProvider<CreateProjectRequest>
    {
        public CreateProjectRequest GetExamples()
        {
            return new CreateProjectRequest
            {
                ProjectId = "201905031630",
                OrderId = "37412374",
                Client = "Volvo AB",
                Name = "House with roof",
                Manager = "John Doe",
                Description = "A house with roof to volvo AB",
                Sector = "Houses",
            };
        }
    }
}
