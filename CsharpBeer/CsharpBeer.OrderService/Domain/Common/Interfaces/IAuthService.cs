namespace CsharpBeer.OrderService.Domain.Common.Interfaces;

public interface IAuthService
{
    Task<(long UserId, string Email)> GetUserInfoAsync(string? token);
    Task<bool> IsAdminAsync(long userId);
}