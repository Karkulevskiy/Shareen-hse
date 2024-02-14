using AutoMapper;
using Shareen.Application;
public class AddUserToLobbyDto : IMapWith<AddUserToLobbyCommand>
{
    public Guid UserId { get; set; }
    public string LobbyLink { get; set; }
    public void Mapping(Profile profile)
    {
        profile.CreateMap<AddUserToLobbyDto, AddUserToLobbyCommand>()
            .ForMember(dto => dto.UserId, com => com.MapFrom(prop => prop.UserId))
            .ForMember(dto => dto.LobbyLink, com => com.MapFrom(prop => prop.LobbyLink));
    }
} 