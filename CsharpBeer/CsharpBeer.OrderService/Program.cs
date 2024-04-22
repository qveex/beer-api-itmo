using CsharpBeer.OrderService.Services;
using CsharpBeer.OrderService;
using CsharpBeer.OrderService.Infrastructure.Database;
using AuthC = Auth.Auth;
using Api;

var builder = WebApplication.CreateBuilder(args);
builder.Configuration.AddEnvironmentVariables();

// Add services to the container.
{
    builder.Logging.ClearProviders();
    builder.Logging.AddConsole();

    builder.Services.AddGrpc(op => 
    {
        op.EnableDetailedErrors = true;
    });
    
    var catalogAddress = builder.Configuration.GetSection(Constants.GRPC_SECTION)[Constants.CATALOG_ADDRESS];
    var identityAddress = builder.Configuration.GetSection(Constants.GRPC_SECTION)[Constants.IDENTITY_ADDRESS];
    builder.Services.AddGrpcClient<Catalog.CatalogClient>(f => f.Address = new Uri(catalogAddress ?? ""));
    builder.Services.AddGrpcClient<AuthC.AuthClient>(f => f.Address = new Uri(identityAddress ?? ""));
    
    builder.Services.AddInfrastructure(builder.Configuration);
}

var app = builder.Build();

// Configure the HTTP request pipeline.
{
    //app.AddInfrastructureMiddleware();
    app.MapGrpcService<OrderService>();
}

app.Run();