using Shareen.Persistence;
using Shareen.Application;
using Swashbuckle.AspNetCore.Swagger;

var builder = WebApplication.CreateBuilder(args);
//Connect 2 libraries with project
builder.Services.AddPersistence(builder.Configuration);
builder.Services.AddApplication();
builder.Services.AddControllers();
builder.Services.AddSwaggerGen();
var app = builder.Build();
app.UseSwagger();
app.UseSwaggerUI();
app.MapGet("/", () => "Hello World!");

app.Run();