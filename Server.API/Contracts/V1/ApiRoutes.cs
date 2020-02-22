namespace Api.Contracts.V1
{
    public static class ApiRoutes
    {

        private const string _Root = "api";
        private const string _Version = "v1";
        public const string Base = _Root + "/" + _Version;

        public static class Project
        {
            public const string GetProjects = Base + "/project";
            public const string GetProject = Base + "/project/{projectId}";
            public const string UpdateProject = Base + "/project/{projectId}";
            public const string DeleteProject = Base + "/project/{projectId}";
            public const string CreateProject = Base + "/project";

            public const string CreateVersion = Base + "/project/{projectId}/version";
            public const string GetVersion = Base + "/project/{projectId}/version/{versionId}";
            public const string GetVersions = Base + "/project/{projectId}/version";
            public const string UpdateVersion = Base + "/project/{projectId}/version/{versionId}";
            public const string DeleteVersion = Base + "/project/{projectId}/version/{versionId}";

        }

        public static class Component
        {
            public const string Create = Base + "/component";
            public const string Get = Base + "/component/{versionId}";
        }
    }
}
