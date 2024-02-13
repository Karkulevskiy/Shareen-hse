namespace Shareen.Domain;

public class Lobby
{
    public Guid Id { get; set; }
    public Guid ChatId { get; set; }
    public Chat Chat { get; set; }
    public List<User> Users { get; set; }
    public string Name { get; set; }
    public DateTime TimeCreated { get; set; }
    public string UniqueLink { get; set; }
}