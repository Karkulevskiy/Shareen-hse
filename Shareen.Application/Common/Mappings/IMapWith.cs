using AutoMapper;

/// <summary>
/// Creating configuration from type T to destination type
/// </summary>
/// <typeparam name="T"></typeparam>
public interface IMapWith<T>
{
    void Mapping(Profile profile) 
        => profile.CreateMap(typeof(T), GetType());    
}