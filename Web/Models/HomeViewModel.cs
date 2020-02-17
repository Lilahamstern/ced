using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Web.Models
{
    public class HomeViewModel
    {
        public List<Project> Projects { get; set; }
        public string Search { get; set; }
    }
}
