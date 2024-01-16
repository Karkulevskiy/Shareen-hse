namespace Shareen.Domain;

public class Chat
{
    public Guid Id { get; set; }
    public Guid LobbyId { get; set; }
    public Lobby Lobby { get; set; }
    public List<string> ListMessages { get; set; }
}