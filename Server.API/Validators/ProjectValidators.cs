using Api.Contracts.V1.Requests.Project;
using FluentValidation;

namespace Api.Validators
{
    public class CreateProjectRequestValidator : AbstractValidator<CreateProjectRequest>
    {
        public CreateProjectRequestValidator()
        {
            RuleFor(x => x.ProjectId)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd);
            RuleFor(x => x.OrderId)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd);
            RuleFor(x => x.Name)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 - \\s]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Manager)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 -]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Sector)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 -]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Client)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 -]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Description)
                .NotNull()
                .WithMessage(ValidatorMessages.NotNull)
                .Matches("^[a-zA-Z0-9 -]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers)
                .MaximumLength(200)
                .WithMessage(ValidatorMessages.MaxLength);
        }
    }

    public class CreateVersionRequestValidator : AbstractValidator<CreateVersionRequest>
    {
        public CreateVersionRequestValidator()
        {
            RuleFor(x => x.Title)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 -]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Description)
                .MaximumLength(200)
                .WithMessage(ValidatorMessages.MaxLength);
        }
    }

    public class UpdateVersionRequestValidator : AbstractValidator<UpdateVersionRequest>
    {
        public UpdateVersionRequestValidator()
        {
            RuleFor(x => x.Title)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 -]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Description)
                .MaximumLength(200)
                .WithMessage(ValidatorMessages.MaxLength);
        }
    }

    public class UpdateProjectRequestValidator : AbstractValidator<UpdateProjectRequest>
    {
        public UpdateProjectRequestValidator()
        {
            RuleFor(x => x.OrderId)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd);
            RuleFor(x => x.Name)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 - \\s]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Manager)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 - \\s]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Sector)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 - \\s]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Description)
                .NotNull()
                .WithMessage(ValidatorMessages.NotNull)
                .Matches("^[a-zA-Z0-9 - \\s]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers)
                .MaximumLength(200)
                .WithMessage(ValidatorMessages.MaxLength);
        }
    }

}
