using Shareen.Domain;

using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

class ChatConfiguration : IEntityTypeConfiguration<Chat>
{
    public void Configure(EntityTypeBuilder<Chat> builder)
    {
        builder.HasIndex(id => id.Id).IsUnique();
        builder.HasKey(id => id.Id);
        builder.HasOne(l => l.Lobby).WithOne(c => c.Chat).IsRequired();
    }
}
