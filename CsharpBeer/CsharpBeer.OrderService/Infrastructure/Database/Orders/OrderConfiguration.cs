using CsharpBeer.OrderService.Domain.Orders;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace CsharpBeer.OrderService.Infrastructure.Database.Orders;

public class OrderConfiguration : IEntityTypeConfiguration<Order>
{
    public void Configure(EntityTypeBuilder<Order> builder)
    {
        builder.ToTable("Orders");
        
        builder.HasKey(o => o.OrderId);

        builder.Property(o => o.OrderId)
            .ValueGeneratedOnAdd();
        
        builder.Property(o => o.UserId)
            .IsRequired();

        builder.Property(o => o.Total)
            .IsRequired();

        builder.HasMany(o => o.Items)
            .WithOne(oi => oi.Order)
            .HasForeignKey(oi => oi.OrderId);
    }
}