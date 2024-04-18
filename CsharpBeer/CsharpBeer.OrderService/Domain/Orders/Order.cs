namespace CsharpBeer.OrderService.Domain.Orders;

public class Order
{
    private List<OrderItem> _items = [];
    public long OrderId { get; set; }
    public long UserId { get; set; }
    public double Total { get; set; }
    public OrderStatus Status { get; set; }
    public IReadOnlyCollection<OrderItem> Items => _items;

    public void AddOrderItem(OrderItem orderItem)
    {
        var item = _items.FirstOrDefault(oi => oi.BeerId == orderItem.BeerId);
        if (item is not null) return;

        orderItem.OrderId = OrderId;
        _items.Add(orderItem);
        Total = _items.Sum(oi => oi.Quantity * oi.Price);
    }

    public void AddOrderItems(IEnumerable<OrderItem> items)
    {
        foreach (var orderItem in items) AddOrderItem(orderItem);
    }
    public static Order Create(long userId, IEnumerable<OrderItem> items)
    {
        var orderItems = items.ToList();
        return new Order()
        {
            UserId = userId,
            Total = orderItems.Sum(oi => oi.Price * oi.Quantity),
            Status = OrderStatus.Created,
            _items = orderItems
        };
    }
}