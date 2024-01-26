using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Configuration;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Interfaces;

namespace Shareen.Persistence;

public static class DependencyInjection
{
    public static IServiceCollection AddPersistence(this IServiceCollection serviceCollection,
        IConfiguration configuration)
    {
        var connectionString = configuration["DbConnectionString"];
        serviceCollection.AddDbContext<AppDbContext>(opt =>
             opt.UseSqlite(connectionString));
        serviceCollection.AddScoped<IAppDbContext>( opt =>
             opt.GetService<AppDbContext>()!);
        return serviceCollection;
    }
}