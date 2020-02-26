using BusinessLayer.Models;
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
                        PId = 233641273,
                        OId = 312651,
                        Name = "Seed Testing 1",
                        Description = "Seed testing 1 description is useless just so u know. I dont event know why I am writing this so long, maybe just to test the ui later :P",
                        Client = "Seeding Company AB",
                        Manager = "John Dope",
                        Sector = "Seeding building",
                    },
                    new Project{
                        PId = 456839876,
                        OId = 904712,
                        Name = "Seed Testing 2",
                        Description = "Seed testing 1 description is useless just so u know :P",
                        Client = "Seeding Company AB",
                        Manager = "Keyboard Super Dope",
                        Sector = "Seeding Enviorment",
                    },
                    new Project{
                        PId = 78378123,
                        OId = 57912,
                        Name = "Seed Testing 3",
                        Description = "",
                        Client = "Fake Airborders",
                        Manager = "John Doe",
                        Sector = "Health of Seeding",
                    }

                };

                context.Projects.AddRange(projects);
                context.SaveChanges();
            }

            if (!context.Versions.Any())
            {
                var versions = new List<Version>()
                {
                    new Version
                    {
                        PId = 78378123,
                        Title = "Super Dope version",
                        Description = "A super dope version with a long description that are useless unless u wanna test the ui without adding new projects all the time."                   
                    },
                    new Version
                    {
                        PId =78378123,
                        Title = "Not so dope",
                        Description = "",
                    },
                    new Version
                    {
                        PId = 456839876,
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
                        VId = 3,
                        CId = 461233721,
                        Level = 0,
                        Material = "Wood",
                        Name = "Super wall",
                        Profile = "400x400",
                        Co = 3712.3412f,
                        Type = "Type"
                    },
                    new Component
                    {
                        VId = 3,
                        CId = 570791312,
                        Level = 0,
                        Material = "Wood",
                        Name = "Super wall XY",
                        Profile = "400x800",
                        Co = 5712.3412f,
                        Type = "Type"
                    },
                    new Component
                    {
                        VId = 3,
                        CId = 723712374,
                        Level = 0,
                        Material = "Stone",
                        Name = "Floor",
                        Profile = "400x400",
                        Co = 2112.3412f,
                        Type = "Type"
                    },
                    new Component
                    {
                        VId = 3,
                        CId = 457213031,
                        Level = 0,
                        Material = "Wood",
                        Name = "Roof",
                        Profile = "400x400x200",
                        Co = 3712.3412f,
                        Type = "Type"
                    },
                };
                context.Components.AddRange(components);
                context.SaveChanges();
            }
        }
    }
}
