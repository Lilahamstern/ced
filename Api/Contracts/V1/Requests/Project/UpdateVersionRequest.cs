using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Requests.Project
{
    public class UpdateVersionRequest
    {
        [Required]
        public string Title { get; set; }
        public string Description { get; set; }
    }
}
