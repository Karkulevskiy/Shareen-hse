using MediatR;

namespace Shareen.Application.LuParser.Commands;
public class CreateLinkToVideoCommandHandler 
    : IRequestHandler<CreateLinkToVideoCommand, string>

{
    private static HttpClient _httpClient;
    static CreateLinkToVideoCommandHandler()
        => _httpClient = new HttpClient(new SocketsHttpHandler
            {
                PooledConnectionLifetime = TimeSpan.FromMinutes(2)
            });
    
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
               var link =  GetHtml(request.link);
               if (link == String.Empty)
               {
                   Console.WriteLine("Any working players not found");
                   //правильно ли отправлять контроллеру невалидную ссылку для проверки??
                   //может надо создать exception или делегировать действие 
                   return string.Empty;
               }
               
               return link;
        }
    }

    private string GetHtml(string url)
    {
        var html = _httpClient.GetStringAsync(url).Result;
        var index = 0;
        List<string> links = [];
        
        while (index < html.Length)
        {
            var findIndex = html.IndexOf("iframe src=", index, StringComparison.Ordinal);
            if (findIndex == -1)
                break;
            findIndex += 12; // skip "iframe src="
            var lastIndex = html.IndexOf("\"", findIndex, StringComparison.Ordinal);
            var str = html.Substring(findIndex, lastIndex - findIndex - 1);
            links.Add(str);
            index = lastIndex + 1;
        }
        
        foreach (var link in links.Where(link => UrlIsValid(link)))
        {
            return link;
        }
        return string.Empty;
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
