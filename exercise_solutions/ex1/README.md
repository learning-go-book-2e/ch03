# Exercise 1

## Question
Write a program that defines a variable named `greetings` of type slice of strings with the following values: `"Hello"`, `"Hola"`, `"नमस्कार"`, `"こんにちは"`, and `"Привіт"`. Create a sub-slice containing the first two values, a second subslice with the second, third, and fourth values, and a third subslice with the fourth and fifth values. Print out all four slices.

## Solution
This problem tests creating a slice literal and calculating slice offsets. The first value in a slice expression is the zero-based starting position. It's optional if you are starting at the beginning of the slice. The second value is one greater than the last position you want to include. The second value is optional if you want to go all the way to the end of the slice. For our sample, we want to create the following three subslices:

```go
	slice1 := greetings[:2]
	slice2 := greetings[1:4]
	slice3 := greetings[3:]
```
