using Api.Contracts.V1.Responses.General;
using Api.Contracts.V1.Responses.Project;
using Swashbuckle.AspNetCore.Filters;

namespace Api.SwaggerExample.Response
{
    public class ProjectResponseExample : IExamplesProvider<ProjectResponse>
    {
        public ProjectResponse GetExamples()
        {
            return new ProjectResponse(2301839);
        }
    }

    public class ErrorResponseExample : IExamplesProvider<ErrorResponse>
    {
        public ErrorResponse GetExamples()
        {
            return new ErrorResponse(
                        new ErrorModel
                        (
                            "Field name",
                            "Error message"
                        ));
        }
    };
}

