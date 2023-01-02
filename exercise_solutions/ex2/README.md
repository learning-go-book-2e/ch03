# Exercise 2

## Question
Write a program that defines a string variable called `message` with the value `"Hi 👩 and 👨"` and prints the 4th rune in it as a character, not a number.

## Solution
This question covers type conversions with strings and runes, as well as the zero-based offset for slices. You need to convert the `message` string to a `[]rune` using a type conversion. Then you get the 4th rune in the rune slice using `runes[3]`. Finally, you need to use a `string` type conversion to convert the rune back into a character. 