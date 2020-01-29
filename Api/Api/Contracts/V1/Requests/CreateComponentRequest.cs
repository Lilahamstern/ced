using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Requests
{
    public class CreateComponentRequest
    {
        public string Version { get; set; }
        public string Description { get; set; }
        public List<CreateComponentDataRequest> Components { get; set; }
    }

    public class CreateComponentDataRequest
    {
        public int CId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public string Co { get; set; }
        public string Type { get; set; }
    }
}
