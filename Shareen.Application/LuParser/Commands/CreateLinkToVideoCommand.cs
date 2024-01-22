using MediatR;

namespace Shareen.Application.LuParser.Commands;

public class CreateLinkToVideoCommand : IRequest<string>
{
    public string link { get; set; }
}