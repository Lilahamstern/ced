using BusinessLayer.Models.Common;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace BusinessLayer.Models
{
    [Table("Project")]
    public class Project : BaseEntity
    {

        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.None)]
        [JsonProperty(PropertyName = "projectId")]
        public int PId { get; set; }
        [JsonProperty(PropertyName = "orderId")]
        public int OId { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public string Manager { get; set; }
        public string Client { get; set; }
        public string Sector { get; set; }
        public virtual IEnumerable<ProjectHistory> ProjectHistory { get; set; }
        public virtual IEnumerable<Version> Versions { get; set; }
    }
}
