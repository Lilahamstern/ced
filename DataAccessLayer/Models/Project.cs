using DataAccessLayer.Models.Common;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace DataAccessLayer.Models
{
    public class Project : BaseEntity
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.None)]
        public int Id { get; set; }
        public List<ProjectInformation> ProjectInformation { get; set; } = new List<ProjectInformation>();
        public List<Version> Versions { get; set; } = new List<Version>();
    }
}
