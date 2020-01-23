using Api.Controllers.V1.Requests;
using Api.Domain.Components;
using Api.Domain.Projects;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
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
        public ICollection<ProjectHistory> ProjectHistory { get; set; }
        public ICollection<Component> Component { get; set; }
    }

    [Table("ProjectHistory")]
    public class ProjectHistory
    {
        [Key]
        public int Id { get; set; }
        [ForeignKey("Project")]
        public int PId { get; set; }
        public string Property { get; set; }
        public string Data { get; set; }
        [DatabaseGenerated(DatabaseGeneratedOption.Computed)]
        public DateTime UpdatedAt { get; set; }
        public Project Project { get; set; }
    }
}
