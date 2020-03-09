using System;
using System.Collections.Generic;
using System.Text;

namespace DataLibrary.Models
{
    public class ProjectInformationModel
    {
        public int Id { get; set; }
        public int ProjectId { get; set; }
        public int OrderId { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public string Manager { get; set; }
        public string Client { get; set; }
        public string Sector { get; set; }
        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }
    }
}
