using Microsoft.AspNetCore.Components;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;
using Web.Data;

namespace Web.Services
{
    public class ProjectService : IProjectService
    {
        private readonly HttpClient _httpClient;
        public ProjectService(HttpClient httpClient)
        {
            _httpClient = httpClient;
        }
        public async Task<List<Project>> FetchProjectsAsync()
        {
            var projects = await _httpClient.GetJsonAsync<List<Project>>("");
            return projects;
        }
    }
}
