using System;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;

namespace Client.Models
{
    public class ProjectModel
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
        [DisplayName("Skapad")]
        [DataType(DataType.Date)]
        public DateTime CreatedAt { get; set; }
    }
}
