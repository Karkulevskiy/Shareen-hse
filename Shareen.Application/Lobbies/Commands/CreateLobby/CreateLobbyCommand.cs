using MediatR;

namespace Shareen.Application.Lobbies.Commands.CreateLobby;

public class CreateLobbyCommand : IRequest<CreateLobbyDto>
{
    public string Name { get; set; }    
}