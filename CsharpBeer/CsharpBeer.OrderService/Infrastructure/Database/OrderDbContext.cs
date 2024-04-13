using System.Reflection;
using CsharpBeer.OrderService.Domain.Orders;
using Microsoft.EntityFrameworkCore;

namespace CsharpBeer.OrderService.Infrastructure.Database;

public class OrderDbContext(DbContextOptions options) : DbContext(options)
{
    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.ApplyConfigurationsFromAssembly(Assembly.GetExecutingAssembly());
    }

    public DbSet<Order> Orders { get; set; } = null!;
    public DbSet<OrderItem> OrderItems { get; set; } = null!;
}