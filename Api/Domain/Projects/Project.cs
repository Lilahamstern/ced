using Api.Domain.Projects;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Runtime.Serialization;
using Version = Api.Domain.Versions.Version;

namespace Api.Domain.Projects
{
    [Table("Project")]
    public class Project
    {

        public Project()
        {
            CreatedAt = DateTime.UtcNow;
        }

        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.None)]
        [JsonProperty(PropertyName = "projectId")]
        public int PId { get; set; }
        [JsonProperty(PropertyName = "orderId")]
        public int OId {get; set;}
        public string Name { get; set; }
        public string Description { get; set; }
        public string Manager { get; set; }
        public string Client { get; set; }
        public string Sector { get; set; }
        [DatabaseGenerated(DatabaseGeneratedOption.Identity), DataMember]
        public System.DateTime? CreatedAt { get; private set; }
        public virtual ICollection<ProjectHistory> ProjectHistory { get; set; }
        public virtual ICollection<Version> Versions { get; set; }
    }
}
