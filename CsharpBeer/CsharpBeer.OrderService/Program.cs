using CsharpBeer.OrderService.Services;
using CsharpBeer.OrderService;
using CsharpBeer.OrderService.Infrastructure.Database;
using Api.Catalog;
using Api.Auth;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
{
    builder.Services.AddGrpc();
    
    var catalogAddress = builder.Configuration.GetSection(Constants.GRPC_SECTION)[Constants.CATALOG_ADDRESS];
    var identityAddress = builder.Configuration.GetSection(Constants.GRPC_SECTION)[Constants.IDENTITY_ADDRESS];
    builder.Services.AddGrpcClient<Catalog.CatalogClient>(f => f.Address = new Uri(catalogAddress ?? ""));
    builder.Services.AddGrpcClient<Auth.AuthClient>(f => f.Address = new Uri(identityAddress ?? ""));
    
    builder.Services.AddInfrastructure(builder.Configuration);
}

var app = builder.Build();

// Configure the HTTP request pipeline.
{
    app.MapGrpcService<OrderService>();
}

app.Run();