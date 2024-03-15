using Shareen.Persistence;
using Shareen.Application;
using AutoMapper;
using Swashbuckle.AspNetCore.Swagger;
using System.Reflection;
using Shareen.Application.Interfaces;
using System.ComponentModel;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddAutoMapper(cfg =>
{
    cfg.AddProfile(new AssemblyMappingProfile(Assembly.GetExecutingAssembly()));
    cfg.AddProfile(new AssemblyMappingProfile(typeof(IAppDbContext).Assembly));
});
builder.Services.AddPersistence(builder.Configuration);
builder.Services.AddApplication();
builder.Services.AddControllers();
builder.Services.AddSwaggerGen();
builder.Services.AddCors();
builder.Services.AddControllers()
    .AddNewtonsoftJson(opt =>
     opt.SerializerSettings.ReferenceLoopHandling 
        = Newtonsoft.Json.ReferenceLoopHandling.Ignore);

var app = builder.Build();

using (var scope = app.Services.CreateScope())
{
    var serviceProvider = scope.ServiceProvider;
    try
    {
        var dbContext = serviceProvider.GetRequiredService<AppDbContext>();
        DbInitializer.Initialize(dbContext);
    }
    catch(Exception exception)
    {
        Console.WriteLine("Can not initialize dbContext");
    }
}

app.UseCors(cfg => cfg.AllowAnyOrigin()); // потому нужно настроить cors

app.UseSwagger();
app.UseSwaggerUI();
app.UseRouting();
app.UseEndpoints(cfg =>
{
    _ = cfg.MapControllers();
});
app.MapGet("/", () => "Hello World!");

app.Run();