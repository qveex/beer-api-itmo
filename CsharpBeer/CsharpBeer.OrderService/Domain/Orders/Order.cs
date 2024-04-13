namespace CsharpBeer.OrderService.Domain.Orders;

public class Order
{
    private List<OrderItem> _items = [];
    public long OrderId { get; private set; }
    public long UserId { get; private set; }
    public decimal Total { get; set; }
    public OrderStatus Status { get; set; }
    public IReadOnlyCollection<OrderItem> Items => _items;
}