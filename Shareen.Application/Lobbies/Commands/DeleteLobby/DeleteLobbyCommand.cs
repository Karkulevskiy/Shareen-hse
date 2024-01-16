using MediatR;

namespace Shareen.Application.Lobbies.Commands.DeleteLobby;

public class DeleteLobbyCommand : IRequest<Unit>
{
    public Guid Id { get; set; }
}