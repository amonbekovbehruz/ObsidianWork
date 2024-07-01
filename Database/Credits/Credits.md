The database that contains all of the data about **credits**

### Histories Changes
```sql
CREATE TABLE [Credits].[HistoriesChanges]
(
	[CreditID] [int] NOT NULL,
	[TranchID] [int] NOT NULL,
	[ChangeDate] [datetime] NOT NULL,
	[ChangeTypeID] [tinyint] NOT NULL,
	[OldInterestRate] [numeric] (15, 2) NOT NULL,
	[OldEndDate] [datetime] NOT NULL,
	[Comment] [nvarchar] (150) COLLATE Cyrillic_General_CI_AS NOT NULL,
	[UserID] [int] NULL,
	[ReservesIgnore] [bit] NULL
)
ALTER TABLE [Credits].[HistoriesChanges] ADD
CONSTRAINT [CK_HistoriesChanges_ChangeTypeID] CHECK (([ChangeTypeID]>=(1) AND [ChangeTypeID]<=(8)))
CREATE NONCLUSTERED INDEX [IX_HistoriesChanges_ChangeDate] ON [Credits].[HistoriesChanges] ([ChangeDate], [ChangeTypeID]) INCLUDE ([CreditID], [ReservesIgnore], [TranchID])
```

