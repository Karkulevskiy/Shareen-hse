namespace Shareen.Domain;

public class User
{
    public Guid Id { get; set; }
    public string Name { get; set; }
    public List<Lobby> Lobbies { get; set; }
}