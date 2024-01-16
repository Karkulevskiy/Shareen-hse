using MediatR;

namespace Shareen.Application.Lobbies.Commands.CreateLobby;

public class CreateLobbyCommand : IRequest<Guid>
{
    public string Name { get; set; }    
}