using Grpc.Core;

namespace CsharpBeer.OrderService.Services.Common.GrpcExtensions;

public static class GrpcExtensions
{
    public static string? GetToken(this ServerCallContext context) => 
        context.RequestHeaders.GetValue("Authorization")?.Split("Bearer ", StringSplitOptions.RemoveEmptyEntries)[0];     
}