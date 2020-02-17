using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace Web.Models
{
    public class Project
    {
        [Display(Name = "Databas Id")]
        public int ProjectId { get; set; }
        [Display(Name = "Order Id")]
        public int OrderId { get; set; }
        [Display(Name = "Namn")]
        public string Name { get; set; }
        [Display(Name = "Beskrivning")]
        public string Description { get; set; }
        [Display(Name = "Manager")]
        public string Manager { get; set; }
        [Display(Name = "Klient")]
        public string Client { get; set; }
        [Display(Name = "Sektor")]
        public string Sector { get; set; }
        [Display(Name = "Co")]
        public float Co { get; set; }
        [DataType(DataType.Date)]
        public DateTime CreatedAt { get; set; }
    }
}
