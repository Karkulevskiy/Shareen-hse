using MediatR;

namespace Shareen.Application.LuParser.Commands;

public class CreateLinkToVideoCommand : IRequest<string>
{
    public string Url { get; set; }
}