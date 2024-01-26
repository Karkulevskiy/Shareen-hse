using Microsoft.EntityFrameworkCore;
using Shareen.Domain;

namespace Shareen.Application.Interfaces;

public interface IAppDbContext
{
    public DbSet<Chat> Chats { get; set; }
    public DbSet<User> Users { get; set; }
    public DbSet<Lobby> Lobbies { get; set; }
    public Task<int> SaveChangesAsync(CancellationToken cancellationToken);
    //to - do остался вопрос с этим метадом, он переопределяется???
}