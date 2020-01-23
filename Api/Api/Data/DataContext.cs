
using Api.Domain;
using Api.Domain.Components;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;

namespace Api.Data
{
    public class DataContext : IdentityDbContext
    {
        public DataContext(DbContextOptions<DataContext> options)
            : base(options)
        {
        }

        public DbSet<Project> Projects { get; set; }

        public DbSet<ProjectHistory> projectHistories { get; set; }

        public DbSet<Component> Components { get; set; }

        public DbSet<ComponentData> ComponentDatas { get; set; }

        public DbSet<RefreshToken> RefreshTokens { get; set; }
    }
}
