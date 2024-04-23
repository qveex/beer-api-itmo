namespace CsharpBeer.OrderService.Domain.Orders;

public enum OrderStatus
{
    None = 0,
    Created,
    Packaging,
    Delivering,
    Delivered,
    Canceled,
    Done
}