using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace Web.Model
{
    public class Project
    {
        [Display(Name = "Namn")]
        public string Name { get; set; }
        [Display(Name = "Shit")]
        public string Manager { get; set; }
    }
}
