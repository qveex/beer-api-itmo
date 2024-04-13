namespace CsharpBeer.OrderService.Domain.Orders;

public class OrderItem
{
    public long OrderId { get; set; }
    public long BeerId { get; set; }
    public int Quantity { get; set; }
    public decimal Price { get; set; }
    
    public virtual Order Order { get; set; }
}