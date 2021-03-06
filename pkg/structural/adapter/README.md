# Паттерн "Адаптер"
## Определение
Преобразует интерфейс класса к другому интерфейсу, на который рассчитан клиент.
Адаптер обеспечивает совместную работу классов, невозможную в обычных условиях
из-за несовместимости интерфейсов. 

## Описание примера
В данном пакете представлено два примера. В первом примере интерфейс утки приводится
к интерфейсу индюшки и наоборот. Во втором примере интерфейс перечисления приводится
к интерфейсу итератора и наоборот. Это достигается за счет отдельной структур,
которая реализует целевой интерфейс и содержит в себе посредством композиции
адаптируемую структуру. При вызове метод адаптера, вызов делегируется адаптируемой
структуре при помощи ряда преобразований.

Таким образом, клиент, рассчитанный на работу только с итераторами, может работать
и с перечислениями при помощи соответствующего адаптера.

## Ключевые моменты
1. Если понадобится использовать существующий класс с неподходящим интерфейсом -
используйте адаптер.
2. Адаптер приводит интерфейс к тому виду, на который рассчитан клиент.
3. Объем работы по реализации адаптера зависит от размера и сложности целевого
интерфейса.
4. Существует две разновидности адаптеров: адаптеры объектов и адаптеры классов.
Для адаптеров классов необходимо множественное наследование.
