using AutoMapper;
using Shareen.Application.Users.Commands.CreateUser;

public class CreateUserDto : IMapWith<CreateUserCommand>
{
    public string Name { get; set; }
    public void Mapping(Profile profile){
        profile.CreateMap<CreateUserDto, CreateUserCommand>()
            .ForMember(userDto => userDto.Name,
                userCommand => userCommand.MapFrom(prop => prop.Name));
    }   
}