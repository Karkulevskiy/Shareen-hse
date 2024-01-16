using MediatR;

namespace Shareen.Application.Users.Queries.GetUsersList;

public class GetUsersListQuery : IRequest<UsersListVm>
{
    public Guid LobbyId { get; set; }
}