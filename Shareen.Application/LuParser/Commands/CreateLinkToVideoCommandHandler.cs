using System.Reflection.Metadata;
using AngleSharp;
using AngleSharp.Html.Parser;
using MediatR;
using Shareen.Domain;
namespace Shareen.Application.LuParser.Commands;
public class CreateLinkToVideoCommandHandler 
    : IRequestHandler<CreateLinkToVideoCommand, string>

{
    private static HttpClient _httpClient;
    private readonly static IConfiguration _config = Configuration.Default.WithDefaultLoader();
    private readonly static IBrowsingContext _context = BrowsingContext.New(_config); 
    private readonly static List<string> _attributes = ["src","data-src"];
    static CreateLinkToVideoCommandHandler()
        => _httpClient = new HttpClient(new SocketsHttpHandler
            {
                PooledConnectionLifetime = TimeSpan.FromMinutes(2)
            });
    public async Task<string> Handle(CreateLinkToVideoCommand request,
        CancellationToken cancellationToken)
    {
        if (!UrlIsValid(request.Url))
            return string.Empty;

        var site = GetSite(request.Url);
        return site?.GetIframe!;
    }

    private Site? GetSite(string url)
    {
        if (url.Contains("youtube"))
            return new YouTube(url);
        if (url.Contains("rutube"))
            return new RuTube(url);
        if (url.Contains("vk.com"))
            return new VkVideo(url);
        if (url.Contains("twitch"))
            return new Twitch(url);
        var _url = TryParseSite(url).Result;
        if (_url is not null)
            return new LordFilm(_url);
        return null;
    }

    private async Task<string?> TryParseSite(string url)
    {
        var validSites = new List<string>();
        var document = await _context.OpenAsync(url);
        foreach(var el in document.QuerySelectorAll("iframe"))
        {
            foreach(var atr in _attributes)
            {
                var res = el.GetAttribute(atr);
                if (res is null || !UrlIsValid(res)) continue;
                validSites.Add(res);
            }
        }
        
        if (validSites.Count == 0)
            return null;

        return validSites[0];
    }

    private bool UrlIsValid(string url)
    {
        try
        {
            var response = _httpClient
                .SendAsync(new HttpRequestMessage(HttpMethod.Get, url));
            return response.Result.IsSuccessStatusCode;
        }
        catch
        {
            Console.WriteLine("Link is not valid or resource is closed!");
            return false;
        }
    }
}
