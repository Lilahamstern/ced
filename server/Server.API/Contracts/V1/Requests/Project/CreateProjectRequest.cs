using System.ComponentModel.DataAnnotations;

namespace Api.Contracts.V1.Requests.Project
{
    public class CreateProjectRequest
    {
        [Required]
        public int ProjectId { get; set; }
        [Required]
        public int OrderId { get; set; }
        [Required]
        public string Name { get; set; }
        public string Description { get; set; }
        [Required]
        public string Manager { get; set; }
        [Required]
        public string Client { get; set; }
        [Required]
        public string Sector { get; set; }
    }
}
