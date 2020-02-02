using Api.Domain.Versions;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;


namespace Api.Domain.Components
{ 

    [Table("Component")]
    public class Component
    {
        [Key]
        public int Id { get; set; }
        [ForeignKey("Version")]
        public int PvId { get; set; }
        public int CId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public double Co { get; set; }
        public int Level { get; set; }
        public string Type { get; set; }
        public virtual Version Version  { get; set; }
    }
}
