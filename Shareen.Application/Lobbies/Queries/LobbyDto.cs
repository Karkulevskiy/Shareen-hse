using AutoMapper;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Queries;

public class LobbyDto : IMapWith<Lobby>
{
    public string Name { get; set; }
    public DateTime TimeCreated { get; set; }
    public List<User> Users { get; set; }

    void Mapping(Profile profile)
    {
        profile.CreateMap<Lobby, LobbyDto>()
            .ForMember(lobbyDto => lobbyDto.Name,
                lobby => lobby.MapFrom(prop => prop.Name))
            .ForMember(LobbyDto => LobbyDto.TimeCreated,
                 lobby => lobby.MapFrom(prop => prop.TimeCreated))
            .ForMember(LobbyDto => LobbyDto.Users,
                 lobby => lobby.MapFrom(prop => prop.Users));        
    }
}