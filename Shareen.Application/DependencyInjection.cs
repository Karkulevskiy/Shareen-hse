using System.Reflection;
using AutoMapper;
using Microsoft.Extensions.DependencyInjection;
using MediatR;
namespace Shareen.Application;

/// <summary>
/// 
/// </summary>
public static class DependencyInjection
{
    public static IServiceCollection AddApplication(
        this IServiceCollection serviceCollection)
    {
        serviceCollection.AddMediatR(cfg 
        => cfg.RegisterServicesFromAssemblies(Assembly.GetExecutingAssembly()));
        return serviceCollection;
    }
}