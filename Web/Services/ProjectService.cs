using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using Newtonsoft;
using System.Threading.Tasks;
using Web.Helpers;
using Web.Models;
using Microsoft.AspNetCore.Components;
using Newtonsoft.Json;

namespace Web.Services
{
    public class ProjectService : IProjectService
    {
        public async Task<List<Project>> GetProjectsAsync()
        {
            var client = HttpClientHelper.GetHttpClient();
            var response = await client.GetAsync("");
            if(response.IsSuccessStatusCode)
            {
                var msg = await response.Content.ReadAsStringAsync();
                List<Project> projects = JsonConvert.DeserializeObject<List<Project>>(msg);
                return projects;
            }

            return new List<Project>();
        }

        public async Task<List<Project>> GetProjectsSearchAsync(string search)
        {
            var client = HttpClientHelper.GetHttpClient();
            var response = await client.GetAsync($"?search={search}");
            if (response.IsSuccessStatusCode)
            {
                var msg = await response.Content.ReadAsStringAsync();
                List<Project> projects = JsonConvert.DeserializeObject<List<Project>>(msg);
                return projects;
            }

            return new List<Project>();

        }
    }
}
