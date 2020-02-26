using BusinessLayer.Models.Common;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace BusinessLayer.Models
{
    [Table("ProjectHistory")]
    public class ProjectHistory : BaseEntity
    {
        [Key]
        public int Id { get; set; }
        [ForeignKey("Project")]
        [JsonProperty(PropertyName = "projectId")]
        public int PId { get; set; }
        public string Property { get; set; }
        public string Data { get; set; }
        public Project Project { get; set; }
    }
}
