using System;
using System.Collections.Generic;
using System.Text;

namespace DataLibrary.Models
{
    public class ComponentModel
    {
        public int Id { get; set; }
        public int ComponentId { get; set; }
        public int VersionId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public float Co { get; set; }
        public int Level { get; set; }
        public string Type { get; set; }
        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }
    }
}
