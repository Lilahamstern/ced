using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using DataLibrary.DataAccess;

namespace Server.gRPC.Installers
{
    public class DatabaseInstaller : IInstaller
    {
        public void InstallServices(IConfiguration configuration, IServiceCollection services)
        {
            DataAccess.SetConnectionString(configuration.GetConnectionString("DefaultConnection"));
        }
    }
}
