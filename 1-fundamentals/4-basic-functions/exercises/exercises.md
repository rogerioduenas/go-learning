1-Create a function with two parameters of a small numeric type and a return value of the same type, performing a simple operation between the two values.

2-Create a variable with an explicit type and assign to it the result of the function created earlier, within the program's entry point (main).

3-Display the value returned by the previous function using the standard output library.

4-Declare another function that returns two results of the same type, representing two different operations with the same arguments.

5-Create a function that receives two numerical parameters and returns two results of the same type, representing two different operations with those values. At the program's entry point (main), call this function and store the two results in separate variables.

6-Create a function that receives two numerical parameters and returns two results of the same type, representing two different operations with those values. At the program's entry point (main), call this function, store the two results in separate variables, and display both values ​​in a single output statement.

7-Create an anonymous function assigned to a variable, which receives a text argument and returns the same value after displaying it.

8-Create an anonymous function, store the return value in a variable, and display the returned value.

9-Call a function that returns multiple values, but discard one of the results without storing it.

10-Reorganize the values ​​returned by the multiple-return function, reversing the order and displaying the new results.

11-Create a new function that, internally, uses the existing function to perform an additional calculation before returning the result.

12-Develop a function that receives another function as a parameter and invokes it internally, displaying the returned result.

13-Implement an anonymous function that uses a previously defined function within its body before returning a value.

14-Create a function that, when called, returns another function. The returned function should be able to be invoked later to calculate and provide a final result.

15-Combine named and anonymous function calls so that the value returned by one serves as an argument for the other.

16-Create a control structure that chooses which function to execute based on a numerical condition, displaying the corresponding result.

17-Create a function that receives a numerical value and returns a function capable of receiving two numbers and performing a mathematical operation combining the original value with the two parameters. Execute the returned function with different arguments.

18-Implement a function that takes another function as an argument, where the latter function returns two values. Inside the first function, call the received function, add the results, and return only the sum.

19-Develop a function that receives a multiplication factor and returns a function that accepts a slice of integers, applying the factor to each element and returning a new slice.

20-Create a function that takes an integer and a conditional function (func(int) bool) as arguments. If the conditional function returns true, perform an operation and return the result; otherwise, return another fixed value.

21-Develop a function that receives a list of functions (func(int) int) and an initial value. Each function should be applied in sequence, so that the result of one is used as the input for the next. Finally, return the result obtained after applying all the functions.

22-Create an anonymous function stored in a variable that returns another anonymous function. Call the inner function and display the result.

23-Implement a function that takes two integers as input and returns a specific function depending on a condition (e.g., if greater than 10, return addition; if less than 10, return subtraction). Call the returned function with the original arguments.

24-Create a function that takes a function as an argument (func(int) int) and an integer. Internally, add 5 to the integer, call the callback function, and return the result. Demonstrate how the outer function "captures" the modified variable before passing it to the callback.

25-Develop a function that returns an inner function capable of accumulating values ​​added to it with each call, returning the accumulated total. Demonstrate with multiple sequential calls.