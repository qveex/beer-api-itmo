using CsharpBeer.OrderService.Domain.Orders;

namespace CsharpBeer.OrderService.Domain.Common.Interfaces;

public interface IOrderRepository
{
    Task<Order> GetOrderById(long id);
    Task<List<Order>> ListOrdersByUserId(long userId);
    Task<Order> CreateOrder(Order order);
    Task<Order> UpdateOrder(Order order);
    Task DeleteOrderById(long id);
}