using System.Security.Cryptography.X509Certificates;
using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using Shareen.Application.Lobbies.Commands.CreateLobby;
using Shareen.Application.Lobbies.Commands.DeleteLobby;
using Shareen.Application.Lobbies.Commands.UpdateLobby;
using Shareen.Application.Lobbies.Queries.GetLobbiesList;
using Shareen.Application.Lobbies.Queries.GetLobby;

[Route("api/[controller]/[action]")]
public class LobbyController(IMapper mapper) : BaseController
{
    private readonly IMapper _mapper = mapper;

    [HttpGet]
    public async Task<IActionResult> GetLobbyById(Guid id)
    {
        var query = new GetLobbyQuery{Id = id};
        var lobby = await Mediator.Send(query);
        return Ok(lobby);
    }

    [HttpGet]
    public async Task<IActionResult> GetLobbyList()
    {
        var query = new GetLobbyListQuery();
        var lobbies = await Mediator.Send(query);
        return Ok(lobbies);
    }

    [HttpPost]
    public async Task<IActionResult> CreateLobby(string lobbyName)
    {
        var command = new CreateLobbyCommand{Name = lobbyName};
        var lobbyId = await Mediator.Send(command);
        return Ok(lobbyId);
    }
    
    [HttpPatch]
    public async Task<IActionResult> UpdateLobby([FromBody] LobbyDto lobbyDto)
    {
        var command = _mapper.Map<UpdateLobbyCommand>(lobbyDto);
        await Mediator.Send(command);
        return NoContent();
    }
    
    [HttpDelete]
    public async Task<IActionResult> DeleteLobby(Guid id)
    {
        var command = new DeleteLobbyCommand{Id = id};
        await Mediator.Send(command);
        return NoContent();
    }

    [HttpGet]
    public async Task<IActionResult> GetLobbyUniqueId(Guid id)
    {
        var query = new GetLobbyUniqueLinkQuery{LobbyId = id};
        var uniqueLink = await Mediator.Send(query);
        return Ok(uniqueLink);
    }

    [HttpGet]
    public async Task<IActionResult> GetLobbyUsers(Guid id)
    {
        var query = new GetLobbyUsersQuery{LobbyId = id};
        var lobbyUsers = await Mediator.Send(query);
        return Ok(lobbyUsers);
    }
}