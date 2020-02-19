using Microsoft.AspNetCore.Authorization;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Authorization
{
    public class CompanyEmailRequirement : IAuthorizationRequirement
    {
        public string DomainName { get; set; }

        public CompanyEmailRequirement(string domainName)
        {
            DomainName = domainName;
        }

    }
}
