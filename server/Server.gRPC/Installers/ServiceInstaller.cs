using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Server.gRPC.Services;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Server.gRPC.Installers
{
    public class ServiceInstaller : IInstaller
    {
        readonly string MyAllowSpecificOrigins = "AllowOrigin";
        public void InstallServices(IConfiguration configuration, IServiceCollection services)
        {
            services.AddCors(o =>
            {
                o.AddPolicy(MyAllowSpecificOrigins, 
                    b =>
                {
                    b.AllowAnyOrigin()
                    .AllowAnyHeader()
                    .AllowAnyMethod()
                    .WithExposedHeaders("Grpc-Status", "Grpc-Message");
                });
            });

            services.AddScoped<IProjectService, ProjectService>();
        }
    }
}
