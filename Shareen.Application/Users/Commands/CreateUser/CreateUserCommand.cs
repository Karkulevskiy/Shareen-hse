using MediatR;

namespace Shareen.Application.Users.Commands.CreateUser;

public class CreateUserCommand : IRequest<Guid>
{
    public string Name { get; set; }
}