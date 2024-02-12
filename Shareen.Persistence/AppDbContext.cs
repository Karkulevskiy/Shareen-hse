using Microsoft.EntityFrameworkCore;
using Shareen.Application.Interfaces;
using Shareen.Domain;
using Microsoft.EntityFrameworkCore.Sqlite;
namespace Shareen.Persistence;

/// <summary>
/// DataBase class with custom Model's
/// </summary>
public class AppDbContext : DbContext, IAppDbContext
{
    public DbSet<User> Users { get; set; }
    public DbSet<Lobby> Lobbies { get; set; }
    public DbSet<Chat> Chats { get; set; }
    
    public AppDbContext(DbContextOptions<AppDbContext> dbContext)
        : base(dbContext) { }
    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        //modelBuilder.ApplyConfiguration(new ChatConfiguration());
        modelBuilder.ApplyConfiguration(new UserConfiguration());
        modelBuilder.ApplyConfiguration(new LobbyConfiguration());
        base.OnModelCreating(modelBuilder);
    }
}