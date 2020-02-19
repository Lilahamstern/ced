using Api.Contracts.V1.Requests.Project;
using Swashbuckle.AspNetCore.Filters;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.SwaggerExample.Request
{
    public class VersionCreateExample : IExamplesProvider<CreateVersionRequest>
    {
        public CreateVersionRequest GetExamples()
        {
            return new CreateVersionRequest
            {
                Title = "Version 1",
                Description = "This is a version 1 that does 1+1"
            };
        }
    }

    public class UpdateVersionRequestExample : IExamplesProvider<UpdateVersionRequest>
    {
        public UpdateVersionRequest GetExamples()
        {
            return new UpdateVersionRequest
            {
                Title = "Updated title",
                Description = "This is a new description"
            };
        }
    }
}
