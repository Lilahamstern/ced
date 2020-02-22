using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace Api
{
    public interface IInstaller
    {
        void InstallServices(IConfiguration configuration, IServiceCollection services);

    }
}
