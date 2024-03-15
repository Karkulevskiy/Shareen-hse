using Shareen.Domain;   
using MediatR;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Microsoft.EntityFrameworkCore;
using AutoMapper;
using System.Collections.Immutable;
using Shareen.Application.Lobbies.Queries;
using Shareen.Application.Lobbies.Queries.GetLobbiesList;

class AddUserToLobbyCommandHandler(IAppDbContext _dbContext, IMapper _mapper) 
    : IRequestHandler<AddUserToLobbyCommand, LobbyDto>
{
    public async Task<LobbyDto> Handle(AddUserToLobbyCommand request,
        CancellationToken cancellationToken)
    {
        //Добавить аутенитифакацию для пользователей
        var user = await _dbContext.Users
            .FirstOrDefaultAsync(u => u.Name == request.UserName);
        if (user is null)
        {
            user = new User{
                Id = Guid.NewGuid(),
                Name = request.UserName,
                Lobbies = []
            };
            await _dbContext.Users.AddAsync(user);

        }
            /* throw new NotFoundException(request.UserId.ToString(), nameof(User)); */
        var lobby = await _dbContext.Lobbies
            .FirstOrDefaultAsync(l => l.UniqueLink == request.LobbyLink, cancellationToken)
            ?? throw new NotFoundException(request.LobbyLink, nameof(Lobby));

        //await _dbContext.SaveChangesAsync(cancellationToken);
        user.Lobbies.Add(lobby);
        lobby.Users.Add(user);
        await _dbContext.SaveChangesAsync(cancellationToken);
        var lobbyUsers = _dbContext.Users.Where( u => u.Lobbies.Contains(lobby));
        lobby.Users = await lobbyUsers.ToListAsync(cancellationToken);
        return _mapper.Map<LobbyDto>(lobby);
    }
}