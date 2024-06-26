﻿using System.Reflection;
using CsharpBeer.OrderService.Domain.Common.Interfaces;
using CsharpBeer.OrderService.Domain.Orders;
using Microsoft.EntityFrameworkCore;

namespace CsharpBeer.OrderService.Infrastructure.Database;

public class OrderDbContext : DbContext, IUnitOfWork
{
    public OrderDbContext(DbContextOptions options) : base(options)
    {
        Database.Migrate();
    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.ApplyConfigurationsFromAssembly(Assembly.GetExecutingAssembly());
        base.OnModelCreating(modelBuilder);
    }

    public async Task CommitChangesAsync() => await SaveChangesAsync();

    public DbSet<Order> Orders { get; set; } = null!;
    public DbSet<OrderItem> OrderItems { get; set; } = null!;
}