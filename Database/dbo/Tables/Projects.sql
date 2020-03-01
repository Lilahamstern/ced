CREATE TABLE [dbo].[Projects] (
    [Id]        INT           NOT NULL,
    [CreatedAt] DATETIME2 (7) NULL,
    [UpdatedAt] DATETIME2 (7) NULL,
    CONSTRAINT [PK_Projects] PRIMARY KEY CLUSTERED ([Id] ASC)
);

