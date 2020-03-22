
using DataLibrary.Models;
using System;
using System.Collections.Generic;
using static DataLibrary.DataAccess.DataAccess;

namespace DataLibrary.BusinessLogic
{
    public static class ProjectProcessor
    {
        public static int CreateProject(int projectId)
        {
            ProjectModel data = new ProjectModel
            {
                Id = projectId,
            };

            string sql = @"insert into project (id)
                        values (@Id);";


            return SaveData(sql, data);
        }

        public static int ProjectExists(int projectId)
        {
            string sql = $"select COUNT(*) from project where id = @data";
            return QueryData(sql, projectId);
        }

        public static int CreateProjectInformation(int projectId, int orderId, string name, string description, string manager,
            string client, string sector)
        {
            ProjectInformationModel data = new ProjectInformationModel
            {
                Project_id = projectId,
                Order_id = orderId,
                Name = name,
                Description = description,
                Manager = manager,
                Client = client,
                Sector = sector,
            };

            string sql = @"insert into project_information (project_id, order_id, name, description, 
                    manager, client, sector)
                        values (@ProjectId, @OrderId, @Name, @Description, @Manager, @Client, @Sector);";

            return SaveData(sql, data);
        }

        public static List<ProjectInformationModel> LoadProjects(string query)
        {
            string sql = @"SELECT * FROM get_projects(@data)";

            return LoadData<ProjectInformationModel>(sql, query);
        }
    }
}
