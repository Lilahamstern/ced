using BusinessLayer.Models.Common;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace BusinessLayer.Models
{
    [Table("Version")]
    public class Version : BaseEntity
    {

        [Key]
        public int Id { get; set; }
        [ForeignKey("Project")]
        [JsonProperty(PropertyName = "projectId")]
        public int PId { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        [JsonProperty(ReferenceLoopHandling = ReferenceLoopHandling.Ignore)]
        public virtual Project Project { get; set; }
        public virtual IEnumerable<Component> Components { get; set; }
    }
}
