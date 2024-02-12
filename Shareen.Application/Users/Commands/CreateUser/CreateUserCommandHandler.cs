using MediatR;
using Shareen.Application.Interfaces;
using Shareen.Domain;

namespace Shareen.Application.Users.Commands.CreateUser;

public class CreateUserCommandHandler(IAppDbContext _dbContext)
    : IRequestHandler<CreateUserCommand, Guid>
{
    public async Task<Guid> Handle(CreateUserCommand request,
        CancellationToken cancellationToken)
    {
        var user = new User()
        {
            Name = request.Name,
            Id = Guid.NewGuid(),
            Lobbies = new()
        };
        
        await _dbContext.Users.AddAsync(user, cancellationToken);
        await _dbContext.SaveChangesAsync(cancellationToken);

        return user.Id;
    }
}