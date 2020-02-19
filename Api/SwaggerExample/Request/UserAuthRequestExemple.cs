using Api.Contracts.V1.Requests.User;
using Swashbuckle.AspNetCore.Filters;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.SwaggerExample.Request
{
    public class UserAuthRequestExemple : IExamplesProvider<UserRegistrationRequest>, IExamplesProvider<UserLoginRequest>
    {
        UserRegistrationRequest IExamplesProvider<UserRegistrationRequest>.GetExamples()
        {
            return new UserRegistrationRequest
            {
                Email = "user@wsp.com",
                Password = "Password123!",
            };
        }

        UserLoginRequest IExamplesProvider<UserLoginRequest>.GetExamples()
        {
            return new UserLoginRequest
            {
                Email = "user@wsp.com",
                Password = "Password123!",
            };
        }
    }
}
