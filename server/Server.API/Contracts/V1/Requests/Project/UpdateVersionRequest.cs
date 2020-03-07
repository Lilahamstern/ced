using System.ComponentModel.DataAnnotations;

namespace Api.Contracts.V1.Requests.Project
{
    public class UpdateVersionRequest
    {
        [Required]
        public string Title { get; set; }
        public string Description { get; set; }
    }
}
