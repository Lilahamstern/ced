using BusinessLayer.Models.EntityFramework;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IComponentService
    {
        Task<bool> AddComponentsAsync(List<Component> components);
    }
}
