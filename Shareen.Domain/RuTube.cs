namespace Shareen.Domain;

public class RuTube : Site
{
    private const string _template = 
    "<iframe width=\"720\" height=\"405\"" +
    "src=\"https://rutube.ru/play/embed/\" frameBorder=\"0\" allow=\"clipboard-write; autoplay\"" + 
    "webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe>";
    /// <summary>
    /// Insert RuTube video id after property embed
    /// </summary>
    /// <param name="url"></param>
    public RuTube(string url)
    {
        var _url = _template.Insert(_template.IndexOf("embed") + 6,
                url.Substring(url.IndexOf("video") + 6,
                                url.Length - (url.IndexOf("video") + 6) - 1));
        SetIframe = _url;
    }
}
