using Api;
using Common.Protobuf;

namespace CsharpBeer.OrderService.Domain.Common.Interfaces;

public interface ICatalogService
{
    Task<Beer> GetBeerByIdAsync(long beerId);
    Task<List<Beer>> GetBeersByIdAsync(IEnumerable<CreateOrderItemDto> ids);
}