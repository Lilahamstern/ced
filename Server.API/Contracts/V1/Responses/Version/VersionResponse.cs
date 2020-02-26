using System;
using VersionModel = BusinessLayer.Models.EntityFramework.Version;

namespace Api.Contracts.V1.Responses.Version
{
    public class VersionResponse
    {
        public VersionResponse(VersionModel version)
        {
            VersionId = version.VersionId;
            ProjectId = version.ProjectId;
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
