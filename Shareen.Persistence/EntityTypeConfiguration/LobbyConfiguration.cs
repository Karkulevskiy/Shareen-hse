
using Shareen.Domain;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using System.Security.Cryptography;

class LobbyConfiguration : IEntityTypeConfiguration<Lobby>
{
    public void Configure(EntityTypeBuilder<Lobby> builder)
    {
        builder.HasIndex(id => id.Id).IsUnique();
        builder.HasKey(id => id.Id);
        builder.HasOne(c => c.Chat).WithOne(l => l.Lobby).HasForeignKey<Chat>(c => c.LobbyId);
        builder.Property(n => n.Name).IsRequired().HasMaxLength(30);
        builder.Property(t => t.TimeCreated).IsRequired();
    }
}