using System;
using VersionModel = Api.Domain.Versions.Version;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1.Responses.Version
{
    public class VersionResponse
    {
        public VersionResponse(VersionModel version)
        {
            VersionId = version.Id;
            ProjectId = version.PId;
            Title = version.Title;
            Description = version.Description;
            CreatedAt = version.CreatedAt;
        }

        public VersionResponse(int versionId)
        {
            VersionId = versionId;
        }
        public int VersionId { get; set; }
        public int ProjectId { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        public DateTime? CreatedAt { get; set; } = null;
    }
}
