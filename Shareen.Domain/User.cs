﻿namespace Shareen.Domain;

public class User
{
    // посмотреть как реализованы связи в EF
    public string Name { get; set; }
    public Guid Id { get; set; }
    public Lobby? Lobby { get; set; }
    public Guid? LobbyId { get; set; }
}