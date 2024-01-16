namespace Shareen.Persistence;

public static class DbInitializer
{
    public static async void Initialize(AppDbContext dbContext,
        CancellationToken cancellationToken)
        => await dbContext.Database.EnsureCreatedAsync(cancellationToken);
}