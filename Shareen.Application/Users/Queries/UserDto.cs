using AutoMapper;
using Shareen.Domain;
namespace Shareen.Application.Users.Queries;

public class UserDto : IMapWith<User>
{
    public string Name { get; set; }
    public void Mapping(Profile profile){
        profile.CreateMap<User, UserDto>()
            .ForMember(userDto => userDto.Name,
                 user => user.MapFrom(prop => prop.Name));

    }
}