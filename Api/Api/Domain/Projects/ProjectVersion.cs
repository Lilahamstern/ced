using Api.Domain.Components;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Runtime.Serialization;
using System.Threading.Tasks;

namespace Api.Domain.Projects
{
    [Table("ProjectVersion")]
    public class ProjectVersion
    {
        public ProjectVersion()
        {
            CreatedAt = DateTime.UtcNow;
        }

        [Key]
        public int Id { get; set; }
        [ForeignKey("Project")]
        public int PId { get; set; }
        public string Version { get; set; }
        public string Description { get; set; }
        [DatabaseGenerated(DatabaseGeneratedOption.Identity), DataMember]
        public DateTime? CreatedAt { get; set; }
        public virtual Project Project { get; set; }
        public virtual ICollection<Component> Components { get; set; }
    }
}
