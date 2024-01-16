namespace Shareen.Domain;

public class User
{
    public string Name { get; set; }
    public Guid Id { get; set; }
    public Lobby? Lobby { get; set; }
    public Guid LobbyId { get; set; }
}