using MediatR;

namespace Shareen.Application.LuParser.Commands;
public class CreateLinkToVideoCommandHandler 
    : IRequestHandler<CreateLinkToVideoCommand, string>
        
{
    public async Task<string> Handle(CreateLinkToVideoCommand request,
        CancellationToken cancellationToken)
    {
        var parsedLinkList = request.link.Split('/');
        var parsedLinkSite = parsedLinkList[2];
        var resultLink = string.Empty;
        const string https = "https://";
        const string embed = "/embed/";
        const string video = "video";
        const string youtubeTemplate = "";
        switch (parsedLinkSite)
        {
            case "www.youtube.com":
                var lastSplash = request.link.LastIndexOf('/');
                resultLink = https + parsedLinkSite
                                   + embed
                                   + request.link[lastSplash..];
                return "<iframe width=\"560\" height=\"315\" src=" + 
                       $"\"{resultLink}\"" +
                "title=\"YouTube video player\"" +
                "frameborder=\"0\"" +
                "allow=\"accelerometer; autoplay; clipboard-write;" +
                       " encrypted-media; gyroscope; picture-in-picture; web-share\"" +
                "allowfullscreen></iframe>";

            case "rutube.ru":
                resultLink = https + video + embed + parsedLinkList[^2] + '/';
                return "<iframe width=\"720\" height=\"405\" src=" +
                    $"\"{resultLink}\"" + 
                    "frameBorder=\"0\"" +
                    "allow=\"clipboard-write; autoplay\"" +
                    "webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe>";
            //case lordfilms, нужно сделать запрос к страничке и получить html, а дальше парсить
            default:
                return string.Empty;
        }
    }
}
