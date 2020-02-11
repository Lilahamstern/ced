using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Runtime.Serialization;
using System.Threading.Tasks;

namespace Api.Domain.Projects
{

    [Table("ProjectHistory")]
    public class ProjectHistory
    {
        [Key]
        public int Id { get; set; }
        [ForeignKey("Project")]
        [JsonProperty(PropertyName = "projectId")]
        public int PId { get; set; }
        public string Property { get; set; }
        public string Data { get; set; }
        [DatabaseGenerated(DatabaseGeneratedOption.Identity), DataMember]
        public DateTime? UpdatedAt { get; set; }
        public Project Project { get; set; }
    }
}
