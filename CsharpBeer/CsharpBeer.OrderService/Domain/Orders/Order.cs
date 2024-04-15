namespace CsharpBeer.OrderService.Domain.Orders;

public class Order
{
    private List<OrderItem> _items = [];
    public long OrderId { get; init; }
    public long UserId { get; init; }
    public decimal Total { get; set; }
    public OrderStatus Status { get; set; }
    public IReadOnlyCollection<OrderItem> Items => _items;

    public static Order Create(long orderId, long userId, IEnumerable<OrderItem> items)
    {
        var orderItems = items.ToList();
        return new Order()
        {
            OrderId = orderId,
            UserId = userId,
            Total = orderItems.Sum(oi => oi.Price * oi.Quantity),
            Status = OrderStatus.CREATED,
            _items = orderItems
        };
    }
}