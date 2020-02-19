using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Responses.User
{
    public class AuthFailResponse
    {
        public IEnumerable<string> Errors { get; set; }

    }
}
