using System.Text;
using AutoMapper;
using MediatR;
using Shareen.Application.Interfaces;
using Shareen.Application.Users.Commands.CreateUser;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Commands.CreateLobby;

public class CreateLobbyCommandHandler(IAppDbContext _dbContext, IMapper _mapper)
    : IRequestHandler<CreateLobbyCommand, CreateLobbyDto>
{
    private static string domain = "http://127.0.0.1:5501/";
    public async Task<CreateLobbyDto> Handle(CreateLobbyCommand request,
        CancellationToken cancellationToken)
    {
        var lobbyId = Guid.NewGuid();
        var lobby = new Lobby()
        {
            Id = lobbyId,
            Name = request.Name,
            TimeCreated = DateTime.Now,
            Users = new(),
            UniqueLink = domain + lobbyId.ToString()
        };

        var chat = new Chat
        {
            Id = Guid.NewGuid(),
            LobbyId = lobby.Id,
            Lobby = lobby,
            ListMessages = new()
        };
        
        lobby.Chat = chat;
        lobby.ChatId = chat.Id;

        await _dbContext.Lobbies.AddAsync(lobby, cancellationToken);
        await _dbContext.Chats.AddAsync(chat, cancellationToken);
        await _dbContext.SaveChangesAsync(cancellationToken);

        return _mapper.Map<CreateLobbyDto>(lobby);
    }
}