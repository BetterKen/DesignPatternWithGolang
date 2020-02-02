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



## 3 六大原则



### 3.1 单一职责原则(Single Responsibility Principle)

定义：就一个类而言， 应该仅有一个引起它变化的原因。

单一职责的划分界限并不是那么清晰，很多时候需要靠个人经验界定。当然最大的问题就是对职责的定义，什么是类的职责，以及怎么划分类的职责。

### 3.2 开放封闭原则（Open Close Principle）

定义：类、模块、函数等应该是可以拓展的，但是不可修改。

开闭原则指导我们，当软件需要变化时，应该尽量通过拓展的方式来实现变化，而不是通过修改已有代码来实现。这里的“应该尽量”4个字说明OCP原则并不是说绝对不可以修改原始类的。当我们嗅到原来的代码“腐化气味”时，应该尽早地重构，以便使代码恢复到正常的“进化”过程，而不是通过集成等方式添加新的实现，这会导致类型的膨胀以及历史遗留代码的冗余。因此，在开发过程中需要自己结合具体情况进行考量，是通过修改旧代码还是通过继承使得软件系统更稳定、更灵活，在保证去除“代码腐化”的同时，也保证原有模块的正确性。

### 3.3 里氏替换原则（Liskov Substitution Principle）

注：它是开闭原则的具体实现手段之一，它的核心原理是抽象

定义：所有引用基类的地方必须能透明地使用其子类的对象。

里氏替换原则的核心原理是抽象，抽象又依赖于继承这个特性，在OOP中，继承的优缺点相当明显，有点如下：

- 代码重用，减少创建类成本，每个子类拥有父类的属性和方法；
- 子类和父类基本相似，但又与父类有所区别；
- 提高代码的可拓展性。

继承的缺点：

- 继承是侵入性的，只要继承就必须拥有弗雷的所有属性和方法
- 可能造成子类代码的冗余、灵活性降低，因为子类必须拥有弗雷的属性和方法。

开闭原则和里氏替换原则往往是生死相依、不离不弃的，通过里氏替换来达到对扩展的开发，对修改的关闭效果。

### 3.4 依赖倒置原则（Dependence Inversion Principle）

注：关系到系统的可拓展性、拥抱变化的能力、开闭原则

定义：高层模块不应该依赖于低层模块，两者都应该依赖于抽象。抽象不应该依赖于细节，细节应该依赖于抽象。

java中抽象指接口或抽象类，两者都不能直接被实例化的；细节就是实现类，实现接口或者集成抽象类而产生的也就细节，也就是可以可以加上yige 关键字new产生的对象。高层模块就是调用端，低层模块就是具体实现类。依赖倒置原则在java中表现就是，模块间依赖通过抽象发生，实现类之间不发生直接依赖关系，其依赖关系是通过接口或者抽象类产生的。如果类与类直接依赖细节，那么久会直接耦合。如此一来当修改时，就会同时修改依赖者代码，这样限制了可拓展性。

### 3.5 接口隔离原则（InterfaceSegregation Principles）

注：最小化， 减少依赖从而降低变更的风险。

定义：一个类对另一个类的依赖应该建立在最小的接口上。

建立单一接口，不要建立庞大臃肿接口；尽量细化接口，接口中方法尽量少。也就是说，我们要为各个类建立专用的接口，而不要试图建立一个很庞大的接口供所有依赖它的类调用。

- 接口尽量小，但是要有限度。对接口进行细化可以提高程序设计的灵活性；但是如果过小，则会造成接口数量过多，使设计复杂化。所以，一定要适度。
- 为依赖接口的类定制服务，只暴露给调用的类需要的方法，它不需要的方法则隐蔽起来。只有专注得为一个模块提供定制服务，才能建立最小的依赖关系。
- 提高内聚，减少对外交互。接口方法尽量少用public修饰。接口是对外的承诺，承诺越少对系统开发越有利，变更风险也会越少。

以上（单一职责、开闭原则、里氏替换、接口隔离、依赖倒置）五个原则被Bob大叔在21世纪早期定义为SOLID原则。作为面向对象编程的5个基本原则，当这些原则被一起使用时，它们使得一个软件系统更清晰、简单，最大程度地拥抱变化。

### 3.6  迪米特原则（Law of Demeter）/最少知识原则

注：通过引入一个合理的第三者降低现有对象之间的耦合度。

定义：一个软件实体应当尽可能少地与其他实体发生相互作用。

一个类应该对自己需要耦合或者调用的类知道最少， 类的内部如何实现与调用者或者依赖关系越密切，耦合度越大，当一个类发生变化时，对另一个类的影响也越大。

- 在类的划分上，应当尽量创建松耦合的类。类之间的耦合度约低，就越有利于服用。一个处于松耦合中的类一旦被修改，则不会对关联的类造成太大的波及。
- 在类的机构设计上， 每一个类都应当尽量降低其成员变量和成员函数的访问权限。
- 在对其他类的引用上， 一个类对其他对象的引用应当降到最低。

