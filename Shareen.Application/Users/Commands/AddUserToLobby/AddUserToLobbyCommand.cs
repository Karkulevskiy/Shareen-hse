
using MediatR;

public class AddUserToLobbyCommand : IRequest<Unit>
{
    public Guid UserId { get; set; }
    public string LobbyLink { get; set; }
}