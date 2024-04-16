using System.Globalization;
using Common.Protobuf;
using CsharpBeer.OrderService.Domain.Orders;

namespace CsharpBeer.OrderService.Domain.Common.Extensions.Orders;

public static class OrderItemExtensions
{
    public static OrderItem CopyFieldsIfNotNull(this OrderItem original, OrderItem other)
    {
        original.Price = other.Price == default ? original.Price : other.Price;
        original.Quantity = other.Quantity == default ? original.Quantity : other.Quantity;

        return original;
    }
    public static OrderItemDto ToDto(this OrderItem item) =>
        new()
        {
            OrderId = item.OrderId,
            BeerId = item.BeerId,
            Price = item.Price.ToString(CultureInfo.InvariantCulture),
            Quantity = item.Quantity
        };

    public static IEnumerable<OrderItemDto> ToDto(this IEnumerable<OrderItem> items) =>
        items.Select(oi => oi.ToDto());

    public static OrderItem ToDomain(this OrderItemDto dto) =>
        OrderItem.Create(
            orderId: dto.OrderId,
            beerId: dto.BeerId,
            quantity: dto.Quantity,
            price: decimal.Parse(dto.Price)
        );  
    
    public static IEnumerable<OrderItem> ToDomain(this IEnumerable<OrderItemDto> items) =>
        items.Select(oi => oi.ToDomain());

    public static OrderItem ToDomain(this CreateOrderItemDto dto) =>
        OrderItem.Create(
            orderId: -1,
            beerId: dto.BeerId,
            quantity: dto.Quantity,
            price: decimal.Parse(dto.Price)
        );  
    
    public static IEnumerable<OrderItem> ToDomain(this IEnumerable<CreateOrderItemDto> items) =>
        items.Select(oi => oi.ToDomain());
}