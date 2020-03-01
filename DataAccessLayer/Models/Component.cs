using DataAccessLayer.Models.Common;
using System.ComponentModel.DataAnnotations;


namespace DataAccessLayer.Models
{
    public class Component : BaseEntity
    {
        [Key]
        public int Id { get; set; }
        [Required]
        public int ComponentId { get; set; }
        [Required]
        [MaxLength(50)]
        public string Name { get; set; }
        [Required]
        [MaxLength(30)]
        public string Profile { get; set; }
        [Required]
        [MaxLength(50)]
        public string Material { get; set; }
        [Required]
        public float Co { get; set; }
        [Required]
        public int Level { get; set; }
        [Required]
        [MaxLength(30)]
        public string Type { get; set; }
        [Required]
        public int VersionId { get; set; }
    }
}
