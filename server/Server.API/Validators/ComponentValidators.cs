using Api.Contracts.V1.Requests.Component;
using FluentValidation;

namespace Api.Validators
{
    public class CreateComponentRequestValidator : AbstractValidator<CreateComponentRequest>
    {
        public CreateComponentRequestValidator()
        {
            RuleFor(x => x.VersionId)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd);
            RuleFor(x => x.Components)
                .Must(list => list.Count >= 4)
                .WithMessage(ValidatorMessages.MinimumArrayListLength);
        }
    }

    public class ComponentRequestValidator : AbstractValidator<ComponentRequest>
    {
        public ComponentRequestValidator()
        {
            RuleFor(x => x.ComponentId)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd);
            RuleFor(x => x.Name)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd);
            RuleFor(x => x.Profile)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd)
                .Matches("^[a-zA-Z0-9 -]*$")
                .WithMessage(ValidatorMessages.OnlyStringAndNumbers);
            RuleFor(x => x.Material)
                .NotEmpty()
                .WithMessage(ValidatorMessages.PropertyRequierd);
        }
    }
}
