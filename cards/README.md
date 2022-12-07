### Notes from course

Variables
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

