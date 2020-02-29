using Api.Filters;
using Api.Services;
using FluentValidation.AspNetCore;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace Api.Installers
{
    public class MvcInstaller : IInstaller
    {

        public void InstallServices(IConfiguration configuration, IWebHostEnvironment env, IServiceCollection services)
        {

            services.AddScoped<IProjectService, ProjectService>();
            services.AddScoped<IVersionService, VersionService>();
            services.AddScoped<IComponentService, ComponentService>();

            services
                .AddMvc(o =>
                {
                    o.EnableEndpointRouting = false;
                    o.Filters.Add<ValidationFilter>();
                })
                .AddNewtonsoftJson(x =>
                {
                    x.SerializerSettings.ReferenceLoopHandling = Newtonsoft.Json.ReferenceLoopHandling.Ignore;
                    x.SerializerSettings.NullValueHandling = Newtonsoft.Json.NullValueHandling.Ignore;
                    x.SerializerSettings.Formatting = Newtonsoft.Json.Formatting.Indented;
                })
                .AddFluentValidation(conf => conf.RegisterValidatorsFromAssemblyContaining<Startup>())
                .SetCompatibilityVersion(CompatibilityVersion.Version_3_0);
            services.AddRazorPages();
        }
    }
}
