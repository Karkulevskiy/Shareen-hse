using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Domain;

namespace Shareen.Application.Users.Commands.DeleteUser;

public class DeleteUserCommandHandler(IAppDbContext _dbContext)
    : IRequestHandler<DeleteUserCommand, Unit>
{
    public async Task<Unit> Handle(DeleteUserCommand request,
        CancellationToken cancellationToken)
    {
        var user = await _dbContext.Users
            .FirstOrDefaultAsync(user => 
                user.Id == request.Id, cancellationToken);

        if (user == null)
            throw new NotFoundException(request.Id.ToString(),
                nameof(User));
        
        _dbContext.Users.Remove(user);
        await _dbContext.SaveChangesAsync(cancellationToken);
        
        return Unit.Value;
    }
}