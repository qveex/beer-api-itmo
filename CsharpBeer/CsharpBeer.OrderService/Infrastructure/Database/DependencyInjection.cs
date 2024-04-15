using CsharpBeer.OrderService.Domain.Common.Interfaces;
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

    private static IServiceCollection AddPersistence(this IServiceCollection services, IConfiguration configuration)
    {
        var connectionString = configuration.GetConnectionString(Constants.POSTGRES_CONNECTION);
        services.AddDbContext<OrderDbContext>(options => options.UseNpgsql(connectionString));
        services.AddTransient<IOrderRepository, OrderRepository>();
        services.AddTransient<IOrderItemRepository, OrderItemRepository>();
        
        return services;
    }
}