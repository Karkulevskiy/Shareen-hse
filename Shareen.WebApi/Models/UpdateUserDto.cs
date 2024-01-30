using AutoMapper;
using Shareen.Application.Users.Commands.UpdateUser;

public class UpdateUserDto : IMapWith<UpdateUserCommand>
{
    public Guid Id { get; set; }
    public string Name { get; set; }
    void Mapping(Profile profile){
        profile.CreateMap<UpdateUserDto, UpdateUserCommand>()
            .ForMember(userDto => userDto.Name,
                 userCommand => userCommand.MapFrom(prop => prop.Name))
            .ForMember(userDto => userDto.Id,
                 userCommand => userCommand.MapFrom(prop => prop.Id));

    }
}