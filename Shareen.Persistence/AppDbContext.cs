using Microsoft.EntityFrameworkCore;
using Shareen.Application.Interfaces;
using Shareen.Domain;
namespace Shareen.Persistence;

public class AppDbContext : DbContext, IAppDbContext
{
    public DbSet<User> Users { get; set; }
    public DbSet<Lobby> Lobbies { get; set; }
    public DbSet<Chat> Chats { get; set; }
    
    public AppDbContext(DbContextOptions<AppDbContext> dbContext)
        : base(dbContext) { }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);
    }
}