using Shareen.Persistence;

static public class DbInitializer
{
    static public async void Initialize(AppDbContext context)
        => await context.Database.EnsureCreatedAsync();
}