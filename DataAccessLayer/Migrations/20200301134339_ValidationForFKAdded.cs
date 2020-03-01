using Microsoft.EntityFrameworkCore.Migrations;

namespace DataAccessLayer.Migrations
{
    public partial class ValidationForFKAdded : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Components_Versions_VersionId",
                table: "Components");

            migrationBuilder.DropForeignKey(
                name: "FK_ProjectInformation_Projects_ProjectId",
                table: "ProjectInformation");

            migrationBuilder.DropForeignKey(
                name: "FK_Versions_Projects_ProjectId",
                table: "Versions");

            migrationBuilder.DropForeignKey(
                name: "FK_Versions_ProjectInformation_ProjectInformationId",
                table: "Versions");

            migrationBuilder.AlterColumn<int>(
                name: "ProjectInformationId",
                table: "Versions",
                nullable: false,
                oldClrType: typeof(int),
                oldType: "int",
                oldNullable: true);

            migrationBuilder.AlterColumn<int>(
                name: "ProjectId",
                table: "Versions",
                nullable: false,
                oldClrType: typeof(int),
                oldType: "int",
                oldNullable: true);

            migrationBuilder.AlterColumn<int>(
                name: "ProjectId",
                table: "ProjectInformation",
                nullable: false,
                oldClrType: typeof(int),
                oldType: "int",
                oldNullable: true);

            migrationBuilder.AlterColumn<int>(
                name: "VersionId",
                table: "Components",
                nullable: false,
                oldClrType: typeof(int),
                oldType: "int",
                oldNullable: true);

            migrationBuilder.AddForeignKey(
                name: "FK_Components_Versions_VersionId",
                table: "Components",
                column: "VersionId",
                principalTable: "Versions",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_ProjectInformation_Projects_ProjectId",
                table: "ProjectInformation",
                column: "ProjectId",
                principalTable: "Projects",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_Versions_Projects_ProjectId",
                table: "Versions",
                column: "ProjectId",
                principalTable: "Projects",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);

            migrationBuilder.AddForeignKey(
                name: "FK_Versions_ProjectInformation_ProjectInformationId",
                table: "Versions",
                column: "ProjectInformationId",
                principalTable: "ProjectInformation",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Components_Versions_VersionId",
                table: "Components");

            migrationBuilder.DropForeignKey(
                name: "FK_ProjectInformation_Projects_ProjectId",
                table: "ProjectInformation");

            migrationBuilder.DropForeignKey(
                name: "FK_Versions_Projects_ProjectId",
                table: "Versions");

            migrationBuilder.DropForeignKey(
                name: "FK_Versions_ProjectInformation_ProjectInformationId",
                table: "Versions");

            migrationBuilder.AlterColumn<int>(
                name: "ProjectInformationId",
                table: "Versions",
                type: "int",
                nullable: true,
                oldClrType: typeof(int));

            migrationBuilder.AlterColumn<int>(
                name: "ProjectId",
                table: "Versions",
                type: "int",
                nullable: true,
                oldClrType: typeof(int));

            migrationBuilder.AlterColumn<int>(
                name: "ProjectId",
                table: "ProjectInformation",
                type: "int",
                nullable: true,
                oldClrType: typeof(int));

            migrationBuilder.AlterColumn<int>(
                name: "VersionId",
                table: "Components",
                type: "int",
                nullable: true,
                oldClrType: typeof(int));

            migrationBuilder.AddForeignKey(
                name: "FK_Components_Versions_VersionId",
                table: "Components",
                column: "VersionId",
                principalTable: "Versions",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);

            migrationBuilder.AddForeignKey(
                name: "FK_ProjectInformation_Projects_ProjectId",
                table: "ProjectInformation",
                column: "ProjectId",
                principalTable: "Projects",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);

            migrationBuilder.AddForeignKey(
                name: "FK_Versions_Projects_ProjectId",
                table: "Versions",
                column: "ProjectId",
                principalTable: "Projects",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);

            migrationBuilder.AddForeignKey(
                name: "FK_Versions_ProjectInformation_ProjectInformationId",
                table: "Versions",
                column: "ProjectInformationId",
                principalTable: "ProjectInformation",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);
        }
    }
}
