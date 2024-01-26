using System.Reflection;
using AutoMapper;
using Microsoft.Extensions.DependencyInjection;
using MediatR;
namespace Shareen.Application;

/// <summary>
/// Класс для добавления Application в WebApi
/// Добавляем MediatR
/// </summary>
public static class DependencyInjection
{
    public static IServiceCollection AddApplication(
        this IServiceCollection serviceCollection)
    {
        serviceCollection.AddMediatR(Assembly.GetExecutingAssembly());
        return serviceCollection;
    }
}