# Паттерн "Стратегия"  
## Определение  
Определяет семейство алгоритмов, инкапсулирует каждый
из них и обеспечивает их взаимозаменяемость. Он позволяет
модифицировать алгоритмы  независимо от их использования на
стороне клиента.  
  
## Используемые принципы  
1. Инкапсулируйте то, что изменяется  
2. Отдавайте предпочтение композиции перед наследованием  
3. Программируйте на уровне интерфейсов, а не реализаций  
  
## Описание примера  
В примере показана реализация данного паттерна на различных
видах уток. Существуют, помимо обычных, еще и
деревянные и резиновые утки, а также утки-приманки.  
У всех уток есть следующие методы:  
1. Display - внешний вид утки  
2. Swim - способность плавать  
3. Quack - способность крякать  
4. Fly - способность летать  
  
Все утки выглядят по-разному, поэтому метод Display
необходимо реализовывать для каждого вида отдельно. Метод
Swim для всех уток один и наследуется от базовой структуры.
А вот методы Quack и Fly для разных видов уток могут как
различаться, так и совпадать. Если применять обычное
наследование, то придется во многих местах эти методы
переопределять, а если делать для каждого вида уток свои
методы, то будет дублирование кода. В любом случае
вероятность будущих ошибок стремительно возрастает,
да и такой код очень тяжело поддерживать.  
  
Данную проблему и решает паттерн "Стратегия".  
Создаются и инкапсулируются отдельные семейства алгоритмов
способностей крякать и летать. В базовой структуре утки
применяется не наследование, а композиция. В данную
структуру добавляются свойства с типами интерфейсов
способностей крякать и летать, а у каждого вида утки в
данные свойства записываются конкретные реализации этих
интерфейсов. Теперь утки вызывают не конкретные методы
способностей крякать и летать, а методы, которые в свою
очередь вызывают методы, описанные интерфейсами данных
способностей.  
  
Паттерн позволяет также устанавливать на лету другие
реализации данных интерфейсов, а также вносить изменения в
поведение уток, не затрагивая основной код.  