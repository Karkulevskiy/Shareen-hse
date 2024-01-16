using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using Shareen.Domain;

namespace Shareen.Persistence.EntityTypeConfiguration;

public class LobbyConfiguration : IEntityTypeConfiguration<Lobby>
{
    public void Configure(EntityTypeBuilder<Lobby> builder)
    {
        builder.HasKey(key => key.Id);
        builder.HasIndex(ind => ind.Id).IsUnique();
        builder.Property(prop => prop.TimeCreated).IsRequired();
    }
}