```csharp
namespace OnlineBank.Enums.Common
{
    /// <summary>
    /// Расширенный тип проводки
    /// </summary> 
    public enum ExtendedTransactionType : byte
    {
        /// <summary>
        /// Депозитная проводка
        /// </summary>
        [EnumDescription("Депозиты")]
        Deposit = 50,

        /// <summary>
        /// Кредитная проводка
        /// </summary>
        [EnumDescription("Кредиты")]
        Credit = 100,

        /// <summary>
        /// Кассовая проводка
        /// </summary>
        [EnumDescription("Касса")]
        Cash = 150,

        /// <summary>
        /// Проводка по депозитарию
        /// </summary>
        [EnumDescription("Депозитарий")]
        Depository = 200,

        /// <summary>
        /// Проводки по РКО
        /// </summary>
        [EnumDescription("РКО")]
        RKO = 90,

        /// <summary>
        /// Проводки по бухгалтерии / главной книге
        /// </summary>
        [EnumDescription("Главная книга")]
        GeneralBook = 250,

        /// <summary>
        /// Операции по пластиковым картам
        /// </summary>
        [EnumDescription("Пластиковые карты")]
        Cards = 30,
    }
}
```
