using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Requests.Project
{
    public class UpdateProjectRequest
    {
        public int OrderId { get; set; }

        public string Name { get; set; }

        public string Description { get; set; }

        public string Manager { get; set; }

        public string Client { get; set; }

        public string Sector { get; set; }
    }
}
