
using DataAccessLayer.Models.Common;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace DataAccessLayer.Models
{
    public class ProjectInformation : BaseEntity
    {
        public int Id { get; set; }
        [Required]
        public int OrderId { get; set; }
        [Required]
        [MaxLength(50)]
        public string Name { get; set; }
        [MaxLength(300)]
        public string Description { get; set; }
        [Required]
        [MaxLength(40)]
        public string Manager { get; set; }
        [Required]
        [MaxLength(50)]
        public string Client { get; set; }
        [Required]
        [MaxLength(50)]
        public string Sector { get; set; }
        [Required]
        public int ProjectId { get; set; }
        public List<Version> Versions { get; set; }
    }
}
