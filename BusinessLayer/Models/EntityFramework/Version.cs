using BusinessLayer.Models.EntityFramework.Common;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Text;

namespace BusinessLayer.Models.EntityFramework
{
    public class Version : BaseEntity
    {
        [Key]
        public int Id { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }

    }
}
