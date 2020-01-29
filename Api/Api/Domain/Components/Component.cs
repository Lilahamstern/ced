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
    public class Component
    {
        [Key]
        public int Id { get; set; }
        [ForeignKey("Project")]
        public int PId { get; set; }
        public string Version { get; set; }
        public string Description { get; set; }
        [DatabaseGenerated(DatabaseGeneratedOption.Identity), DataMember]
        public DateTime? CreatedAt { get; set; }
        public Project Project { get; set; }
        public ICollection<ComponentData> ComponentDatas { get; set; }
    }


    [Table("ComponentData")]
    public class ComponentData
    {
        [Key]
        public int Id { get; set; }
        [ForeignKey("Component")]
        public int CId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public float Co { get; set; }
        public string Type { get; set; }
        public Component Component { get; set; }
    }
}
