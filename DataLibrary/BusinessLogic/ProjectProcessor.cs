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

            string sql = @"insert into project (Id, CreatedAt, UpdatedAt)
                        values (@Id, @CreatedAt, @UpdatedAt);";


            return DataAccess.DataAccess.SaveData(sql, data);
        }

        public static int ProjectExists(int projectId)
        {
            string sql = $"select COUNT(*) from project where id = @data";
            return DataAccess.DataAccess.QueryData<int>(sql, projectId);
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

            string sql = @"insert into projectInformation (ProjectId, OrderId, Name, Description, 
                    Manager, Client, Sector, CreatedAt, UpdatedAt)
                        values (@ProjectId, @OrderId, @Name, @Description, @Manager, @Client, @Sector, @CreatedAt, @UpdatedAt);";

            return DataAccess.DataAccess.SaveData(sql, data);
        }

        public static List<ProjectInformationModel> LoadProjects()
        {
            string sql = @"select ProjectId, OrderId, Name, Description, 
                    Manager, Client, Sector, CreatedAt, UpdatedAt from projectInformation LIMIT 10;";

            return DataAccess.DataAccess.LoadData<ProjectInformationModel>(sql);
        }
    }
}
