1-Create a simple structure that represents a generic entity with two basic fields, one text and one numeric.

2-Create a second structure that incorporates the first, representing a compositional relationship between types.

3-At the program's entry point, declare a base structure variable and assign values ​​to its fields.

4-Create a base structure called person and another composite structure called student that contains person as a field. At the program's entry point, declare a variable for student and assign it the previously created instance of person.

5-Display the full value of a compound structure in standard output.

6-Access a field inherited from the base structure directly through the composite structure and display its value.

7-Create a base structure representing a generic type of worker and another derived structure that adds specific fields, such as job title and workplace.

8-Instantiate the base structure separately and then use it to create the derived structure.

9-Display both inherited and derived structure-specific fields in the standard output.

10-Create a new instance of the derived structure by initializing all fields directly, without creating an intermediate variable for the base structure.

11-Change the value of an inherited field after the derived structure has been created and display the updated result.

12-Create a function that receives a derived structure and displays a field inherited from the base structure.

13-Develop a function that receives the base structure and returns a derived structure containing that embedded instance along with additional fields.

14-Create a function that receives the derived structure and modifies both inherited fields and specific fields before returning it.

15-Implement a chain of structures, where a base structure is incorporated into an intermediate one, which in turn is incorporated into a third, representing multiple levels of composition.

16-Display the value of a field belonging to the innermost structure using the outermost structure.

17-Create a function that receives two employee-type structures, uses the inherited data from each—such as name and job title—and combines this information to generate a new salary-type structure, incorporating the inherited fields and adding a salary value.

18-Simulate a type hierarchy where two different structures share the same base structure and create instances for both, showing how inherited fields are reused.