using BusinessLayer.Models.EntityFramework;
using System.Collections.Generic;
using System.Linq;


namespace DataAccessLayer.Seeder
{
    public class DataSeeder
    {
        public static void Initialize(DataContext context)
        {
            if (!context.Projects.Any())
            {
                var projects = new List<Project>()
                {
                    new Project{
                        ProjectId = 373091
                    },
                    new Project{
                        ProjectId = 571290
                    },
                    new Project{
                        ProjectId = 928731
                    }

                };

                context.Projects.AddRange(projects);
                context.SaveChanges();
            }

            if (!context.ProjectInformation.Any())
            {
                var projectInformation = new List<ProjectInformation>()
                {
                    new ProjectInformation
                    {
                        ProjectId = 373091,
                        OrderId = 940971,
                        Name = "Helipad 2000s",
                        Description = "A new helipad at volvo with a useless description that is prob gonna be to long, for the ui",
                        Manager = "John Dope",
                        Client = "Volvo",
                        Sector = "Construction",
                    },
                    new ProjectInformation
                    {
                        ProjectId = 373091,
                        OrderId = 940971,
                        Name = "Helipad 3000s",
                        Description = "A new helipad at volvo with a useless description that is prob gonna be to long, for the ui, with some new values for the second version",
                        Manager = "John Dope",
                        Client = "Volvo",
                        Sector = "Construction and healthcare",
                    },
                    new ProjectInformation
                    {
                        ProjectId = 571290,
                        OrderId = 473921,
                        Name = "Car Production Hall",
                        Description = "",
                        Manager = "Sara Dope",
                        Client = "Ikea",
                        Sector = "School",
                    },
                };

                context.ProjectInformation.AddRange(projectInformation);
                context.SaveChanges();
            }

            if (!context.Versions.Any())
            {
                var versions = new List<Version>()
                {
                    new Version
                    {
                        ProjectId = 373091,
                        ProjectInformationId = 1,
                        Title = "Money Cash for Helipad",
                        Description = "A super dope version with a long description that are useless unless u wanna test the ui without adding new projects all the time."                   
                    },
                    new Version
                    {
                        ProjectId = 373091,
                        ProjectInformationId = 2,
                        Title = "Helipad v2",
                        Description = "",
                    },
                    new Version
                    {
                        ProjectId = 571290,
                        ProjectInformationId = 3,
                        Title = "Useless Version",
                        Description = "This is a a useless version so dont ask why"
                    }
                };
                context.Versions.AddRange(versions);
                context.SaveChanges();
            }

            if (!context.Components.Any())
            {
                var components = new List<Component>()
                {
                    new Component
                    {
                        VersionId = 4,
                        ComponentId = 461233,
                        Level = 0,
                        Material = "Concreat",
                        Name = "Helipad left leg",
                        Profile = "400x400",
                        Co = 3712.3412f,
                        Type = "Type"
                    },
                    new Component
                    {
                        VersionId = 4,
                        ComponentId = 477409,
                        Level = 0,
                        Material = "Concreat",
                        Name = "Helipad right leg",
                        Profile = "200x400",
                        Co = 3000.3412f,
                        Type = "Type"
                    },
                    new Component
                    {
                        VersionId = 4,
                        ComponentId = 389123,
                        Level = 0,
                        Material = "Concreat",
                        Name = "Helipad Middle leg",
                        Profile = "200x200",
                        Co = 2000f,
                        Type = "Type"
                    },
                    new Component
                    {
                        VersionId = 4,
                        ComponentId = 472818,
                        Level = 0,
                        Material = "Steel",
                        Name = "Helipad plate",
                        Profile = "2000x2000",
                        Co = 53012.312f,
                        Type = "Type"
                    },
                };
                context.Components.AddRange(components);
                context.SaveChanges();
            }
        }
    }
}
