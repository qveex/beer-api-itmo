using CsharpBeer.OrderService.Domain.Common.Interfaces;
using CsharpBeer.OrderService.Infrastructure.Auth;
using CsharpBeer.OrderService.Infrastructure.Catalog;
using CsharpBeer.OrderService.Infrastructure.Database.OrderItems;
using CsharpBeer.OrderService.Infrastructure.Database.Orders;
using Microsoft.EntityFrameworkCore;

namespace CsharpBeer.OrderService.Infrastructure.Database;

public static class DependencyInjection
{
    public static IServiceCollection AddInfrastructure(this IServiceCollection services, IConfiguration configuration)
    {
        return services.AddPersistence(configuration);
    }

    // public static IApplicationBuilder AddInfrastructureMiddleware(this IApplicationBuilder builder)
    // {
    //     builder.UseMiddleware<ConsistencyMiddleware>();

    //     return builder;
    // }

    private static IServiceCollection AddPersistence(this IServiceCollection services, IConfiguration configuration)
    {
        var connectionString = configuration.GetConnectionString(Constants.POSTGRES_CONNECTION);
        services.AddDbContext<OrderDbContext>(options => options.UseNpgsql(connectionString));
        services.AddScoped<IOrderRepository, OrderRepository>();
        services.AddScoped<IOrderItemRepository, OrderItemRepository>();
        services.AddScoped<IAuthService, AuthService>();
        services.AddScoped<ICatalogService, CatalogService>();
        services.AddScoped<IUnitOfWork>(serviceProvider => serviceProvider.GetRequiredService<OrderDbContext>());
        
        return services;
    }
}