using CsharpBeer.OrderService.Domain.Common.Extensions.Orders;
using CsharpBeer.OrderService.Domain.Common.Interfaces;
using CsharpBeer.OrderService.Domain.Orders;
using Microsoft.EntityFrameworkCore;

namespace CsharpBeer.OrderService.Infrastructure.Database.Orders;

public class OrderRepository(OrderDbContext context) : IOrderRepository
{
    private readonly OrderDbContext _context = context;

    public async Task<Order> GetOrderById(long id)
    {
        var order = await _context.Orders.FirstOrDefaultAsync(o => o.OrderId == id);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");

        return order;
    }

    public Task<List<Order>> ListOrdersByUserId(long userId) =>
        _context.Orders.Where(o => o.UserId == userId).ToListAsync();

    public async Task<Order> CreateOrder(Order order)
    {
        await _context.Orders.AddAsync(order);

        return order;
    }

    public async Task<Order> UpdateOrder(Order order)
    {
        var orderInDb = await _context.Orders.FirstOrDefaultAsync(o => o.OrderId == order.OrderId);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");
        orderInDb.CopyFieldsIfNotNull(order);

        return orderInDb;
    }

    public async Task DeleteOrderById(long id)
    {
        var order = await _context.Orders.FirstOrDefaultAsync(o => o.OrderId == id);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");
        _context.Orders.Remove(order);
    }

    public async Task<Order> GetOrderByIdAsNoTrack(long id)
    {
        var order = await _context.Orders.AsNoTracking().FirstOrDefaultAsync(o => o.OrderId == id);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");

        return order;
    }

    public Task<List<Order>> ListOrdersByUserIdAsNoTrack(long userId) => 
        _context.Orders.AsNoTracking().Where(o => o.UserId == userId).ToListAsync();
}