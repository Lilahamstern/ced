using Api.Controllers.V1.Requests;
using Microsoft.AspNetCore.Identity;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Domain
{
    [Table("Project")]
    public class Project
    {
        [Key]
        public int PId { get; set; }    
        public int OId {get; set;}
        public string Name { get; set; }
        public string? Description { get; set; }
        public string Manager { get; set; }
        public string Client { get; set; }
        public string Sector { get; set; }

        [DatabaseGenerated(DatabaseGeneratedOption.Computed)]
        public DateTime CreatedAt { get; set; }
    }
}
