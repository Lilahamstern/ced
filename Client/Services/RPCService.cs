using Grpc.Net.Client;
using Server.gRPC.Controllers;

namespace Client.Services
{
    public class RPCService : IRPCService
    {
        private readonly GrpcChannel _channel;
        public RPCService()
        {
            _channel = GrpcChannel.ForAddress("http://server-grpc/");
        }

        public Project.ProjectClient GetProjectClient()
        {
            return new Project.ProjectClient(_channel);
        }

        //Console.WriteLine("Now");
        //var channel = GrpcChannel.ForAddress("http://server-grpc/");
        //var client = new Greeter.GreeterClient(channel);
        //var reply = await client.SayHelloAsync(new HelloRequest { Name = message });
        //response = reply.Message;
        //Console.WriteLine($"{response} this is greeting");
        //Console.WriteLine(reply.Message);
        //Console.WriteLine(reply.CalculateSize().ToString());
    }
}
