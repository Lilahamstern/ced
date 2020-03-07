using Api.Contracts.V1.Responses.General;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Filters
{
    public class ValidationFilter : IAsyncActionFilter
    {
        public async Task OnActionExecutionAsync(ActionExecutingContext context, ActionExecutionDelegate next)
        {
            if (!context.ModelState.IsValid)
            {
                var errors = context.ModelState
                    .Where(x => x.Value.Errors.Count > 0)
                    .ToDictionary(kvp => kvp.Key, kvp => kvp.Value.Errors
                    .Select(x => x.ErrorMessage))
                    .ToArray();
                List<ErrorModel> errorModel = new List<ErrorModel>();

                foreach (var error in errors)
                {
                    foreach (var sub in error.Value)
                    {
                        errorModel.Add(new ErrorModel
                        (
                            error.Key.ToLower(),
                            sub
                        ));
                    }
                }

                context.Result = new BadRequestObjectResult(new ErrorResponse(errorModel));
                return;
            }

            await next();
        }
    }
}
