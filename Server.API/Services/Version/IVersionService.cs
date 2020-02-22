using BusinessLayer.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IVersionService
    {
        public Task<bool> CreateVersionAsync(Version version);

        public Task<bool> UpdateVersionAsync(int versionId, Version versionToUpdate);
        public Task<bool> DeleteVersionAsync(int projectId, int versionId);
        public Task<List<Version>> GetVersionsAsync(int projectId);
        public Task<Version> GetVersionByTitleAsync(int projectId, string title);
        public Task<Version> GetVersionByIdAsync(int versionId, int projectId);
        public Task<Version> GetVersionByIdAsync(int versionId);
        public Task<float> GetNewestVersionCo(int projectId);

    }
}
