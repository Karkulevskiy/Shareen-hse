using System.Reflection;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Application.Users.Commands.DeleteUser;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Commands.DeleteLobby;

public class DeleteLobbyCommandHandler(IAppDbContext _dbContext) :
    IRequestHandler<DeleteUserCommand, Unit>
{
    public async Task<Unit> Handle(DeleteUserCommand request,
        CancellationToken cancellationToken)
    {
        var lobby = await _dbContext.Lobbies
            .FirstOrDefaultAsync(lobby => 
                lobby.Id == request.Id, cancellationToken);

        if (lobby == null)
            throw new NotFoundException(request.Id.ToString(), nameof(Lobby));
        
        _dbContext.Lobbies.Remove(lobby);
        await _dbContext.SaveChangesAsync(cancellationToken);
        
        return Unit.Value;
    }
}