using Shareen.Domain;

public class LordFilm : Site
{
    private const string _template = "<iframe width=\"560\" height=\"400\" src=\"\" loading=\"lazy\" frameborder=\"0\" allowfullscreen=\"\"></iframe>";
    public LordFilm(string url)
    {
        SetIframe = _template.Insert(_template.IndexOf("\"\"") + 1, url);
    }
}