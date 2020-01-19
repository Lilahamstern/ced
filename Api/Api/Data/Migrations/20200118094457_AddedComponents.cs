using System;
using Microsoft.EntityFrameworkCore.Migrations;

namespace Api.Data.Migrations
{
    public partial class AddedComponents : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Components",
                columns: table => new
                {
                    Id = table.Column<string>(nullable: false),
                    PId = table.Column<string>(nullable: true),
                    CId = table.Column<int>(nullable: false),
                    Name = table.Column<string>(nullable: true),
                    Profile = table.Column<string>(nullable: true),
                    Material = table.Column<string>(nullable: true),
                    Co = table.Column<string>(nullable: true),
                    Type = table.Column<string>(nullable: true),
                    Version = table.Column<int>(nullable: false),
                    CreatedAt = table.Column<DateTime>(nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Components", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Components_Project_PId",
                        column: x => x.PId,
                        principalTable: "Project",
                        principalColumn: "PId",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateIndex(
                name: "IX_Components_PId",
                table: "Components",
                column: "PId");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Components");
        }
    }
}
