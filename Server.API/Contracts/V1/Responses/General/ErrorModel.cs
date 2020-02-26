namespace Api.Contracts.V1.Responses.General
{
    public class ErrorModel
    {
        public ErrorModel(string fieldName, string message)
        {
            FieldName = fieldName;
            Message = message;
        }

        public string FieldName { get; set; }
        public string Message { get; set; }
    }
}
