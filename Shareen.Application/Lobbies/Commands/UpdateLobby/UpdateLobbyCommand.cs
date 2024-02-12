using MediatR;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Commands.UpdateLobby;

public class UpdateLobbyCommand : IRequest<Unit>
{
    public Guid Id { get; set; }
    public string Name { get; set; }
    public List<User> Users { get; set; }
}