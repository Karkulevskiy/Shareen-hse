using AutoMapper;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Queries;

public class LobbyDto : IMapWith<Lobby>
{
    public string Name { get; set; }
    public DateTime TimeCreated { get; set; }
    public List<User> Users { get; set; }
    public string LobbyLink { get; set; }

    public void Mapping(Profile profile)
    {
        profile.CreateMap<Lobby, LobbyDto>()
            .ForMember(lobbyDto => lobbyDto.Name,
                 lobby => lobby.MapFrom(prop => prop.Name))
            .ForMember(lobbyDto => lobbyDto.TimeCreated,
                 lobby => lobby.MapFrom(prop => prop.TimeCreated))
            .ForMember(lobbyDto => lobbyDto.Users,
                 lobby => lobby.MapFrom(prop => prop.Users))
            .ForMember(lobbyDto => lobbyDto.LobbyLink,
                 lobby => lobby.MapFrom(prop => prop.UniqueLink));        
    }
}