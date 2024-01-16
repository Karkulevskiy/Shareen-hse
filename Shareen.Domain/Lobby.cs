namespace Shareen.Domain;

public class Lobby
{
    public Guid Id { get; set; }
    public string Name { get; set; }
    public Guid ChatId { get; set; }
    public Chat Chat { get; set; }
    public DateTime TimeCreated { get; set; }
    public int NumberOfUsers { get; set; }
    public List<User> Users { get; set; }
}