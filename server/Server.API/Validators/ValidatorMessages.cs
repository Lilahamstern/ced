namespace Api.Validators
{
    public static class ValidatorMessages
    {
        public const string PropertyRequierd = "{PropertyName} is requierd";
        public const string MaxLength = "Maximum length of {PropertyName} is {MaxLength} characters";
        public const string NotNull = "{PropertyName} cannot be null";
        public const string OnlyStringAndNumbers = "{PropertyName} can only contain letter and numbers";

        public const string MinimumArrayListLength = "{PropertyName} need to contain atleast 4 components";
    }
}
