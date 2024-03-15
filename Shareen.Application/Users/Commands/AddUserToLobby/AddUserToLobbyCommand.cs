
using MediatR;
using Shareen.Application.Lobbies.Queries;

public class AddUserToLobbyCommand : IRequest<LobbyDto>
{
    public string UserName { get; set; }
    public string LobbyLink { get; set; }
}