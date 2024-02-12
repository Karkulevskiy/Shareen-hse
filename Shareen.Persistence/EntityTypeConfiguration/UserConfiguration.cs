
using Shareen.Domain;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

class UserConfiguration : IEntityTypeConfiguration<User>
{
    public void Configure(EntityTypeBuilder<User> builder)
    {
        builder.HasIndex(id => id.Id).IsUnique();
        builder.HasKey(id => id.Id);
        builder.Property(n => n.Name).IsRequired().HasMaxLength(30);
        builder.HasMany(l => l.Lobbies).WithMany(u => u.Users);
    }
}