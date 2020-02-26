using BusinessLayer.Models.EntityFramework.Common;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace BusinessLayer.Models.EntityFramework
{
    public class ProjectInformation : BaseEntity
    {
        [Key]
        public int ProjectInformationId { get; set; }
        public int OrderId { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public string Manager { get; set; }
        public string Client { get; set; }
        public string Sector { get; set; }
        public int ProjectId { get; set; }
        public Project Project { get; set; }
        public virtual IEnumerable<Version> Versions { get; set; }
    }
}
