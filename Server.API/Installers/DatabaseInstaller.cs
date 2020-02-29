using DataAccessLayer;
using Microsoft.AspNetCore.Hosting;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;


namespace Api.Installers
{
    public class DatabaseInstaller : IInstaller
    {
        public void InstallServices(IConfiguration configuration, IWebHostEnvironment env,IServiceCollection services)
        {
            services.AddDbContext<DataContext>(o => o.UseSqlServer(configuration.GetConnectionString("DefaultConnection")));
        }
    }
}
