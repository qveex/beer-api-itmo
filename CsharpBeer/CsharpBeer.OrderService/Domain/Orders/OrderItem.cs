namespace CsharpBeer.OrderService.Domain.Orders;

public class OrderItem
{
    public long OrderId { get; set; }
    public long BeerId { get; set; }
    public int Quantity { get; set; }
    public double Price { get; set; }
    
    public virtual Order Order { get; set; }

    public static OrderItem Create(long beerId,
        int quantity,
        double price) =>
        new()
        {
            BeerId = beerId,
            Quantity = quantity,
            Price = price
        };
}