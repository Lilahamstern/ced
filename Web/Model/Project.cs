using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace Web.Model
{
    public class Project
    {
        [DisplayName("Uppdrags ID")]
        public int OrderId { get; set; }
        [DisplayName("Namn")]
        public string Name { get; set; }
        [DisplayName("Manager")]
        public string Manager { get; set; }
        [DisplayName("Klient")]
        public string Client { get; set; }
        [DisplayName("Sektor")]
        public string Sector { get; set; }
        [DisplayName("Co")]
        public float Co { get; set; }
        [DisplayName("Databas Id")]
        public int ProjectId { get; set; }
        [DisplayName("Skapad pa")]
        [DataType(DataType.Date)]
        public DateTime CreatedAt { get; set; }
    }
}
