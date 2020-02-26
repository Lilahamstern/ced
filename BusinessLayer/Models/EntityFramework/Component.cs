﻿using BusinessLayer.Models.EntityFramework.Common;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace BusinessLayer.Models.EntityFramework
{
    public class Component : BaseEntity
    {
        [Key]
        public int Id { get; set; }
        public int ComponentId { get; set; }
        public string Name { get; set; }
        public string Profile { get; set; }
        public string Material { get; set; }
        public float Co { get; set; }
        public int Level { get; set; }
        public string Type { get; set; }
        public int VersionId { get; set; }
        public Version Version { get; set; }
    }
}