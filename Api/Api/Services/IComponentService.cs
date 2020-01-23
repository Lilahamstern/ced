﻿using Api.Domain.Components;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IComponentService
    {
        Task<bool> CreateComponents(List<Component> components);
    }
}