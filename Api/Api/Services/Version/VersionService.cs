using Api.Data;
using System;
using System.Collections.Generic;
using System.Linq;
using Version = Api.Domain.Versions.Version;
using System.Threading.Tasks;
using System.Data.Entity;

namespace Api.Services
{
    public class VersionService : IVersionService
    { 
        private readonly DataContext _dataContext;

        public VersionService(DataContext dataContext)
        {
            _dataContext = dataContext;
        }

        public async Task<bool> CreateVersionAsync(Version version)
        {
            await _dataContext.Versions.AddAsync(version);

            var created = await _dataContext.SaveChangesAsync();
            return created > 0;

        }

        public Task<bool> DeleteVersionAsync(int versionId)
        {
            throw new NotImplementedException();
        }

        public async Task<Version> GetVersionByIdAsync(int versionId)
        {
            return await _dataContext.Versions.SingleOrDefaultAsync(x => x.Id == versionId);
        }

        public async Task<Version> GetVersionByTitleAsync(string title)
        {
            return await _dataContext.Versions.SingleOrDefaultAsync(x => x.Title == title);
        }

        public Task<List<Version>> GetVersionsAsync(int projectId)
        {
            throw new NotImplementedException();
        }
    }
}
