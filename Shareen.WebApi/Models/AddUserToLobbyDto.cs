using AutoMapper;
using Shareen.Application;
public class AddUserToLobbyDto : IMapWith<AddUserToLobbyCommand>
{
    public string UserName { get; set; }
    public string LobbyLink { get; set; }
    public void Mapping(Profile profile)
    {
        profile.CreateMap<AddUserToLobbyDto, AddUserToLobbyCommand>()
            .ForMember(dto => dto.UserName, com => com.MapFrom(prop => prop.UserName))
            .ForMember(dto => dto.LobbyLink, com => com.MapFrom(prop => prop.LobbyLink));
    }
} 