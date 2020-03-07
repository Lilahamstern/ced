using System.ComponentModel.DataAnnotations;

namespace Api.Contracts.V1.Requests.Project
{
    public class CreateVersionRequest
    {
        [Required]
        public string Title { get; set; }
        public string Description { get; set; }
    }
}
