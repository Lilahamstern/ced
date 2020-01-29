using Api.Contracts.V1.Responses;
using Api.Controllers.V1.Responses;
using Swashbuckle.AspNetCore.Filters;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.SwaggerExample.Response
{
    public class ProjectResponseExample : IExamplesProvider<ProjectResponse>
    {
        public ProjectResponse GetExamples()
        {
            return new ProjectResponse
            {
                PId = 2301839,
            };
        }
    }

    public class ErrorResponseExample : IExamplesProvider<ErrorResponse>
    {
        public ErrorResponse GetExamples()
        {
            return new ErrorResponse
            {
                Errors = new List<ErrorModel>
                    {
                        new ErrorModel
                        {
                            FieldName = "Field name",
                            Message = "Error message",
                        }
                    }
            };
        }
    }
}
