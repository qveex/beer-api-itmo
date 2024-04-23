using System.Globalization;
using Api;
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

    public static OrderItem ToDomain(this Beer beer, long orderId = -1,int quantity = 1) => 
        new()
        {
            OrderId = orderId,
            BeerId = beer.BeerId,
            Price = beer.Price,
            Quantity = quantity,
        };
    
    public static OrderItemDto ToDto(this OrderItem item) =>
        new()
        {
            OrderId = item.OrderId,
            BeerId = item.BeerId,
            Price = item.Price,
            Quantity = item.Quantity
        };

    public static IEnumerable<OrderItemDto> ToDto(this IEnumerable<OrderItem> items) =>
        items.Select(oi => oi.ToDto());

    public static OrderItem ToDomain(this OrderItemDto dto, long orderId)
    {
        return new()
        {
            BeerId = dto.BeerId,
            OrderId = orderId,
            Price = dto.Price,
            Quantity = dto.Quantity
        };
    }  
    
    public static IEnumerable<OrderItem> ToDomain(this IEnumerable<OrderItemDto> items, long orderId) =>
        items.Select(oi => oi.ToDomain(orderId));

    public static OrderItem ToDomain(this CreateOrderItemDto dto, Beer beer) =>
        OrderItem.Create(
            beerId: dto.BeerId,
            quantity: dto.Quantity,
            price: beer.Price
        );

    public static IEnumerable<OrderItem> ToDomain(this IEnumerable<CreateOrderItemDto> dtos, IEnumerable<Beer> beers) =>
        dtos.Zip(beers).Select(t => t.First.ToDomain(t.Second));
}