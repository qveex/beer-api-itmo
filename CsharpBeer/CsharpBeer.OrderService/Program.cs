using CsharpBeer.OrderService.Infrastructure;
using CsharpBeer.OrderService.Services;
using CsharpBeer.OrderService;
using Api;
using CsharpBeer.OrderService.Infrastructure.Database;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
{
    builder.Services.AddGrpc();
    
    var catalogAddress = builder.Configuration.GetSection(Constants.GRPC_SECTION)[Constants.CATALOG_ADDRESS];
    builder.Services.AddGrpcClient<Api.Catalog.CatalogClient>(f => f.Address = new Uri(catalogAddress ?? ""));
    
    builder.Services.AddInfrastructure(builder.Configuration);
}

var app = builder.Build();

// Configure the HTTP request pipeline.
{
    app.MapGrpcService<OrderService>();
}

app.Run();