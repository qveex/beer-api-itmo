﻿using CsharpBeer.OrderService.Domain.Common.Extensions.Orders;
using CsharpBeer.OrderService.Domain.Common.Interfaces;
using CsharpBeer.OrderService.Domain.Orders;
using Microsoft.EntityFrameworkCore;

namespace CsharpBeer.OrderService.Infrastructure.Database.OrderItems;

public class OrderItemRepository(OrderDbContext context) : IOrderItemRepository
{
    private readonly OrderDbContext _context = context;

    public async Task<OrderItem> GetOrderItem(long orderId, long beerId)
    {
        var orderItem = await _context.OrderItems.FirstOrDefaultAsync(o =>  o.OrderId == orderId && o.BeerId == beerId);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");

        return orderItem;
    }

    public async Task<List<OrderItem>> GetOrderItemsByOrderId(long orderId)
    {
        var orderItem = await _context.OrderItems.Where(o =>  o.OrderId == orderId).ToListAsync();
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");

        return orderItem;
    }

    public async Task<OrderItem> CreateOrder(OrderItem orderItem)
    {
        await _context.OrderItems.AddAsync(orderItem);

        return orderItem;
    }

    public async Task<OrderItem> UpdateOrder(OrderItem orderItem)
    {
        var orderInDb =
            await _context.OrderItems.FirstOrDefaultAsync(o =>
                o.OrderId == orderItem.OrderId && o.BeerId == orderItem.BeerId);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");
        orderInDb.CopyFieldsIfNotNull(orderItem);

        return orderInDb;
    }

    public async Task DeleteOrderItemById(long orderId, long beerId)
    {
        var orderItem = await _context.OrderItems.FirstOrDefaultAsync(o =>  o.OrderId == orderId && o.BeerId == beerId);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");
        _context.OrderItems.Remove(orderItem);
    }

    public async Task<OrderItem> GetOrderItemAsNoTrack(long orderId, long beerId)
    {
        var orderItem = await _context.OrderItems.AsNoTracking().FirstOrDefaultAsync(o =>  o.OrderId == orderId && o.BeerId == beerId);
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");

        return orderItem;
    }

    public async Task<List<OrderItem>> GetOrderItemsByOrderIdAsNoTrack(long orderId)
    {
        var orderItem = await _context.OrderItems.AsNoTracking().Where(o =>  o.OrderId == orderId).ToListAsync();
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");

        return orderItem;
    }

    public async Task DeleteAllOrderItemsByOrderId(long orderId)
    {
        var orderItems = await _context.OrderItems.Where(oi =>  oi.OrderId == orderId).ToListAsync();
        
        //TODO result patter to project
        //if (order is null) return Error.NotFound($"Order with id={id} not found.");
        _context.OrderItems.RemoveRange(orderItems);
    }
}