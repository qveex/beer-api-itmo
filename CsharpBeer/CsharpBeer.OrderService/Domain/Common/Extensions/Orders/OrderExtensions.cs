using System.Globalization;
using Api;
using Common.Protobuf;
using CsharpBeer.OrderService.Domain.Orders;
using Grpc.Core;

namespace CsharpBeer.OrderService.Domain.Common.Extensions.Orders;

public static class OrderExtensions
{
    public static Order CopyFieldsIfNotNull(this Order original, Order other)
    {
        //TODO add items conversion
        //original.Items
        original.Total = other.Total == default ? original.Total : other.Total;
        original.Status = other.Status == OrderStatus.None ? original.Status : other.Status;

        return original;
    }

    public static Order ToDomain(this OrderDto dto)
    {
        var order = new Order { OrderId = dto.OrderId, UserId = dto.UserId };
        order.AddOrderItems(dto.Items.ToDomain(order.OrderId));
        return order;
    }

    public static IEnumerable<Order> ToDomain(this IEnumerable<OrderDto> dtos) =>
        dtos.Select(d => d.ToDomain());

    public static Order ToDomain(this CreateOrderDto dto, long userId, IEnumerable<Beer> beers) => 
        Order.Create(userId: userId, items: dto.Items.ToDomain(beers));

    public static OrderDto ToDto(this Order order)
    {
        var dto = new OrderDto()
        {
            OrderId = order.OrderId,
            UserId = order.UserId,
            Status = order.Status.ToDto(),
            Total = order.Total
        };
        dto.Items.AddRange(order.Items.ToDto());
        return dto;
    }

    public static IEnumerable<OrderDto> ToDto(this IEnumerable<Order> dtos) =>
        dtos.Select(d => d.ToDto());
}       