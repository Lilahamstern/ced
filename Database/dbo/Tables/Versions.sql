CREATE TABLE [dbo].[Versions] (
    [Id]                   INT            IDENTITY (1, 1) NOT NULL,
    [CreatedAt]            DATETIME2 (7)  NULL,
    [UpdatedAt]            DATETIME2 (7)  NULL,
    [Title]                NVARCHAR (30)  NOT NULL,
    [Description]          NVARCHAR (200) NULL,
    [ProjectId]            INT            NOT NULL,
    [ProjectInformationId] INT            NOT NULL,
    CONSTRAINT [PK_Versions] PRIMARY KEY CLUSTERED ([Id] ASC),
    CONSTRAINT [FK_Versions_ProjectInformation_ProjectInformationId] FOREIGN KEY ([ProjectInformationId]) REFERENCES [dbo].[ProjectInformation] ([Id]) ON DELETE CASCADE,
    CONSTRAINT [FK_Versions_Projects_ProjectId] FOREIGN KEY ([ProjectId]) REFERENCES [dbo].[Projects] ([Id])
);


GO
CREATE NONCLUSTERED INDEX [IX_Versions_ProjectId]
    ON [dbo].[Versions]([ProjectId] ASC);


GO
CREATE NONCLUSTERED INDEX [IX_Versions_ProjectInformationId]
    ON [dbo].[Versions]([ProjectInformationId] ASC);

