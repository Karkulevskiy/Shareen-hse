

using AutoMapper;
using Shareen.Domain;

public class CreateLobbyDto : IMapWith<Lobby>
{
    public Guid Id { get; set; }
    public string UniqueLink { get; set; }
    public List<UserInLobbyDto> Users { get; set; }
    public Guid ChatId { get; set; }
    public void Mapping(Profile profile)
    {
        profile.CreateMap<Lobby, CreateLobbyDto>()
         .ForMember(l => l.Id,
         dto => dto.MapFrom(prop => prop.Id))
         .ForMember(l => l.UniqueLink,
         dto => dto.MapFrom(prop => prop.UniqueLink))
         .ForMember(l => l.ChatId,
         dto => dto.MapFrom(prop => prop.ChatId));
    }

}