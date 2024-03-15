using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using Shareen.Application.Users.Commands.CreateUser;
using Shareen.Application.Users.Commands.DeleteUser;
using Shareen.Application.Users.Commands.UpdateUser;
using Shareen.Application.Users.Queries.GetUser;

[Route("api/[controller]/[action]")]
public class UserController(IMapper mapper) : BaseController
{
    private readonly IMapper _mapper = mapper;

    [HttpPost]
    public async Task<IActionResult> CreateUser([FromBody] CreateUserDto createUserDto)
    {
        var command = _mapper.Map<CreateUserCommand>(createUserDto);
        var userId = await Mediator.Send(command);
        return Ok(userId);
    }

    [HttpDelete]
    public async Task<IActionResult> DeleteUser(Guid id)
    {
        var command = new DeleteUserCommand{Id = id};
        await Mediator.Send(command);
        return NoContent();
    }

    [HttpPut]
    public async Task<IActionResult> UpdateUser([FromBody] UpdateUserDto updateUserDto)
    {
        var command = _mapper.Map<UpdateUserCommand>(updateUserDto);
        await Mediator.Send(command);
        return NoContent();
    }
    
    [HttpGet]
    public async Task<IActionResult> GetUser(Guid id)
    {
        var command = new GetUserQuery{Id = id};
        var user = await Mediator.Send(command);
        return Ok(user);
    }

    [HttpPost]
    public async Task<IActionResult> AddUserToLobby([FromBody] AddUserToLobbyDto userToLobbyDto)
    {
        var command = _mapper.Map<AddUserToLobbyCommand>(userToLobbyDto);
        var lobbyDto = await Mediator.Send(command);
        return Ok(lobbyDto);
    }
    //реализовать потом получение пользователей
}