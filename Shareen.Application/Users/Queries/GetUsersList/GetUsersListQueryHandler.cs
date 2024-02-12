using AutoMapper;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Application.Users.Queries.GetUser;
using Shareen.Domain;

namespace Shareen.Application.Users.Queries.GetUsersList;

public class GetUsersListQueryHandler(IAppDbContext _dbContext, IMapper _mapper) 
    : IRequestHandler<GetUsersListQuery, UsersListVm>
{
    public async Task<UsersListVm> Handle(GetUsersListQuery request,
        CancellationToken cancellationToken)
    {
        var users = _dbContext.Users;

        if (users == null)
            throw new NotFoundException(request.LobbyId.ToString(),
                nameof(Lobby));
        
        return new UsersListVm
        {
            Users = await _mapper
                .ProjectTo<UserDto>(users)
                .ToListAsync(cancellationToken)
        };
    }
}