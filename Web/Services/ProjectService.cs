using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using Newtonsoft;
using System.Threading.Tasks;
using Web.Helpers;
using Web.Models;
using Microsoft.AspNetCore.Components;

namespace Web.Services
{
    public class ProjectService : IProjectService
    {
        public async Task<List<Project>> GetProjectsAsync()
        {
            var client = HttpClientHelper.GetHttpClient();
            var response = await client.GetJsonAsync<List<Project>>("");

            return response;
        }
    }
}
