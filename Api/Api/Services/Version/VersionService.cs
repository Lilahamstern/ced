using Api.Data;
using System.Collections.Generic;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using Version = Api.Domain.Versions.Version;
using System.Threading.Tasks;

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

        public async Task<bool> DeleteVersionAsync(int projectId, int versionId)
        {
            return true;
        }

        public async Task<Version> GetVersionByTitleAsync(int projectId, string title)
        {
           return await (from v in _dataContext.Versions
                         where v.Title == title && v.PId == projectId
                         select v).FirstOrDefaultAsync();
        }

        public async Task<Version> GetVersionByIdAsync(int projectId, int versionId)
        {
            return await (from v in _dataContext.Versions
                         where v.Id == versionId && v.PId == projectId
                         select v).SingleOrDefaultAsync();
        }

        public async Task<List<Version>> GetVersionsAsync(int projectId)
        {
            var versions = await (from v in _dataContext.Versions
                          where v.PId == projectId
                          select v).ToListAsync();

            versions.ForEach((version) => version.Project = null);
            return versions;
        }
    }
}
