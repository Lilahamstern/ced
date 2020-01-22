using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Domain.Projects
{
    [Table("ProjectHistory")]
    public class ProjectHistory
    {
        [Key]
        public int Id { get; set; }
        public string Property { get; set; }
        public string Data { get; set; }

        [DatabaseGenerated(DatabaseGeneratedOption.Computed)]
        public DateTime UpdatedAt { get; set; }
    }
}
