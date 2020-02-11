using Api.Domain.Components;
using Api.Domain.Projects;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Runtime.Serialization;
using System.Threading.Tasks;

namespace Api.Domain.Versions
{
    [Table("Version")]
    public class Version
    {
        public Version()
        {
            CreatedAt = DateTime.UtcNow;
        }

        [Key]
        public int Id { get; set; }
        [ForeignKey("Project")]
        [JsonProperty(PropertyName = "projectId")]
        public int PId { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        [DatabaseGenerated(DatabaseGeneratedOption.Identity), DataMember]
        public DateTime? CreatedAt { get; set; }
        [JsonProperty(ReferenceLoopHandling = ReferenceLoopHandling.Ignore)]
        public virtual Project Project { get; set; }
        public virtual ICollection<Component> Components { get; set; }
    }
}
