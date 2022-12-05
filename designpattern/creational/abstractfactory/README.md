# AbstractFactory

## Intent

Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying 
their concrete classes.

![alt text](https://refactoring.guru/images/patterns/content/abstract-factory/abstract-factory-en.png)

## Problem 

Imagine that you’re creating a furniture shop simulator. Your code consists of classes that represent:

A family of related products, say: Chair + Sofa + CoffeeTable.

Several variants of this family. For example, products Chair + Sofa + CoffeeTable are available in these variants: 
Modern, Victorian, ArtDeco.

![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/problem-en.png)


You need a way to create individual furniture objects so that they match other objects of the same family. Customers 
get quite mad when they receive non-matching furniture.

Also, you don’t want to change existing code when adding new products or families of products to the program. Furniture 
vendors update their catalogs very often, and you wouldn’t want to change the core code each time it happens.

## Solution 

The first thing the Abstract Factory pattern suggests is to explicitly declare interfaces for each distinct product of
the product family (e.g., chair, sofa or coffee table). Then you can make all variants of products follow those 
interfaces. For example, all chair variants can implement the Chair interface; all coffee table variants can implement 
the CoffeeTable interface, and so on.

![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/solution1.png)

The next move is to declare the Abstract Factory—an interface with a list of creation methods for all products that are 
part of the product family (for example, createChair, createSofa and createCoffeeTable). These methods must return 
abstract product types represented by the interfaces we extracted previously: Chair, Sofa, CoffeeTable and so on.

![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/solution2.png)

Now, how about the product variants? For each variant of a product family, we create a separate factory class based on 
the AbstractFactory interface. A factory is a class that returns products of a particular kind. For example, the 
ModernFurnitureFactory can only create ModernChair, ModernSofa and ModernCoffeeTable objects.

The client code has to work with both factories and products via their respective abstract interfaces. This lets you 
change the type of factory that you pass to the client code, as well as the product variant that the client code 
receives, without breaking the actual client code.

![alt text](https://refactoring.guru/images/patterns/content/abstract-factory/abstract-factory-comic-2-en.png)

Say the client wants a factory to produce a chair. The client doesn’t have to be aware of the factory’s class, nor does 
it matter what kind of chair it gets. Whether it’s a Modern model or a Victorian-style chair, the client must treat all 
chairs in the same manner, using the abstract Chair interface. With this approach, the only thing that the client knows 
about the chair is that it implements the sitOn method in some way. Also, whichever variant of the chair is returned, 
it’ll always match the type of sofa or coffee table produced by the same factory object.

There’s one more thing left to clarify: if the client is only exposed to the abstract interfaces, what creates the 
actual factory objects? Usually, the application creates a concrete factory object at the initialization stage. Just 
before that, the app must select the factory type depending on the configuration or the environment settings.

## Structure

![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/structure.png)

## Pseudocode

This example illustrates how the Abstract Factory pattern can be used for creating cross-platform UI elements without 
coupling the client code to concrete UI classes, while keeping all created elements consistent with a selected operating 
system.

![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/example.png)

The same UI elements in a cross-platform application are expected to behave similarly, but look a little bit different 
under different operating systems. Moreover, it’s your job to make sure that the UI elements match the style of the 
current operating system. You wouldn’t want your program to render macOS controls when it’s executed in Windows.

The Abstract Factory interface declares a set of creation methods that the client code can use to produce different 
types of UI elements. Concrete factories correspond to specific operating systems and create the UI elements that match 
that particular OS.

It works like this: when an application launches, it checks the type of the current operating system. The app uses this 
information to create a factory object from a class that matches the operating system. The rest of the code uses this 
factory to create UI elements. This prevents the wrong elements from being created.

With this approach, the client code doesn’t depend on concrete classes of factories and UI elements as long as it works 
with these objects via their abstract interfaces. This also lets the client code support other factories or UI elements 
that you might add in the future.

As a result, you don’t need to modify the client code each time you add a new variation of UI elements to your app. You 
just have to create a new factory class that produces these elements and slightly modify the app’s initialization code 
so it selects that class when appropriate.


## Applicability

Use the Abstract Factory when your code needs to work with various families of related products, but you don’t want it 
to depend on the concrete classes of those products—they might be unknown beforehand or you simply want to allow for 
future extensibility.

The Abstract Factory provides you with an interface for creating objects from each class of the product family. As long 
as your code creates objects via this interface, you don’t have to worry about creating the wrong variant of a product 
which doesn’t match the products already created by your app.

Consider implementing the Abstract Factory when you have a class with a set of Factory Methods that blur its primary 
responsibility.

In a well-designed program each class is responsible only for one thing. When a class deals with multiple product types, 
it may be worth extracting its factory methods into a stand-alone factory class or a full-blown Abstract Factory 
implementation.

## How to Implement

Map out a matrix of distinct product types versus variants of these products.

Declare abstract product interfaces for all product types. Then make all concrete product classes implement these 
interfaces.

Declare the abstract factory interface with a set of creation methods for all abstract products.

Implement a set of concrete factory classes, one for each product variant.

Create factory initialization code somewhere in the app. It should instantiate one of the concrete factory classes, 
depending on the application configuration or the current environment. Pass this factory object to all classes that 
construct products.

Scan through the code and find all direct calls to product constructors. Replace them with calls to the appropriate 
creation method on the factory object.





