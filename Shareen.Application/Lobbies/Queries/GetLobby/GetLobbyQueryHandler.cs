using AutoMapper;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Domain;

namespace Shareen.Application.Lobbies.Queries.GetLobby;

public class GetLobbyQueryHandler(IAppDbContext _dbContext, IMapper _mapper) 
    : IRequestHandler<GetLobbyQuery, LobbyDto>
{
    public async Task<LobbyDto> Handle(GetLobbyQuery request,
        CancellationToken cancellationToken)
    {
        var lobby = await _dbContext.Lobbies
            .FirstOrDefaultAsync(lobby =>
                lobby.Id == request.Id, cancellationToken);
        if (lobby == null)
            throw new NotFoundException(request.Id.ToString(), nameof(Lobby));
        return _mapper.Map<LobbyDto>(lobby);
    }
}