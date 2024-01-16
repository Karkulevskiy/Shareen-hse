using MediatR;

namespace Shareen.Application.Users.Queries.GetUser;

public class GetUserQuery : IRequest<UserDto>
{
    public Guid Id { get; set; }
}