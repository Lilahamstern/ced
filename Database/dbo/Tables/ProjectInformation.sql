CREATE TABLE [dbo].[ProjectInformation] (
    [Id]          INT            IDENTITY (1, 1) NOT NULL,
    [CreatedAt]   DATETIME2 (7)  NULL,
    [UpdatedAt]   DATETIME2 (7)  NULL,
    [OrderId]     INT            NOT NULL,
    [Name]        NVARCHAR (50)  NOT NULL,
    [Description] NVARCHAR (300) NULL,
    [Manager]     NVARCHAR (40)  NOT NULL,
    [Client]      NVARCHAR (50)  NOT NULL,
    [Sector]      NVARCHAR (50)  NOT NULL,
    [ProjectId]   INT            NOT NULL,
    CONSTRAINT [PK_ProjectInformation] PRIMARY KEY CLUSTERED ([Id] ASC),
    CONSTRAINT [FK_ProjectInformation_Projects_ProjectId] FOREIGN KEY ([ProjectId]) REFERENCES [dbo].[Projects] ([Id]) ON DELETE CASCADE
);


GO
CREATE NONCLUSTERED INDEX [IX_ProjectInformation_ProjectId]
    ON [dbo].[ProjectInformation]([ProjectId] ASC);

