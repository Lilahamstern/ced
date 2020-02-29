using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.OpenApi.Models;
using Swashbuckle.AspNetCore.Filters;
using System;
using System.Collections.Generic;
using System.IO;
using System.Reflection;

namespace Api.Installers
{
    public class SwaggerInstaller : IInstaller
    {
        public void InstallServices(IConfiguration configuration, IWebHostEnvironment env, IServiceCollection services)
        {
            services.AddSwaggerGen(x =>
            {
                x.SwaggerDoc("v1", new OpenApiInfo
                {
                    Title = "Ced API",
                    Version = "V1",
                    Description = "Rest-Api for Ced",
                    Contact = new OpenApiContact
                    {
                        Email = "support@hamsterapps.net",
                        Name = "Tech support",
                        Url = new Uri("https://hamsterapps.net")
                    }
                });

                x.EnableAnnotations(true);

                x.ExampleFilters();

                var xmlFile = $"{Assembly.GetExecutingAssembly().GetName().Name}.xml";
                var xmlPath = Path.Combine(AppContext.BaseDirectory, xmlFile);

                x.IncludeXmlComments(xmlPath);

            });

            services.AddSwaggerExamplesFromAssemblyOf<Startup>();
        }
    }
}
