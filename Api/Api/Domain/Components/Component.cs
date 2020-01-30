using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Runtime.Serialization;
using System.Threading.Tasks;

namespace Api.Domain.Components
{
    [Table("Component")]
    public class ComponentInformation
    {
        public ComponentInformation()
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
        public virtual ICollection<ComponentData> ComponentDatas { get; set; }
    }


    [Table("ComponentData")]
    public class ComponentData
    {
        [Key]
        public int Id { get; set; }
        [ForeignKey("ComponentInformation")]
        public int CId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public float Co { get; set; }
        public int Level { get; set; }
        public string Type { get; set; }
        public virtual ComponentInformation Component { get; set; }
    }
}
