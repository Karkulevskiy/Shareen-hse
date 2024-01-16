using AutoMapper;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Interfaces;

namespace Shareen.Application.Lobbies.Queries.GetLobbiesList;

public class GetLobbyListQueryHandler(IAppDbContext _dbContext, IMapper _mapper) 
    : IRequestHandler<GetLobbyListQuery, LobbiesListVm>
{
    public async Task<LobbiesListVm> Handle(GetLobbyListQuery request,
        CancellationToken cancellationToken)
    {
        var lobbies = await _dbContext.Lobbies.ToListAsync(cancellationToken);
        return new LobbiesListVm()
        {
            //Lobbies = _mapper.ProjectTo<LobbyDto>(lobbies).ToListAsync(cancellationToken);
        }//таже хуйня
    }
}