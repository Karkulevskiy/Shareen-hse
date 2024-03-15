using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using Shareen.Application.LuParser.Commands;

[ApiController]
public class VideoPlayerController(IMapper mapper) : BaseController
{
    private readonly IMapper _mapper = mapper;

    [HttpPost]
    public async Task<IActionResult> CreateLinkToVideo(string url)
    {
        //HttpContext.Response.Headers.Add("Set-Cookie", "SameSite=None; Secure");
        //IHeaderDictionary.Append("Set-Cookie", "SameSite=None; Secure");
        var command = new CreateLinkToVideoCommand{Url = url};
        var res = await Mediator.Send(command);
        Console.WriteLine(res);
        return Ok(res);
    }
}