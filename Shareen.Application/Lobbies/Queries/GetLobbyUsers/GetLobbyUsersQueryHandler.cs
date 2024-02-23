using AutoMapper;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Domain;

public class GetLobbyUsersQueryHandler(IAppDbContext _dbContext, IMapper _mapper) 
: IRequestHandler<GetLobbyUsersQuery, LobbyUsersListDto>
{
    public async Task<LobbyUsersListDto> Handle(GetLobbyUsersQuery request,
        CancellationToken cancellationToken)
    {
        var lobby = await _dbContext
            .Lobbies
            .FirstOrDefaultAsync(l => l.Id == request.LobbyId);
        
        if (lobby is null)
            throw new NotFoundException(request.LobbyId.ToString(), nameof(Lobby));

        return new LobbyUsersListDto
        {
            Users = await _mapper
                .ProjectTo<LobbyUserDto>(lobby.Users.AsQueryable())
                .ToListAsync(cancellationToken)
        };     
    }
}