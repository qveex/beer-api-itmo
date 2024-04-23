using Auth;
using CsharpBeer.OrderService.Domain.Common.Interfaces;
using CsharpBeer.OrderService.Services.Common.Errors;
using AuthC = Auth.Auth;

namespace CsharpBeer.OrderService.Infrastructure.Auth;

public class AuthService(AuthC.AuthClient authClient) : IAuthService
{
    private readonly AuthC.AuthClient _authClient = authClient;

    public async Task<(long UserId, string Email)> GetUserInfoAsync(string? token)
    {
        if (token is null) throw RpcErrors.InvalidToken;

        var request = new GetUserInfoRequest() { Token = token };
        var response = await _authClient.GetUserInfoAsync(request);

        return (response.UserId, response.Email);
    }

    public async Task<bool> IsAdminAsync(long userId)
    {
        var request = new IsAdminRequest() { UserId = userId };
        var response = await _authClient.IsAdminAsync(request);

        return response.IsAdmin;
    }
}