using Api.Contracts.V1.Requests;
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
                Information = new ComponentInformation
                {
                    Description = "Version description, is not requierd",
                    Version = "Economy shit",
                },
                Components = new List<ComponentData>
                {
                   new ComponentData
                   {
                      ComponentId = 3126312,
                      Co = 341.32,
                      Name = "Wall 1",
                      Profile = "500x500",
                      Material = "Betong",
                      Type = "Vägg",
                   },
                   new ComponentData
                   {
                      ComponentId = 3122321,
                      Co = 521.32,
                      Name = "Golv",
                      Profile = "500x5000",
                      Material = "Betong",
                      Type = "Golv",
                   }
                }
            };
        }
    }
}
