using System.Reflection;
using System.Security.Cryptography.X509Certificates;
using AutoMapper;
/// <summary>
/// Apply all mappings implementing interface IMapWith
/// </summary>
public class ApplyMappingProfile : Profile
{
    public ApplyMappingProfile(Assembly assembly) 
        => ApplyMappingsFromAssembly(assembly);

    public void ApplyMappingsFromAssembly(Assembly assembly)
    {
        var types = assembly.GetExportedTypes()
            .Where(type => type.GetInterfaces()
            .Any(i => i.IsGenericType 
                    && i.GetGenericTypeDefinition() == typeof(IMapWith<>)))
            .ToList();
        foreach(var type in types)
        {
            var instance = Activator.CreateInstance(type);
            var methodInfo = type.GetMethod("Mapping");
            methodInfo?.Invoke(instance, new object[]{this});
        }    
    }

    
}