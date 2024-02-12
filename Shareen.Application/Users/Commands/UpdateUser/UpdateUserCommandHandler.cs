using MediatR;
using Microsoft.EntityFrameworkCore;
using Shareen.Application.Exceptions;
using Shareen.Application.Interfaces;
using Shareen.Domain;

namespace Shareen.Application.Users.Commands.UpdateUser;

public class UpdateUserCommandHandler(IAppDbContext _dbContext) :
    IRequestHandler<UpdateUserCommand, Unit>
{
    public async Task<Unit> Handle(UpdateUserCommand request,
        CancellationToken cancellationToken)
    {
        var user = await _dbContext.Users
            .FirstOrDefaultAsync(user => 
                user.Id == request.Id, cancellationToken);

        if (user == null)
            throw new NotFoundException(request.Id.ToString(), nameof(User));
        
        user.Name = request.Name;
        
        await _dbContext.SaveChangesAsync(cancellationToken);
        return Unit.Value;
    }
}