using MediatR;

public class GetLobbyUsersQuery : IRequest<LobbyUsersListDto>
{
    public Guid LobbyId { get; set; }
}