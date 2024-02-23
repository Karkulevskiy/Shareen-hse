using AutoMapper;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Domain;

public class GetLobbyUniqueLinkQueryHandler(IAppDbContext _dbContext) 
    : IRequestHandler<GetLobbyUniqueLinkQuery, string>
{
    public async Task<string> Handle(GetLobbyUniqueLinkQuery request,
        CancellationToken cancellationToken)
    {
        var lobby = await _dbContext
            .Lobbies
            .FirstOrDefaultAsync(lobby => lobby.Id == request.LobbyId);
        
        if (lobby is null)
            throw new NotFoundException(request.LobbyId.ToString(), nameof(Lobby));
        
        return lobby.UniqueLink;
    }
}