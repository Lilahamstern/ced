using Api.Controllers.V1.Requests;
using FluentValidation;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Api.Validators
{
    public class CreateProjectRequestValidator : AbstractValidator<CreateProjectRequest>
    {
        public CreateProjectRequestValidator()
        {
            RuleFor(x => x.ProjectId)
                .NotEmpty();
            RuleFor(x => x.OrderId)
                .NotEmpty();
            RuleFor(x => x.Name)
                .NotEmpty()
                .Matches("^[a-zA-Z0-9 -]*$");
            RuleFor(x => x.Manager)
                .NotEmpty()
                .Matches("^[a-zA-Z0-9 -]*$");
            RuleFor(x => x.Sector)
                .NotEmpty()
                .Matches("^[a-zA-Z0-9 -]*$"); 
            RuleFor(x => x.Description)
                .NotEmpty()
                .Matches("^[a-zA-Z0-9 -]*$")
                .MaximumLength(200);
        }
    }

}
