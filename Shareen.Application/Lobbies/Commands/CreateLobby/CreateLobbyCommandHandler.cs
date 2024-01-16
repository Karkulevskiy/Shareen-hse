using MediatR;
using Shareen.Application.Interfaces;
using Shareen.Application.Users.Commands.CreateUser;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Commands.CreateLobby;

public class CreateLobbyCommandHandler(IAppDbContext _dbContext)
    : IRequestHandler<CreateLobbyCommand, Guid>
{
    public async Task<Guid> Handle(CreateLobbyCommand request,
        CancellationToken cancellationToken)
    {
        var lobby = new Lobby()
        {
            Id = Guid.NewGuid(),
            TimeCreated = DateTime.Now,
            NumberOfUsers = 0,
            Users = new()
        };
        await _dbContext.Lobbies.AddAsync(lobby, cancellationToken);
        await _dbContext.SaveChangesAsync(cancellationToken);
        return lobby.Id;
    }
}