using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Domain.Projects
{
    public class ProjectResult
    {

        public bool Success { get; set; }

        public IEnumerable<string> Errors { get; set; }
    }
}
