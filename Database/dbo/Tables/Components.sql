CREATE TABLE [dbo].[Components] (
    [Id]          INT           IDENTITY (1, 1) NOT NULL,
    [CreatedAt]   DATETIME2 (7) NULL,
    [UpdatedAt]   DATETIME2 (7) NULL,
    [ComponentId] INT           NOT NULL,
    [Name]        NVARCHAR (50) NOT NULL,
    [Profile]     NVARCHAR (30) NOT NULL,
    [Material]    NVARCHAR (50) NOT NULL,
    [Co]          REAL          NOT NULL,
    [Level]       INT           NOT NULL,
    [Type]        NVARCHAR (30) NOT NULL,
    [VersionId]   INT           NOT NULL,
    CONSTRAINT [PK_Components] PRIMARY KEY CLUSTERED ([Id] ASC),
    CONSTRAINT [FK_Components_Versions_VersionId] FOREIGN KEY ([VersionId]) REFERENCES [dbo].[Versions] ([Id]) ON DELETE CASCADE
);


GO
CREATE NONCLUSTERED INDEX [IX_Components_VersionId]
    ON [dbo].[Components]([VersionId] ASC);

