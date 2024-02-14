

using System.Text.RegularExpressions;
using Shareen.Domain;

public class VkVideo : Site
{
    private const string _template = "<iframe src=\"https://vk.com/video_ext.php?oid=-\" width=\"853\" height=\"480\" " +
    "allow=\"autoplay; encrypted-media; fullscreen; picture-in-picture;\" frameborder=\"0\" " + 
    "allowfullscreen></iframe>";
    public VkVideo(string url)
    {
        var regex = new Regex(@"\d*[0-9]_\d*[0-9]", RegexOptions.Compiled);
        var matchId = regex.Match(url).Value.Split('_');
        var replacedId = matchId[0] + "&id=" + matchId[1] + "&hd=2";
        SetIframe = _template.Insert(_template.IndexOf("=-") + 2, replacedId);
    }
}