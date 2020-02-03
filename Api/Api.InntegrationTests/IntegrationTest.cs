
using Api.Contracts.V1;
using Api.Contracts.V1.Requests;
using Api.Contracts.V1.Responses;
using Api.Data;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using Newtonsoft.Json;
using System;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using System.Threading.Tasks;


namespace Api.IntegrationTest
{
    public class IntegrationTest : IDisposable
    {

        protected readonly HttpClient TestClient;
        private readonly IServiceProvider _serviceProvider;
        protected IntegrationTest()
        {
            var appFactory = new WebApplicationFactory<Startup>()
                .WithWebHostBuilder(builder =>
                {
                    builder.ConfigureServices(s =>
                    {
                        s.RemoveAll(typeof(DataContext));
                        s.AddDbContext<DataContext>(o =>
                        {
                            o.UseInMemoryDatabase("TestDb");
                        });
                    });
                });

            _serviceProvider = appFactory.Services;
            TestClient = appFactory.CreateClient();
        }

        protected async Task AuthenticateAsync()
        {
            TestClient.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("bearer", await  GetJwtAsync());
        }

        protected async Task<ProjectResponse> CreateProjectAsync(CreateProjectRequest request)
        {
            var response = await TestClient.PostAsJsonAsync(ApiRoutes.Project.CreateProject, request);
            return await response.Content.ReadAsAsync<ProjectResponse>();
        }

        private async Task<string> GetJwtAsync()
        {


            var response = await TestClient.PostAsJsonAsync(ApiRoutes.Auth.Register, new UserRegistrationRequest
            {
                Email = "test@gmail.com",
                Password = "12345AbC!"
            });

            var regResponse = await response.Content.ReadAsAsync<AuthSuccessResponse>();

            return regResponse.Token;
        }

        public void Dispose()
        {
            using(var serviceScope = _serviceProvider.CreateScope())
            {
                var context = serviceScope.ServiceProvider.GetService<DataContext>();
                context.Database.EnsureDeleted();
            }
        }
    }
}
