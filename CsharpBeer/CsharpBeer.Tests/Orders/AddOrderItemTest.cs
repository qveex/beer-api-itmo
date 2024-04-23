using CsharpBeer.Tests.Common;
using FluentAssertions;

namespace CsharpBeer.Tests.Orders;

public class AddOrderItemTest
{
    [Fact]
    public void AddOrderItem_WhenOrderItemWithThisIdAlreadyExists_ShouldDoNothing()
    {
        //Arrange
        var order = OrderFactory.CreateOrder();

        var orderItem1 = OrderItemFactory.CreateOrderItem();
        var orderItem2 = OrderItemFactory.CreateOrderItem();
        
        //Act
        order.AddOrderItem(orderItem1);
        order.AddOrderItem(orderItem2);

        //Assertion
        order.Items.Count.Should().Be(1);
    }
    
    [Fact]
    public void AddOrderItem_WhenAddNewItem_ShouldDefineOrderIdInOrderItem()
    {
        //Arrange
        const int orderId = 1;
        
        var order = OrderFactory.CreateOrder();
        order.OrderId = orderId;

        var orderItem1 = OrderItemFactory.CreateOrderItem();
        
        //Act
        order.AddOrderItem(orderItem1);

        //Assertion
        order.Items.First().OrderId.Should().Be(orderId);
    }
    
    [Fact]
    public void AddOrderItem_WhenAddNewItem_ShouldRecalculateTotal()
    {
        //Arrange
        const double beerPrice = 50d;
        const int quantity = 2;
        
        var order = OrderFactory.CreateOrder();

        var orderItem1 = OrderItemFactory.CreateOrderItem(quantity: quantity, price: beerPrice);
        
        //Act
        order.AddOrderItem(orderItem1);

        //Assertion
        order.Total.Should().Be(quantity * beerPrice);
    }
}