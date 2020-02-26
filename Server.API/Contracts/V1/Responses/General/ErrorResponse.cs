using System.Collections.Generic;

namespace Api.Contracts.V1.Responses.General
{
    public class ErrorResponse
    {

        public ErrorResponse(List<ErrorModel> errors)
        {
            Errors = errors;
        }

        public ErrorResponse(ErrorModel error)
        {
            Error = error;
        }

        public ErrorModel Error { get; set; }
        public List<ErrorModel> Errors { get; set; }
    }
}
