using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Controllers.V1.Requests
{
    public class CreateProjectRequest
    {
        public string ProjectId { get; set; }

        public string OrderId { get; set; }

        public string Name { get; set; }

        public string Description { get; set; }

        public string Manager { get; set; }

        public string Client { get; set; }

        public string Sector { get; set; }
    }
}
