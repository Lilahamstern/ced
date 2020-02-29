using Server.gRPC.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Client.Services
{
    public interface IRPCService
    {
        Project.ProjectClient GetProjectClient();
    }
}
