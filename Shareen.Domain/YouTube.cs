namespace Shareen.Domain;

public class YouTube : Site
{
    private const string _template 
        = "<iframe width=\"560\" " +
        "height=\"315\" src=\"https://www.youtube.com/embed/\" " +
        "title=\"YouTube video player\" frameborder=\"0\" " +
        "allow=\"accelerometer; autoplay; clipboard-write; " +
        "encrypted-media; gyroscope; picture-in-picture; " +
        "\"web-share\" allowfullscreen></iframe>";

    /// <summary>
    /// Insert YouTube video id after property embed
    /// </summary>
    /// <param name="url"></param>

    public YouTube(string url)
    {
        var _url = _template.Insert(_template.IndexOf("embed") + 6,
                url.Substring(url.IndexOf('=') + 1));
        SetIframe = _url;
    }
} 