# My source Code for Stephen Grider GO udemy course

Source Code for GO course 

## Notes from course
### Cards project section
#### Variables
- define variable in multiple ways with identifying type or using :=
    - var card string = "Ace of Spades"
    - var card := "Acre of Spades"
- can redefine variable with =
- can initialize variable without assigning a value
- can initialize variable outside of function without assigning value
- can only assign value to variable within function

Array - Fixed Length of list
Slice - an arrray that can grow or shrink
- indexed from 0 
- to get a subset range use example [0:4]
    - can optionally leave off 1 number of the range [:4] to start from begining or [0:] to go to the end of the slice

Functions/Return Types
- must define type of data that is returned from function
- function receivers are conventially called with a single letter variable
- define a variable as an _ if you don't want to use it to avoid compile error
    - ex: for _, suits := range cardSuits
- must define an arguments type of data if not defined previously
    - ex: func(d, deck, handSize int) 

#### Testing
- tests functions are started with uppercase letters
    - ex: TestNewDeck(t *testing.T)
        - t is the test handler
- must run go mod init <project-directory-name-here>  within project directory for go test to work 
    - need to provide a module path for test to run 
    - ex go mod init cards

### structs project section
##### Structs
- structs are like an object in javascript
    - dictionary in python 
    - or hash in ruby
- define each type of data in struct
    - type person struct 
        - firstName string
        - lastName string
- can nest structs inside of structs
    - embedded struct 
- structs can be a reciever to a function

##### Pointers
- the & operator turns a value into a pointer
    - name := joe
    - fmt.Println(&name) - prints memory address
- the * operator turns a pointer into a value

###### Reference vs Value types

Value types - use pointers to change these things in a function
- int
- float 
- string 
- bool
- structs 
Reference types - don't worry about pointers with these
- slices
- maps
- channels
- pointers
- functions

### Maps section
- like structs Maps are similiar to an object in javascipt
    - hash in ruby
    - dictionary in python
##### Maps Vs Structs

###### Maps
- All keys must be the same type
- All values must be the same type
- keys are indexed so we can iterate over them 
- used to represent a collection of related properties 
- dont need to know all the keys at compile time
- reference type

###### Structs
- Values can be of different type
- Keys dont support indexing
- you need to know all the different fields at compile time 
- Use to represent a "thing" with a lot of different properties
- Value type 

### Interfaces
Interfaces are for guiding you down the correct path they don't necessarily help you write correct code
-  interfaces are NOT generic types
- interfaces are 'implicit'
- they are a contract to help us manage types
    - garbage in => garbage out
    - if custom type implementation of a function is broken the interfaced wont help
- interfaces are tough 
    - learn and understand how to read them 
        - use standard library for practice on reading interfaces
- interfaces could be seen as a multiple function holder

