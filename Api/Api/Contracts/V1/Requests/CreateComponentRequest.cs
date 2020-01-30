using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Requests
{
    public class CreateComponentRequest
    {
        public ComponentInformationRequest Information {get; set; }
        public List<ComponentDataRequest> Components { get; set; }
    }

    public class ComponentInformationRequest
    {
        public string Version { get; set; }
        public string Description { get; set; }
    }

    public class ComponentDataRequest
    {
        public int ComponentId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public double Co { get; set; }
        public int Level { get; set; }
        public string Type { get; set; }
    }
}
