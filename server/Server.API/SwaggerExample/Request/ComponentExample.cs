using Api.Contracts.V1.Requests.Component;
using Swashbuckle.AspNetCore.Filters;
using System.Collections.Generic;

namespace Api.SwaggerExample.Request
{
    public class CreateComponentExample : IExamplesProvider<CreateComponentRequest>
    {
        public CreateComponentRequest GetExamples()
        {
            return new CreateComponentRequest
            {
                VersionId = 1,
                Components = new List<ComponentRequest>
                {
                   new ComponentRequest
                   {
                      ComponentId = 3126312,
                      Co = 341.32f,
                      Name = "Wall 1",
                      Profile = "500x500",
                      Material = "Betong",
                      Level = 0,
                      Type = "Vägg",
                   },
                   new ComponentRequest
                   {
                      ComponentId = 3122325,
                      Co = 430.32f,
                      Name = "Floor",
                      Profile = "500x5000",
                      Material = "Betong",
                      Level = 0,
                      Type = "Golv",
                   },
                   new ComponentRequest
                   {
                      ComponentId = 3122327,
                      Co = 521.32f,
                      Name = "Wall 2",
                      Profile = "300x500",
                      Material = "Betong",
                      Level = 0,
                      Type = "Golv",
                   },
                   new ComponentRequest
                   {
                      ComponentId = 3122321,
                      Co = 300.32f,
                      Name = "Wall 3",
                      Profile = "500x400",
                      Material = "Betong",
                      Level = 0,
                      Type = "Golv",
                   }
                }
            };
        }
    }
}
