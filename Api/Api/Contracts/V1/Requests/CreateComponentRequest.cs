using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Requests
{
    public class CreateComponentRequest
    {
        public ComponentInformation Information {get; set; }
        public List<ComponentData> Components { get; set; }
    }

    public class ComponentInformation
    {
        public string Version { get; set; }
        public string Description { get; set; }
    }

    public class ComponentData
    {
        public int ComponentId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public double Co { get; set; }
        public string Type { get; set; }
    }
}
