using Api.Contracts.V1.Requests.Project;
using Swashbuckle.AspNetCore.Filters;

namespace Api.SwaggerExample.Request
{
    public class CreateProjectRequestExample : IExamplesProvider<CreateProjectRequest>
    {
        public CreateProjectRequest GetExamples()
        {
            return new CreateProjectRequest
            {
                ProjectId = 20190105,
                OrderId = 3123810,
                Client = "Volvo AB",
                Name = "House with roof",
                Manager = "John Doe",
                Description = "A house with roof to volvo AB",
                Sector = "Houses",
            };
        }
    }

    public class UpdateProjectRequestExample : IExamplesProvider<UpdateProjectRequest>
    {
        public UpdateProjectRequest GetExamples()
        {
            return new UpdateProjectRequest
            {
                OrderId = 3223810,
                Client = "Toys R rush",
                Name = "House without roof",
                Manager = "John Doe",
                Description = "A house with roof to volvo AB",
                Sector = "Houses",
            };
        }

    }
}
