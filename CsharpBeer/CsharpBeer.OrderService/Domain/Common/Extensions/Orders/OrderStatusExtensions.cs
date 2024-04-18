using Common.Protobuf;
using CsharpBeer.OrderService.Domain.Orders;

namespace CsharpBeer.OrderService.Domain.Common.Extensions.Orders;

public static class OrderStatusExtensions
{
    public static StatusOrderDto ToDto(this OrderStatus status) => status switch
    {
        OrderStatus.None => StatusOrderDto.None,
        OrderStatus.Created => StatusOrderDto.Created,
        OrderStatus.Packaging => StatusOrderDto.Packaging,
        OrderStatus.Delivering => StatusOrderDto.Delivering,
        OrderStatus.Delivered => StatusOrderDto.Delivered,
        OrderStatus.Canceled => StatusOrderDto.Canceled,
        OrderStatus.Done => StatusOrderDto.Done,
        _ => StatusOrderDto.None
    };
    
    public static OrderStatus ToDomain(this StatusOrderDto status) => status switch
    {
        StatusOrderDto.None => OrderStatus.None,
        StatusOrderDto.Created => OrderStatus.Created,
        StatusOrderDto.Packaging => OrderStatus.Packaging,
        StatusOrderDto.Delivering => OrderStatus.Delivering,
        StatusOrderDto.Delivered => OrderStatus.Delivered,
        StatusOrderDto.Canceled => OrderStatus.Canceled,
        StatusOrderDto.Done => OrderStatus.Done,
        _ => OrderStatus.None
    };
}