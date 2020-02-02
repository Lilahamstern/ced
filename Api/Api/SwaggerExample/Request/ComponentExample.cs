using Api.Contracts.V1.Requests.Component;
using Swashbuckle.AspNetCore.Filters;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.SwaggerExample.Request
{
    public class CreateComponentExample : IExamplesProvider<CreateComponentRequest>
    {
        public CreateComponentRequest GetExamples()
        {
            return new CreateComponentRequest {
                Components = new List<ComponentDataRequest>
                {
                   new ComponentDataRequest
                   {
                      ComponentId = 3126312,
                      Co = 341.32,
                      Name = "Wall 1",
                      Profile = "500x500",
                      Material = "Betong",
                      Level = 0,
                      Type = "Vägg",
                   },
                   new ComponentDataRequest
                   {
                      ComponentId = 3122321,
                      Co = 521.32,
                      Name = "Golv",
                      Profile = "500x5000",
                      Material = "Betong",
                      Level = 0,
                      Type = "Golv",
                   }
                }
            };
        }
    }
}
