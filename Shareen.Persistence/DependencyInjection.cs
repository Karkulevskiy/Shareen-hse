using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Configuration;

namespace Shareen.Persistence;

public static class DependencyInjection
{
    public static IServiceCollection AddPersistence(this IServiceCollection serviceCollection,
        IConfiguration configuration)
    {
        var connectionString = configuration["DbConnectionString"];
        //serviceCollection.AddDbContext<AppDbContext>(opt => opt.Use);
        // to-do : Какую бд использовать?
        return serviceCollection;
    }
}