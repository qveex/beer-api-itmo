using Grpc.Core;
using Common.Protobuf;
using Google.Protobuf.WellKnownTypes;

namespace CsharpBeer.OrderService.Services;

public class OrderService(ILogger<OrderService> logger) : OrderServiceGrpc.OrderServiceGrpcBase
{
    private readonly ILogger<OrderService> _logger = logger;

    public override Task<OrderDto> GetOrder(GetOrderRequest request, ServerCallContext context)
    {
        //TODO
        return base.GetOrder(request, context);
    }

    public override Task<ListOrdersResponse> ListOrders(ListOrdersRequest request, ServerCallContext context)
    {
        //TODO
        return base.ListOrders(request, context);
    }

    public override Task<OrderDto> CreateOrder(CreateOrderRequest request, ServerCallContext context)
    {
        //TODO
        return base.CreateOrder(request, context);
    }

    public override Task<OrderDto> UpdateOrder(UpdateOrderRequest request, ServerCallContext context)
    {
        //TODO
        return base.UpdateOrder(request, context);
    }

    public override Task<Empty> DeleteOrder(DeleteOrderRequest request, ServerCallContext context)
    {
        //TODO
        return base.DeleteOrder(request, context);
    }
}