using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Web.Data;

namespace Web.Services
{
    public interface IProjectService
    {
        public Task<List<Project>> FetchProjectsAsync();
    }
}
