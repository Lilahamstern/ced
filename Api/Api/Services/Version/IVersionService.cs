using System;
using System.Collections.Generic;
using System.Linq;
using Version = Api.Domain.Versions.Version;
using System.Threading.Tasks;

namespace Api.Services
{
    public interface IVersionService
    {
        public Task<bool> CreateVersionAsync(Version version);
        public Task<bool> DeleteVersionAsync(int versionId);
        public Task<List<Version>> GetVersionsAsync(int projectId);
        public Task<Version> GetVersionByTitleAsync(string title);
        public Task<Version> GetVersionByIdAsync(int versionId);

    }
}
