using Shareen.Domain;   
using MediatR;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Microsoft.EntityFrameworkCore;

class AddUserToLobbyCommandHandler(IAppDbContext _dbContext) : IRequestHandler<AddUserToLobbyCommand, Unit>
{
    public async Task<Unit> Handle(AddUserToLobbyCommand request,
        CancellationToken cancellationToken)
    {
        var user = await _dbContext.Users
            .FirstOrDefaultAsync(u => u.Id == request.UserId);
        
        if (user is null)
            throw new NotFoundException(request.UserId.ToString(), nameof(User));

        var lobby = await _dbContext.Lobbies
            .FirstOrDefaultAsync(l => l.UniqueLink == request.LobbyLink);

        if (lobby is null)
            throw new NotFoundException(request.LobbyLink, nameof(Lobby));

        lobby.Users = [user];
        user.Lobbies = [lobby];

        await _dbContext.SaveChangesAsync(cancellationToken);

        return Unit.Value;
    }
}