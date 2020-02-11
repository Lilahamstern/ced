using Api.Contracts.V1.Responses;
using Api.Contracts.V1.Responses.User;
using Swashbuckle.AspNetCore.Filters;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.SwaggerExample.Response
{
    public class AuthSuccessResponseExample : IExamplesProvider<AuthSuccessResponse>
    {
        public AuthSuccessResponse GetExamples()
        {
            return new AuthSuccessResponse
            {

                RefreshToken = "Refreshtoken for user",
                Token = "Token for user"
            };
        }
    }
}
