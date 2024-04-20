using CsharpBeer.OrderService.Domain.Orders;

namespace CsharpBeer.Tests.Common;

public static class OrderItemFactory
{
    public static OrderItem CreateOrderItem(
        long? beerId = null,
        int? quantity = null,
        double? price = null) =>
        OrderItem.Create(
            beerId ?? 0,
            quantity ?? 1,
            price ?? 20d);
}