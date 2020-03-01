using DataLibrary.DataAccess;
using DataLibrary.Models;
using System;
using System.Collections.Generic;
using System.Text;

namespace DataLibrary.BusinessLogic
{
    public static class ProjectProcessor
    {
        public static int  CreateProject(int projectId)
        {
            ProjectModel data = new ProjectModel
            {
                Id = projectId,
                CreatedAt = DateTime.UtcNow,
                UpdatedAt = DateTime.UtcNow,
            };

            string sql = @"insert into dbo.Project (Id, CreatedAt, UpdatedAt)
                        values (@Id, @CreatedAt, @UpdatedAt);";


            return SqlDataAccess.SaveData(sql, data);
        }

        public static int ProjectExists(int projectId)
        {
            string sql = $"select COUNT(*) from dbo.Project where id = @projectId";
            return SqlDataAccess.QueryData<int>(sql, projectId);
        }

        public static int CreateProjectInformation(int projectId, int orderId, string name, string description, string manager,
            string client, string sector)
        {
            ProjectInformationModel data = new ProjectInformationModel
            {
                ProjectId = projectId,
                OrderId = orderId,
                Name = name,
                Description = description,
                Manager = manager,
                Client = client,
                Sector = sector,
            };

            string sql = @"insert into dbo.ProjectInformation (ProjectId, OrderId, Name, Description, 
                    Manager, Client, Sector, CreatedAt, UpdatedAt)
                        values (@ProjectId, @OrderId, @Name, @Description, @Manager, @Client, @Sector, @CreatedAt, @UpdatedAt);";

            return SqlDataAccess.SaveData(sql, data);
        }

        public static List<ProjectInformationModel> LoadProjects()
        {
            string sql = @"select ProjectId, OrderId, Name, Description, 
                    Manager, Client, Sector, CreatedAt, UpdatedAt from dbo.ProjectInformation;";

            return SqlDataAccess.LoadData<ProjectInformationModel>(sql);
        }
    }
}
