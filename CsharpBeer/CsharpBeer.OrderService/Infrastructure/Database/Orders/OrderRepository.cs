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

    public Task<List<Order>> ListOrdersByUserId(long userId)
    {
        return _context.Orders.Where(o => o.UserId == userId).ToListAsync();
    }

    public async Task<Order> CreateOrder(Order order)
    {
        await _context.Orders.AddAsync(order);
        await _context.SaveChangesAsync();

        return order;
    }

    public async Task<Order> UpdateOrder(Order order)
    {
        var orderInDb = await _context.Orders.FirstOrDefaultAsync(o => o.OrderId == order.OrderId);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");
        orderInDb.CopyFieldsIfNotNull(order);

        await _context.SaveChangesAsync();

        return orderInDb;
    }

    public async Task DeleteOrderById(long id)
    {
        var order = await _context.Orders.FirstOrDefaultAsync(o => o.OrderId == id);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");
        _context.Orders.Remove(order);

        await _context.SaveChangesAsync();
    }
}