
using AutoMapper;
using Microsoft.EntityFrameworkCore.Query;
using Shareen.Domain;

public class UserInLobbyDto : IMapWith<User>
{
    public string Name {get; set;}
    public Guid Id {get;set;}

    public void Mapping(Profile profile)
    {
        profile.CreateMap<User, UserInLobbyDto>()
            .ForMember(dto => dto.Name, u => u.MapFrom(prop => prop.Name))
            .ForMember(dto => dto.Id, u => u.MapFrom(prop => prop.Id));
    }
}