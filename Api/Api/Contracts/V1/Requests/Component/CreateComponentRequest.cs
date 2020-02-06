using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Requests.Component
{
    public class CreateComponentRequest
    {
        public int VersionId { get; set; }
        public List<ComponentRequest> Components { get; set; }
    }

    public class ComponentRequest
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
