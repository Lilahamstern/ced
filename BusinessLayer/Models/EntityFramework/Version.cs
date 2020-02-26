using BusinessLayer.Models.EntityFramework.Common;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace BusinessLayer.Models.EntityFramework
{
    public class Version : BaseEntity
    {
        [Key]
        public int VersionId { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        public int ProjectId { get; set; }
        public Project Project { get; set; }
        public int ProjectInformationId { get; set; }
        public ProjectInformation ProjectInformation { get; set; }
        public virtual IEnumerable<Component> Components { get; set; }

    }
}
