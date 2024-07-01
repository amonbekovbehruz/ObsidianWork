```sql
CREATE TABLE [Cards].[Cards]
(
	[CardID] [int] NOT NULL IDENTITY(1, 1),
	[CompanyID] [int] NOT NULL,
	[CardNo] [nvarchar] (50) COLLATE Cyrillic_General_CI_AS NULL,
	[CardType] [int] NOT NULL CONSTRAINT [DF_Account_CardType] DEFAULT ((1)),
	[Codeword] [nvarchar] (20) COLLATE Cyrillic_General_CI_AS NULL,
	[CustomerNameForPrint] [nvarchar] (24) COLLATE Cyrillic_General_CI_AS NOT NULL,
	[CustomerCategoryID] [int] NULL,
	[UseSmsInforming] [bit] NOT NULL,
	[AutoLoad] [bit] NOT NULL CONSTRAINT [DF_Account_AutoLoad] DEFAULT ((0)),
	[PromoCode] [nvarchar] (50) COLLATE Cyrillic_General_CI_AS NULL,
	[OpenUserID] [int] NOT NULL,
	[CloseUserID] [int] NULL,
	[LastChangeUserID] [int] NOT NULL,
	[ProductID] [int] NOT NULL,
	[EMail] [nvarchar] (100) COLLATE Cyrillic_General_CI_AS NULL,
	[Phone] [nvarchar] (15) COLLATE Cyrillic_General_CI_AS NULL,
	[RiskGroupID] [int] NULL,
	[SourceInformationID] [int] NULL,
	[CreditID] [int] NULL
)
```

