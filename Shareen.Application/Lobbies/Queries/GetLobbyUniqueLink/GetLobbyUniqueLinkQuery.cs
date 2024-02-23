using MediatR;

public class GetLobbyUniqueLinkQuery : IRequest<string>
{
    public Guid LobbyId { get; set; }
}