# Switches

A `switch` statement is an alternative (and idiomatic) syntax for displaying if-else if-else chains. Here is an example of a switch:

```
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

The function `unhex` is used to convert a `byte`, or an `int8` in binary, in hexadecimal form to an integer in base-10. Let's break it down:

- The `switch` keyword intitiates the statement (this is sometimes followed by a 'value', the variable/constant that we are comparing to a number of different reference points in the **cases** that follow)
- Inside the `switch` block, `{ ... }`, we have multiple `case` statements making certain comparison checks, and if we have a value following `switch`, these checks will simply comprise of another value, and check if the two are equal
- Once a condition is met, the code inside that `case` is executed, and the `switch` statement is exited
- At the end of the switch statement we may also have a default statement, which runs if no other `case` is satisfied

## Scoped Short Declaration

This is a short variable declaration within a `switch` statement (it can also be used in if-else if-else chains). The syntax is as follows:

```
switch season := "summer" ; season {
case "summer":
  fmt.Println("Go out and enjoy the sun!")
}
```

Here we declare the varriable `season` within the `switch` statement to begin with, and then use the `switch` syntax wherein we are checking equality. An advantage of this is that now we ave 'scoped' the declaration of `season`, and so it is only accessible within the `switch`, thus saving on memory.

### References

- [Golang switch docs](https://go.dev/doc/effective_go#switch)