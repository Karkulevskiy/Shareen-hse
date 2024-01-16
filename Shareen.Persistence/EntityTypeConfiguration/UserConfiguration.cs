using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using Shareen.Domain;

namespace Shareen.Persistence.EntityTypeConfiguration;

public class UserConfiguration : IEntityTypeConfiguration<User>
{
    public void Configure(EntityTypeBuilder<User> builder)
    {
        builder.HasKey(key => key.Id);
        builder.HasIndex(ind => ind.Id).IsUnique();
        builder.Property(prop => prop.Name).HasMaxLength(20);
    }
}