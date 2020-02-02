using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Contracts.V1
{
    public static class ApiRoutes
    {

        private const string _Root = "api";
        private const string _Version = "v1";
        public const string Base = _Root + "/" + _Version;

        public static class Project
        {
            public const string GetAll = Base + "/project";
            public const string Get = Base + "/project/{projectId}";
            public const string Update = Base + "/project/{projectId}";
            public const string Delete = Base + "/project/{projectId}";
            public const string Create = Base + "/project";
        }

        public static class Component
        {
            public const string Create = Base + "/component/{versionId}";
        }

        public static class Version
        {
            public const string Get = Base + "/version/{projectId}";
            public const string Create = Base + "/version/{projectId}";
            public const string Delete = Base + "/version/{versionId}";
        }

        public static class Auth
        {
            public const string Login = Base + "/auth/login";
            public const string Register = Base + "/auth/register";
            public const string Refresh = Base + "/auth/refresh";
        }
    }
}
