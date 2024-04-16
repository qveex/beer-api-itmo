using System.Globalization;
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

    public static Order ToDomain(this OrderDto dto) => 
        Order.Create(orderId: dto.OrderId, userId: dto.UserId, items: dto.Items.ToDomain());

    public static IEnumerable<Order> ToDomain(this IEnumerable<OrderDto> dtos) =>
        dtos.Select(d => d.ToDomain());

    public static Order ToDomain(this CreateOrderDto dto) => 
        Order.Create(-1, userId: dto.UserId, items: dto.Items.ToDomain());

    public static IEnumerable<Order> ToDomain(this IEnumerable<CreateOrderDto> dtos) =>
        dtos.Select(d => d.ToDomain());

    public static OrderDto ToDto(this Order order)
    {
        var dto = new OrderDto()
        {
            OrderId = order.OrderId,
            UserId = order.UserId,
            Status = order.Status.ToDto(),
            Total = order.Total.ToString(CultureInfo.InvariantCulture)
        };
        dto.Items.AddRange(order.Items.ToDto());
        return dto;
    }

    public static IEnumerable<OrderDto> ToDto(this IEnumerable<Order> dtos) =>
        dtos.Select(d => d.ToDto());
}       