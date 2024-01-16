using MediatR;
using Shareen.Domain;

namespace Shareen.Application.Users.Commands.UpdateUser;

public class UpdateUserCommand : IRequest<Unit>
{
    public Guid Id { get; set; }
    public Guid LobbyId { get; set; }
    public Lobby? Lobby { get; set; }
    public string Name { get; set; }
}