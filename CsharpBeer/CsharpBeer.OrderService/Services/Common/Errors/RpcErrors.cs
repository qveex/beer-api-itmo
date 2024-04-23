using Grpc.Core;

namespace CsharpBeer.OrderService.Services.Common.Errors;

public static class RpcErrors
{
    public static RpcException InvalidToken => new(new Status(StatusCode.Unauthenticated, "Token is invalid."));
    public static RpcException PermissionDenied => new(new Status(StatusCode.PermissionDenied, "User is not owner or admin."));
    public static RpcException OrderNotFound => new(new Status(StatusCode.NotFound, "Order is not found."));
}