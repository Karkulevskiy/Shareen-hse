using Microsoft.EntityFrameworkCore;
using Shareen.Domain;

namespace Shareen.Application.Interfaces;

public interface IAppDbContext
{
    public DbSet<Chat> Chats { get; set; }
    public DbSet<User> Users { get; set; }
    public DbSet<Lobby> Lobbies { get; set; }

    Task SaveChangesAsync(CancellationToken cancellationToken);
    //to-do: узнать почему при Task пропадает ошибка при реализации интерфейса
}