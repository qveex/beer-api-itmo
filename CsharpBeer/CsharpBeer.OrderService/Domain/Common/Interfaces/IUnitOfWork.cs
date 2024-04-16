namespace CsharpBeer.OrderService.Domain.Common.Interfaces;

public interface IUnitOfWork
{
    Task CommitChangesAsync();   
}