using Shareen.Domain;

namespace Shareen.Application.Lobbies.Queries;

public class LobbyDto
{
    public string Name { get; set; }
    public int NumberOfPeople { get; set; }
    public DateTime TimeCreated { get; set; }
    public List<User> Users { get; set; }
}