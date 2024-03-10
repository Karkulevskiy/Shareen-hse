using MediatR;

namespace Shareen.Application.Lobbies.Queries.GetLobby;

public class GetLobbyQuery : IRequest<LobbyDto>
{
    public string Link { get; set; }
}