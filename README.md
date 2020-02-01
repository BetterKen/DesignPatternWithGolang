#  Go 语言设计模式

## 1 设计模式分类

- 创建型模式
    - [简单工厂模式（Simple Factory）](DesignPattern/CreationalPatterns/SimpleFactory)
    - [工厂方法模式（Factory Method）](DesignPattern/CreationalPatterns/Factory)
    - [抽象工厂模式（Abstract Factory）](DesignPattern/CreationalPatterns/AbstractFactory)
    - [创建者模式（Builder）](DesignPattern/CreationalPatterns/Builder)
    - [原型模式（Prototype）](DesignPattern/CreationalPatterns/Prototype)
    - [单例模式（Singleton）](DesignPattern/CreationalPatterns/Singleton)

- 结构型模式
    - [外观模式（Facade）](DesignPattern/StructuralPatterns/Facade)
    - [适配器模式（Adapter）](DesignPattern/StructuralPatterns/Adapter)
    - [代理模式（Proxy）](DesignPattern/StructuralPatterns/Proxy)
    - [组合模式（Composite）](DesignPattern/StructuralPatterns/Composite)
    - [享元模式（Flyweight）](DesignPattern/StructuralPatterns/FlyWeight)
    - [装饰模式（Decorator）](DesignPattern/StructuralPatterns/Decorator)
    - [桥模式（Bridge）](DesignPattern/StructuralPatterns/Bridge)

- 行为型模式
    - [中介者模式（Mediator）](DesignPattern/BehavioralPatterns/Mediator)
    - [观察者模式（Observer）](DesignPattern/BehavioralPatterns/Observer)
    - [命令模式（Command）](DesignPattern/BehavioralPatterns/Command)
    - [迭代器模式（Iterator）](DesignPattern/BehavioralPatterns/Iterator)
    - [模板方法模式（Template Method）](DesignPattern/BehavioralPatterns/Template)
    - [策略模式（Strategy）](DesignPattern/BehavioralPatterns/Strategy)
    - [状态模式（State）](DesignPattern/BehavioralPatterns/State)
    - [备忘录模式（Memento）](DesignPattern/BehavioralPatterns/Memento)
    - [解释器模式（Interpreter）](DesignPattern/BehavioralPatterns/Interpreter)
    - [职责链模式（Chain of Responsibility）](DesignPattern/BehavioralPatterns/Chain)
    - [访问者模式（Visitor）](DesignPattern/BehavioralPatterns/Vistor)



## 2 设计模式功能介绍

- **单例（Singleton）模式**：某个类只能生成一个实例，该类提供了一个全局访问点供外部获取该实例，其拓展是有限多例模式。
- **原型（Prototype）模式**：将一个对象作为原型，通过对其进行复制而克隆出多个和原型类似的新实例。
- **工厂方法（Factory Method）模式**：定义一个用于创建产品的接口，由子类决定生产什么产品。
- **抽象工厂（AbstractFactory）模式**：提供一个创建产品族的接口，其每个子类可以生产一系列相关的产品。
- **建造者（Builder）模式**：将一个复杂对象分解成多个相对简单的部分，然后根据不同需要分别创建它们，最后构建成该复杂对象。
- **代理（Proxy）模式**：为某对象提供一种代理以控制对该对象的访问。即客户端通过代理间接地访问该对象，从而限制、增强或修改该对象的一些特性。
- **适配器（Adapter）模式**：将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类能一起工作。
- **桥接（Bridge）模式**：将抽象与实现分离，使它们可以独立变化。它是用组合关系代替继承关系来实现，从而降低了抽象和实现这两个可变维度的耦合度。
- **装饰（Decorator）模式**：动态的给对象增加一些职责，即增加其额外的功能。
- **外观（Facade）模式**：为多个复杂的子系统提供一个一致的接口，使这些子系统更加容易被访问。
- **享元（Flyweight）模式**：运用共享技术来有效地支持大量细粒度对象的复用。
- **组合（Composite）模式**：将对象组合成树状层次结构，使用户对单个对象和组合对象具有一致的访问性。
- **模板方法（TemplateMethod）模式**：定义一个操作中的算法骨架，而将算法的一些步骤延迟到子类中，使得子类可以不改变该算法结构的情况下重定义该算法的某些特定步骤。
- **策略（Strategy）模式**：定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的改变不会影响使用算法的客户。
- **命令（Command）模式**：将一个请求封装为一个对象，使发出请求的责任和执行请求的责任分割开。
- **职责链（Chain of Responsibility）模式**：把请求从链中的一个对象传到下一个对象，直到请求被响应为止。通过这种方式去除对象之间的耦合。
- **状态（State）模式**：允许一个对象在其内部状态发生改变时改变其行为能力。
- **观察者（Observer）模式**：多个对象间存在一对多关系，当一个对象发生改变时，把这种改变通知给其他多个对象，从而影响其他对象的行为。
- **中介者（Mediator）模式**：定义一个中介对象来简化原有对象之间的交互关系，降低系统中对象间的耦合度，使原有对象之间不必相互了解。
- **迭代器（Iterator）模式**：提供一种方法来顺序访问聚合对象中的一系列数据，而不暴露聚合对象的内部表示。
- **访问者（Visitor）模式**：在不改变集合元素的前提下，为一个集合中的每个元素提供多种访问方式，即每个元素有多个访问者对象访问。
- **备忘录（Memento）模式**：在不破坏封装性的前提下，获取并保存一个对象的内部状态，以便以后恢复它。
- **解释器（Interpreter）模式**：提供如何定义语言的文法，以及对语言句子的解释方法，即解释器。

