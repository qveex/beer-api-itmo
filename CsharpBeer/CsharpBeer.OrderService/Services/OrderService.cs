using Api;
using Grpc.Core;
using Common.Protobuf;
using CsharpBeer.OrderService.Domain.Common.Extensions.Orders;
using CsharpBeer.OrderService.Domain.Common.Interfaces;
using CsharpBeer.OrderService.Domain.Orders;
using CsharpBeer.OrderService.Services.Common.GrpcExtensions;
using Google.Protobuf.WellKnownTypes;
using CsharpBeer.OrderService.Services.Common.Errors;

namespace CsharpBeer.OrderService.Services;

public class OrderService(
    ILogger<OrderService> logger,
    IOrderRepository orderRepository,
    IOrderItemRepository orderItemRepository,
    IUnitOfWork unitOfWork,
    IAuthService authClient,
    ICatalogService catalogClient) : OrderServiceGrpc.OrderServiceGrpcBase
{
    private readonly ILogger<OrderService> _logger = logger;
    private readonly IOrderRepository _orderRepository = orderRepository;
    private readonly IOrderItemRepository _orderItemRepository = orderItemRepository;
    private readonly IUnitOfWork _unitOfWork = unitOfWork;
    private readonly IAuthService _authService = authClient;
    private readonly ICatalogService _catalogService = catalogClient;

    public override async Task<OrderDto> GetOrder(GetOrderRequest request, ServerCallContext context)
    {
        var token = context.GetToken();
        var (userId, _) = await _authService.GetUserInfoAsync(token);
        var order = await _orderRepository.GetOrderByIdAsNoTrack(request.OrderId) ?? throw RpcErrors.OrderNotFound;

        if (order.UserId != userId || !await _authService.IsAdminAsync(userId)) throw RpcErrors.PermissionDenied;

        var orderItems = await _orderItemRepository.GetOrderItemsByOrderIdAsNoTrack(request.OrderId);
        
        var orderDto = order.ToDto();
        var beer = await GetOrderItemsDto(orderItems);
        orderDto.Items.AddRange(beer);
        
        return orderDto;
    }

    public override async Task<ListOrdersResponse> ListOrders(ListOrdersRequest request, ServerCallContext context)
    {
        var token = context.GetToken();
        var (userId, _) = await _authService.GetUserInfoAsync(token);

        var orders = request.UserId is null ? 
            await _orderRepository.ListOrdersByUserIdAsNoTrack(userId) :
            await _authService.IsAdminAsync(userId) ? 
                await _orderRepository.ListOrdersByUserIdAsNoTrack(request.UserId.Value) :
                throw RpcErrors.PermissionDenied; 

        List<OrderDto> orderDtos = [];
        foreach (var order in orders)
        {
            var orderDto = order.ToDto();
            var orderItems = await _orderItemRepository.GetOrderItemsByOrderIdAsNoTrack(order.OrderId);
            var beer = await GetOrderItemsDto(orderItems);
            orderDto.Items.AddRange(beer);
            orderDtos.Add(orderDto);
        }

        var response = new ListOrdersResponse();
        response.Orders.AddRange(orderDtos);
        return response;
    }

    public override async Task<OrderDto> CreateOrder(CreateOrderRequest request, ServerCallContext context)
    {
        var token = context.GetToken();
        var (userId, _) = await _authService.GetUserInfoAsync(token);
        
        var beers = await _catalogService.GetBeersByIdAsync(request.Order.Items);
        var order = request.Order.ToDomain(userId, beers);
        await _orderRepository.CreateOrder(order);

        await _unitOfWork.CommitChangesAsync();

        return order.ToDto();
    }

    public override async Task<OrderDto> UpdateOrder(UpdateOrderRequest request, ServerCallContext context)
    {
        var token = context.GetToken();
        var (userId, _) = await _authService.GetUserInfoAsync(token);

        var orderInApp = await _orderRepository.GetOrderById(request.Order.OrderId) ?? throw RpcErrors.OrderNotFound;
        var orderItemsInApp = await _orderItemRepository.GetOrderItemsByOrderId(orderInApp.OrderId);
        if (orderInApp.UserId != userId || !await _authService.IsAdminAsync(userId)) throw RpcErrors.PermissionDenied;
        
        var orderUpdate = request.Order.ToDomain();
        orderInApp.CopyFieldsIfNotNull(orderUpdate);
        // TODO update orderItems

        await _unitOfWork.CommitChangesAsync();

        var orderDto = orderInApp.ToDto();
        orderDto.Items.AddRange(orderItemsInApp.ToDto());

        return orderDto;
    }

    public override async Task<Empty> DeleteOrder(DeleteOrderRequest request, ServerCallContext context)
    {
        var token = context.GetToken();
        var (userId, _) = await _authService.GetUserInfoAsync(token);

        var orderInApp = await _orderRepository.GetOrderById(request.OrderId) ?? throw RpcErrors.OrderNotFound;
        if (orderInApp.UserId != userId) throw RpcErrors.PermissionDenied;
        
        await _orderRepository.DeleteOrderById(orderInApp.OrderId);
        await _orderItemRepository.DeleteAllOrderItemsByOrderId(orderInApp.OrderId);
        await _unitOfWork.CommitChangesAsync();
        
        return new Empty();
    }

    private async Task<List<OrderItemDto>> GetOrderItemsDto(IEnumerable<OrderItem> orderItems) => [.. 
        (await Task.WhenAll(orderItems.Select(async oi => {
            var orderItemDto = oi.ToDto();

            var beerResponse = await _catalogService.GetBeerByIdAsync(oi.BeerId);

            orderItemDto.Beer = beerResponse;
            return orderItemDto;
        })))
    ];

    private List<OrderItem> AddBeersToOrderItems(IEnumerable<Beer> beers, IEnumerable<CreateOrderItemDto> orderItems) => 
        beers.Zip(orderItems).Select(t => t.Second.ToDomain(t.First)).ToList();

    /* List<OrderItemDto> listOrderItems = [];
     foreach (var orderItem in orderItems)
     {
         var orderItemDto = orderItem.ToDto();

         var beerResponse = await _catalogService.GetBeerByIdAsync(orderItem.BeerId);

         //TODO if beer not exists -> WHAT TO DO?
         //if (beerResponse is null) continue || Error.NotFound();

         orderItemDto.Beer = beerResponse;
         listOrderItems.Add(orderItemDto);
     }

    return listOrderItems;*/
}