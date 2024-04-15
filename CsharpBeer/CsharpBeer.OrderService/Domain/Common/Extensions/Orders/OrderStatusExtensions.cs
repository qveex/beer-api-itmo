using Common.Protobuf;
using CsharpBeer.OrderService.Domain.Orders;

namespace CsharpBeer.OrderService.Domain.Common.Extensions.Orders;

public static class OrderStatusExtensions
{
    public static StatusOrderDto ToDto(this OrderStatus status) => status switch
    {
        OrderStatus.None => StatusOrderDto.None,
        OrderStatus.CREATED => StatusOrderDto.Created,
        OrderStatus.PACKAGING => StatusOrderDto.Packaging,
        OrderStatus.DELIVERING => StatusOrderDto.Delivering,
        OrderStatus.DELIVERED => StatusOrderDto.Delivered,
        OrderStatus.CANCELED => StatusOrderDto.Canceled,
        OrderStatus.DONE => StatusOrderDto.Done,
        _ => StatusOrderDto.None
    };
    
    public static OrderStatus ToDomain(this StatusOrderDto status) => status switch
    {
        StatusOrderDto.None => OrderStatus.None,
        StatusOrderDto.Created => OrderStatus.CREATED,
        StatusOrderDto.Packaging => OrderStatus.PACKAGING,
        StatusOrderDto.Delivering => OrderStatus.DELIVERING,
        StatusOrderDto.Delivered => OrderStatus.DELIVERED,
        StatusOrderDto.Canceled => OrderStatus.CANCELED,
        StatusOrderDto.Done => OrderStatus.DONE,
        _ => OrderStatus.None
    };
}