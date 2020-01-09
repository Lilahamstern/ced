using Api.Contracts.V1;
using Api.Controllers.V1.Requests;
using Api.Domain;
using FluentAssertions;
using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace Api.IntegrationTest
{
    public class PostControllerTests : IntegrationTest
    {
        [Fact]
        public async Task GetAll_WithoutAnyProjects_ReturnsEmptyResponse()
        {
            // Arrage
            await AuthenticateAsync();
            // Act
            var response = await TestClient.GetAsync(ApiRoutes.Projects.GetAll);
            // Assert
            response.StatusCode.Should().Be(HttpStatusCode.OK);
            (await response.Content.ReadAsAsync<List<Project>>()).Should().BeEmpty();
        }

        [Fact]
        public async Task Get_ReturnsProject_WhenProjectExistsInTheDatabase()
        {
            // Arrage
            await AuthenticateAsync();
            var createdProject = await CreateProjectAsync(new CreateProjectRequest { Name = "Test Project" });

            // act
            var response = await TestClient.GetAsync(ApiRoutes.Projects.Get.Replace("{projectId}", createdProject.Id.ToString()));

            // Assert
            response.StatusCode.Should().Be(HttpStatusCode.OK);
            var returnedProject = await response.Content.ReadAsAsync<Project>();
            returnedProject.Id.Should().Be(createdProject.Id);
            returnedProject.Name.Should().Be("Test Project");
        }
    }
}
