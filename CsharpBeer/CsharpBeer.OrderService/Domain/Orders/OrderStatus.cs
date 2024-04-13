namespace CsharpBeer.OrderService.Domain.Orders;

public enum OrderStatus
{
    None = 0,
    CREATED,
    PACKAGING,
    DELIVERING,
    DELIVERED,
    CANCELED,
    DONE
}