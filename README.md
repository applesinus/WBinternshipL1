[Jump to English](#English)

<a name="Russian"></a>
# Русский
<p id="ru"><h3>Задачи Lv1 на стажировке Wildberries</h3></p>

<p>Каждая задача - один файл с соответствующим названием, запускаются они все из файла main через интерактивный ввод.</p>
<ol>
  <li>Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)</li>
  <li>Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout</li>
  <li>Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте</li>
  <li>Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров</li>
  <li>Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться</li>
  <li>Реализовать все возможные способы остановки выполнения горутины</li>
  <li>Реализовать конкурентную запись данных в map</li>
  <li>Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0</li>
  <li>Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout</li>
  <li>Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc</li>
  <li>Реализовать пересечение двух неупорядоченных множеств</li>
  <li>Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество</li>
  <li>Поменять местами два числа без создания временной переменной</li>
  <li>Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}</li>
  <li>
  <p>К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации</p>
<pre>
<code>var justString string
  func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
  }
  func main() {
    someFunc()
  }</code>
</pre>
  </li>
  <li>Реализовать быструю сортировку массива (quicksort) встроенными методами языка</li>
  <li>Реализовать бинарный поиск встроенными методами языка</li>
  <li>Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика</li>
  <li>Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode</li>
  <li>Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow»</li>
  <li>Реализовать паттерн «адаптер» на любом примере</li>
  <li>Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20</li>
  <li>Удалить i-ый элемент из слайса</li>
  <li>Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором</li>
  <li>Реализовать собственную функцию sleep</li>
  <li>Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой. Например: abcd — true, abCdefAaf — false, aabcd — false</li>
</ol>

[Перейти к русскому](#Russian)

<a name="English"></a>
# English
<p id="ru"><h3>Exercises on Lv1 at the Wildberries internship</h3></p>

<p>Each task is one file with the corresponding name, they are all launch from the main file through interactive input.</p>
<ol>
   <li>The Human structure is given (with an arbitrary set of fields and methods). Implement embedding of methods in the Action structure from the parent Human structure (analogous to inheritance)</li>
   <li>Write a program that competitively calculates the value of the squares of numbers taken from the array (2,4,6,8,10) and outputs their squares to stdout</li>
   <li>Implement constant writing of data to the channel (main thread). Implement a set of N workers that read arbitrary data from a channel and output it to stdout. It is necessary to be able to select the number of workers at startup</li>
   <li>The program should end by pressing Ctrl+C. Select and justify a method for completing the work of all workers</li>
   <li>Develop a program that will sequentially send values to the channel, and read from the other side of the channel. After N seconds the program should exit</li>
   <li>Implement all possible ways to stop the execution of a goroutine</li>
   <li>Implement concurrent data recording in map</li>
   <li>An int64 variable is given. Develop a program that sets the i-th bit to 1 or 0</li>
   <li>Develop a number pipeline. Two channels are given: numbers (x) from the array are written into the first, the result of the operation x*2 is written into the second, after which the data from the second channel should be output to stdout</li>
   <li>The sequence of temperature fluctuations is given: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Combine these values into groups in 10 degree increments. The sequence in subsets is not important
Example: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc</li>
   <li>Implement the intersection of two unordered sets</li>
   <li>There is a sequence of strings - (cat, cat, dog, cat, tree) create your own set for it</li>
   <li>Swap two numbers without creating a temporary variable</li>
   <li>Develop a program that can determine the type of variable in runtime: int, string, bool, channel from a variable of type interface{}</li>
   <li>
   <p>What negative consequences can this piece of code lead to, and how can I fix it? Please provide a correct implementation example</p>
<pre>
<code>var justString string
  func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
  }
  func main() {
    someFunc()
  }</code>
</pre>
   </li>
   <li>Implement quick array sorting (quicksort) using built-in language methods</li>
   <li>Implement binary search using built-in language methods</li>
   <li>Implement a counter structure that will be incremented in a competitive environment. Upon completion, the program should print the final counter value</li>
   <li>Develop a program that reverses the string supplied for the move (for example: “headfish - abyrvalg”). Characters can be unicode</li>
   <li>Develop a program that reverses words in a string. Example: “snow dog sun - sun dog snow”</li>
   <li>Implement the “adapter” pattern using any example</li>
   <li>Develop a program that multiplies, divides, adds, subtracts two numeric variables a,b, the value of which is > 2^20</li>
   <li>Remove the i-th element from the slice</li>
   <li>Develop a program for finding the distance between two points, which are represented as a Point structure with encapsulated parameters x,y and a constructor</li>
   <li>Implement your own sleep function</li>
   <li>Develop a program that checks that all characters in a string are unique (true if unique, false etc). The check function must be case insensitive. For example: abcd - true, abCdefAaf - false, aabcd - false</li>
</ol>
