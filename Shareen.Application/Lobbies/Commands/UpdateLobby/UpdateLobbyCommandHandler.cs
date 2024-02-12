using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Application.Users.Commands.UpdateUser;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Commands.UpdateLobby;

public class UpdateLobbyCommandHandler(IAppDbContext _dbContext) 
    : IRequestHandler<UpdateLobbyCommand, Unit>
{
    public async Task<Unit> Handle(UpdateLobbyCommand request,
        CancellationToken cancellationToken)
    {
        var lobby = await _dbContext.Lobbies
            .FirstOrDefaultAsync(lobby =>
                lobby.Id == request.Id, cancellationToken);

        if (lobby == null)
            throw new NotFoundException(request.Id.ToString(), nameof(Lobby));
        
        lobby.Users = request.Users;
        lobby.Name = request.Name;

        await _dbContext.SaveChangesAsync(cancellationToken);
        
        return Unit.Value;
    }
}