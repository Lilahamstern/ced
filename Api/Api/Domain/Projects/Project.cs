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
        public string PId { get; set; }

        public string OId {get; set;}
        
        public string Name { get; set; }

        public string Description { get; set; }

        public string Manager { get; set; }

        public string Client { get; set; }

        public string Sector { get; set; }

        public void UpdateAgianstUpdateRequest(UpdateProjectRequest request)
        { 
            if (!String.IsNullOrEmpty(request.Name)) this.Name = request.Name;
            if (!String.IsNullOrEmpty(request.OrderId)) this.OId = request.OrderId;
            if (!String.IsNullOrEmpty(request.Client)) this.Client = request.Client;
            if (!String.IsNullOrEmpty(request.Manager)) this.Manager = request.Manager;
            if (!String.IsNullOrEmpty(request.Description)) this.Description = request.Description;
            if (!String.IsNullOrEmpty(request.Sector)) this.Sector = request.Sector;
        }
    }
}
