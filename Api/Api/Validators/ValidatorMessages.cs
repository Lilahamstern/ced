using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Validators
{
    public static class ValidatorMessages
    {
        public const string PropertyRequierd = "{PropertyName} is requierd";
        public const string MaxLength = "Maximum length of {PropertyName} is {MaxLength} characters";
        public const string NotNull = "{PropertyName} cannot be null";
        public const string OnlyStringAndNumbers = "{PropertyName} can only contain letter and numbers";
    }
}
