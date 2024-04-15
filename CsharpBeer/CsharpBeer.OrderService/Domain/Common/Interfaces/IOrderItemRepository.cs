using CsharpBeer.OrderService.Domain.Orders;

namespace CsharpBeer.OrderService.Domain.Common.Interfaces;

public interface IOrderItemRepository
{
    Task<OrderItem> GetOrderItem(long orderId, long beerId);
    Task<List<OrderItem>> GetOrderItemsByOrderId(long orderId);
    Task<OrderItem> CreateOrder(OrderItem order);
    Task<OrderItem> UpdateOrder(OrderItem order);
    Task DeleteOrderItemById(long orderId, long beerId);
}