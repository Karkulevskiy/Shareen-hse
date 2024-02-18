

using Shareen.Domain;

public class Twitch : Site
{
    private const string _template = "<div id=\"twitch-embed\"></div>" +
    "<script src=\"https://player.twitch.tv/js/embed/v1.js\"></script>" +
    "<script type=\"text/javascript\">" + 
    "new Twitch.Player(\"twitch-embed\", {channel: \"\"});</script>";
    public Twitch(string url)
    {
        var _url = _template.Insert(_template.IndexOf("\"\"") + 1,
                url.Substring(url.LastIndexOf('/') + 1));
        SetIframe = _url;
    }
}
