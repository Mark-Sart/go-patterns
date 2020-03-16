# Паттерн "Наблюдатель"  
## Определение  
Определяет отношение "один-ко-многим" между объектами таким
образом, что при изменении состояния одного объекта
происходит автоматическое оповещение и обновление всех
зависимым объектов.  
  
## Используемые принципы  
1. Инкапсулируйте то, что изменяется  
2. Отдавайте предпочтение композиции перед наследованием  
3. Программируйте на уровне интерфейсов, а не реализаций  
4. Стремитесь к слабой связанности взаимодействующих
объектов.  
  
## Описание примера  
В примере показана реализация данного паттерна в приложении
для метеостанции.  
  
Необходимо реализовать такое приложение, которое будет
отображать информацию на трех визуальных элементах.
К тому же необходимо добавить возможность реализовывать
визуальные элементы сторонними разработчиками.
Так как в приложении используется три визуальных элемента,
к тому же, еще планируется добавлять новые, нельзя напрямую
обновлять информацию в этих элементах, так как придется
вносить правки в основной код приложения при каждом
добавлении нового визуального элемента. Поддержка данного
приложения становится очень ресурсозатратной.  
  
В данной ситуации и применяется паттерн "Наблюдатель".  
1. Создаются интерфейсы наблюдателя и субъекта, располагающего
данными;  
2. Реализуются методы регистрации и удаления
наблюдателей у субъекта;  
3. Реализуется оповещение зарегистрированных наблюдателей
у субъекта;  
4. Реализуется активная доставка информации прямо в
оповещении или же get-методы для запроса необходимой
информации у субъекта (считается, что лучше использовать запросы, так
как различным наблюдателям нужны различные наборы информации
субъекта);  
5. Реализуется обновление информации у наблюдателя.  
  
Для получения информации объект регистрируется у субъекта
в качестве наблюдателя. Для прекращения получения информации
объект удаляется из списка наблюдателей у субъекта.  
  
Паттерн использует слабую связь между субъектом и
наблюдателями. Это реализовано за счет то, что субъект и
наблюдатель общаются только интерфейсами. Друг о друге
данные объекты ничего не знают.  