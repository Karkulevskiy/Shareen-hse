using AutoMapper;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Domain;

namespace Shareen.Application.Users.Queries.GetUser;

public class GetUserQueryHandler(IAppDbContext _dbContext, IMapper _mapper)
    : IRequestHandler<GetUserQuery, UserDto>
{
    public async Task<UserDto> Handle(GetUserQuery request,
        CancellationToken cancellationToken)
    {
        var user = await _dbContext.Users
            .FirstOrDefaultAsync(user =>
                user.Id == request.Id, cancellationToken);
        if (user == null)
            throw new NotFoundException(request.Id.ToString(),nameof(User));
        return _mapper.Map<UserDto>(user);
    }
}