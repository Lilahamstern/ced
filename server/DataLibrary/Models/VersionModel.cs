using System;
using System.Collections.Generic;
using System.Text;

namespace DataLibrary.Models
{
    public class VersionModel
    {
        public int Id { get; set; }
        public int ProjectId { get; set; }
        public int ProjectInformationId { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }
    }
}
