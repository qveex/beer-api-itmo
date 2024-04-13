using Microsoft.EntityFrameworkCore;

namespace CsharpBeer.OrderService.Infrastructure.Database;

public static class DependencyInjection
{
    public static IServiceCollection AddInfrastructure(this IServiceCollection services, IConfiguration configuration)
    {
        return services.AddPersistence(configuration);
    }
    public static IServiceCollection AddPersistence(this IServiceCollection services, IConfiguration configuration)
    {
        var connectionString = configuration.GetConnectionString(Constants.POSTGRES_CONNECTION);
        services.AddDbContext<OrderDbContext>(options => options.UseNpgsql(connectionString));

        return services;
    }
}