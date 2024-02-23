using AutoMapper;
using Shareen.Domain;
public class LobbyUserDto : IMapWith<User>
{
    public Guid UserId { get; set; }
    public string? UserName { get; set; }
    public void Mapping(Profile profile)
    {
        profile.CreateMap<User, LobbyUserDto>()
            .ForMember(userDto => userDto.UserId,
                user => user.MapFrom(prop => prop.Id))
            .ForMember(userDto => userDto.UserName,
                user => user.MapFrom(prop => prop.Name));
    }
}