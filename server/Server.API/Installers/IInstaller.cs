using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace Api
{
    public interface IInstaller
    {
        void InstallServices(IConfiguration configuration, IWebHostEnvironment env, IServiceCollection services);

    }
}
