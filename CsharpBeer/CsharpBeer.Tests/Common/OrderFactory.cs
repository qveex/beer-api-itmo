using CsharpBeer.OrderService.Domain.Orders;

namespace CsharpBeer.Tests.Common;

public static class OrderFactory
{
    public static Order CreateOrder(
        long? userId = null,
        IEnumerable<OrderItem>? orderItems = null) => 
        Order.Create(
            userId ?? 0, 
            orderItems ?? Enumerable.Empty<OrderItem>());
}