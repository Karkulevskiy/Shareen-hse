using Shareen.Persistence;
using Shareen.Application;using AutoMapper;
using Swashbuckle.AspNetCore.Swagger;
using System.Reflection;
using Shareen.Application.Interfaces;

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
var app = builder.Build();
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