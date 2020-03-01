using DataAccessLayer.Models.Common;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;


namespace DataAccessLayer.Models 
{ 
    public class Version : BaseEntity
    {
        [Key]
        public int Id { get; set; }
        [Required]
        [MaxLength(30)]
        public string Title { get; set; }
        [MaxLength(200)]
        public string Description { get; set; }
        [Required]
        public int ProjectId { get; set; }
        [Required]
        public int ProjectInformationId { get; set; }
        public List<Component> Components { get; set; } = new List<Component>();

    }
}
