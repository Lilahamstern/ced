using Api.Data;
using System.Collections.Generic;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using Version = Api.Domain.Versions.Version;
using System.Threading.Tasks;
using Api.Domain.Components;
using System;

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
            var version = await (from v in _dataContext.Versions
                                 where v.Id == versionId && v.PId == projectId
                                 select v).FirstOrDefaultAsync();

            _dataContext.Versions.Remove(version);
            var deleted = await _dataContext.SaveChangesAsync();
            return deleted > 0;
        }

        public async Task<Version> GetVersionByTitleAsync(int projectId, string title)
        {
            var version = await (from v in _dataContext.Versions
                                 where v.Title == title && v.PId == projectId
                                 select v).FirstOrDefaultAsync();
            if (version == null)
                return null;

            version.Project = null;
            return version;
        }

        public async Task<Version> GetVersionByIdAsync(int versionId, int projectId)
        {
            var version = await (from v in _dataContext.Versions
                                 where v.Id == versionId && v.PId == projectId
                                 select v).FirstOrDefaultAsync();
            if (version == null)
                return null;

            version.Project = null;
            return version;
        }

        public async Task<Version> GetVersionByIdAsync(int versionId)
        {
            var version = await (from v in _dataContext.Versions
                                 where v.Id == versionId
                                 select v).FirstOrDefaultAsync();
            if (version == null)
                return null;

            version.Project = null;
            return version;
        }

        public async Task<bool> UpdateVersionAsync(int versionId, Version versionToUpdate)
        {
            var result = await _dataContext.Versions.SingleOrDefaultAsync(x => x.Id == versionId);
            if (result != null)
            {
                result.Title = versionToUpdate.Title;
                result.Description = versionToUpdate.Description;
                var updated = await _dataContext.SaveChangesAsync();
                return updated > 0;
            }
            return false;
        }

        public async Task<List<Version>> GetVersionsAsync(int projectId)
        {
           var versions = await (from v in _dataContext.Versions
                          where v.PId == projectId
                          select v).ToListAsync();

           versions.ForEach((version) => version.Project = null);
           return versions;
        }

        public async Task<float> GetNewestVersionCo(int projectId)
        {

                var query =  await (from v in _dataContext.Versions
                            join c in _dataContext.Components on v.Id equals c.PvId
                            where v.PId == projectId
                            select new { Version = v, Component = c }).ToListAsync();
                foreach (var item in query)
                {
                    Console.WriteLine($"{item.Component.Name} - {item.Version.CreatedAt}");
                }

                return 0.0f;
        }
    }
}
