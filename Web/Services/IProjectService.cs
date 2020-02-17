using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Web.Models;

namespace Web.Services
{
    public interface IProjectService
    {
        Task<List<Project>> GetProjectsAsync();
        Task<List<Project>> GetProjectsSearchAsync(string search);
    }
}
