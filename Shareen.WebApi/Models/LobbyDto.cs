using AutoMapper;
using Shareen.Application.Lobbies.Commands.UpdateLobby;
using Shareen.Domain;
public class LobbyDto : IMapWith<UpdateLobbyCommand>
{
    public Guid Id { get; set; }
    public string Name { get; set; }
    public List<User> Users { get; set; }

    public void Mapping(Profile profile)
    {
        profile.CreateMap<LobbyDto, UpdateLobbyCommand>()
            .ForMember(prop => prop.Id, p => p.MapFrom(prop => prop.Id))
            .ForMember(prop => prop.Name, p => p.MapFrom(prop => prop.Name))
            .ForMember(prop => prop.Users, p => p.MapFrom(prop => prop.Users));
    }
}