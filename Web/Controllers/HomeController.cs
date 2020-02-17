using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using Web.Models;
using Web.Services;

namespace Web.Controllers
{
    public class HomeController : Controller
    {
        private readonly ILogger<HomeController> _logger;
        private readonly IProjectService _projectService;

        public HomeController(IProjectService projectService, ILogger<HomeController> logger)
        {
            _projectService = projectService;
            _logger = logger;
        }

        public async Task<IActionResult> Index(string search)
        {
            List<Project> projects;
            if (!String.IsNullOrEmpty(search))
            {
                projects = await _projectService.GetProjectsSearchAsync(search);
            } else
            {
                projects = await _projectService.GetProjectsAsync();
                if (projects.Count <= 0)
                {
                    return NotFound();
                }
            }
           

            var homeVM = new HomeViewModel
            {
                Projects = projects
            };

            return View(homeVM);
        }

        public IActionResult Privacy()
        {
            return View();
        }

        [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}
