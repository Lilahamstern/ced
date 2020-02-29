using DataAccessLayer;
using DataAccessLayer.Seeder;
using Microsoft.AspNetCore;
using Microsoft.AspNetCore.Hosting;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Storage;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Linq;
using System.Threading.Tasks;

namespace Api
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var host = CreateWebHostBuilder(args).Build();

            using (var scope = host.Services.CreateScope())
            {
                var services = scope.ServiceProvider;
                try
                {
                    var context = services.GetRequiredService<DataContext>();
                    context.Database.GetAppliedMigrations().ToList().ForEach(c =>
                    {
                        Console.WriteLine(c);
                    });
                    if (!context.Database.GetAppliedMigrations().Contains("20200226205043_InitalCreate"))
                    {
                        context.Database.Migrate();
                    }
                    
                    DataSeeder.Initialize(context);
                }
                catch (Exception e)
                {
                    Console.WriteLine($"Error occuerd while migrating and seed database error: {e.Message}");
                }
            }

            await host.RunAsync();
        }

        public static IWebHostBuilder CreateWebHostBuilder(string[] args) =>
            WebHost.CreateDefaultBuilder(args)
            .UseStartup<Startup>();
    }
}
