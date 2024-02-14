namespace Shareen.Domain;

public class Site
{
    private string _iframe;
    public string GetIframe
    {
        get => _iframe;
    }
    public string? SetIframe
    {
        init => _iframe = value;
    }
}