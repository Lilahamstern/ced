using BusinessLayer.Models.EntityFramework.Common;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace BusinessLayer.Models.EntityFramework
{
    public class Project : BaseEntity
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.None)]
        public int ProjectId { get; set; }
        public virtual IEnumerable<ProjectInformation> ProjectInformation { get; set; }
        public virtual IEnumerable<Version> Versions { get; set; }
    }
}
