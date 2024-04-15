using Api.Auth;
using Api.Catalog;
using Grpc.Core;
using Common.Protobuf;
using CsharpBeer.OrderService.Domain.Common.Extensions.Orders;
using CsharpBeer.OrderService.Domain.Common.Interfaces;
using CsharpBeer.OrderService.Domain.Orders;
using Google.Protobuf.WellKnownTypes;

namespace CsharpBeer.OrderService.Services;

public class OrderService(
    ILogger<OrderService> logger,
    IOrderRepository orderRepository,
    IOrderItemRepository orderItemRepository,
    Auth.AuthClient authClient,
    Catalog.CatalogClient catalogClient) : OrderServiceGrpc.OrderServiceGrpcBase
{
    private readonly ILogger<OrderService> _logger = logger;
    private readonly IOrderRepository _orderRepository = orderRepository;
    private readonly IOrderItemRepository _orderItemRepository = orderItemRepository;
    private readonly Auth.AuthClient _authClient = authClient;
    private readonly Catalog.CatalogClient _catalogClient = catalogClient;

    public override async Task<OrderDto> GetOrder(GetOrderRequest request, ServerCallContext context)
    {
        //TODO Check if user is the owner of the order
        
        
        var order = await _orderRepository.GetOrderById(request.OrderId);
        var orderItems = await _orderItemRepository.GetOrderItemsByOrderId(request.OrderId);
        
        var orderDto = order.ToDto();
        orderDto.Items.AddRange(GetOrderItemsDto(orderItems));
        
        return orderDto;
    }

    public override async Task<ListOrdersResponse> ListOrders(ListOrdersRequest request, ServerCallContext context)
    {
        //TODO Extract userId from token to get orders
        //var token = context.RequestHeaders.GetValue("Authentication")?.Split("Bearer ")[0];
        //var userId = token.ExtractClaims();

        var token = request.Token;
        var userId = long.Parse(token);

        var orders = await _orderRepository.ListOrdersByUserId(userId);

        List<OrderDto> orderDtos = [];
        foreach (var order in orders)
        {
            var orderDto = order.ToDto();
            var orderItems = await _orderItemRepository.GetOrderItemsByOrderId(order.OrderId);
            orderDto.Items.AddRange(GetOrderItemsDto(orderItems));
            orderDtos.Add(orderDto);
        }

        var response = new ListOrdersResponse();
        response.Orders.AddRange(orderDtos);
        return response;
    }

    public override Task<OrderDto> CreateOrder(CreateOrderRequest request, ServerCallContext context)
    {
        //TODO validations checks
        
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

    private IEnumerable<OrderItemDto> GetOrderItemsDto(IEnumerable<OrderItem> orderItems)
    {
        List<OrderItemDto> listOrderItems = [];
        foreach (var orderItem in orderItems)
        {
            var orderItemDto = orderItem.ToDto();
            
            var catalogRequest = new GetBeerRequest { BeerId = orderItem.BeerId };
            var beerResponse = _catalogClient.GetBeer(catalogRequest);
            
            //TODO if beer not exists -> WHAT TO DO?
            //if (beerResponse.Beer is null) continue || Error.NotFound();

            orderItemDto.Beer = beerResponse.Beer;
            listOrderItems.Add(orderItemDto);
        }

        return listOrderItems;
    }
}