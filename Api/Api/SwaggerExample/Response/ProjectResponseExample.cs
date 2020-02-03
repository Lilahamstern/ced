using Api.Contracts.V1.Responses;
using Swashbuckle.AspNetCore.Filters;
using System.Collections.Generic;

namespace Api.SwaggerExample.Response
{
    public class ProjectResponseExample : IExamplesProvider<ProjectResponse>
    {
        public ProjectResponse GetExamples()
        {
            return new ProjectResponse
            {
                ProjectId = 2301839,
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
