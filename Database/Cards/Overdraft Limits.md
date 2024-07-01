```sql
CREATE TABLE [Cards].[OverdraftLimits]
(
	[ID] [int] NOT NULL IDENTITY(1, 1),
	[CardID] [int] NOT NULL,
	[LimitValue] [numeric] (15, 2) NOT NULL,
	[EndDate] [date] NULL,
	[CreateDate] [datetime] NOT NULL CONSTRAINT [DF_OverdraftLimits_CreateDate] DEFAULT (getdate()),
	[CreateUserID] [int] NOT NULL,
	[ApproveDate] [datetime] NULL,
	[ApproveUserID] [int] NULL
)
```

Shows the limits for [[Cards]] based on its ID