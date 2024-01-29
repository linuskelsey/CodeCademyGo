# 1. Numeric Types

- 15 different numeric types in Go
- 3 broad categories: `int`, `float`, `complex`
- 11 diff `int`'s, 2 `float`'s and 2 `complex` 
- `bool`, `int8`, `int16` and so on as in many other languages, and unsigned versions too
- Again, as in other languages, the number at the end represents the number of bits used to store them (and `bool` is one bit)
- Good practice to store a variable in the right data type so as to optimise performance - a variable which is true or false is much better stored in a `bool` than in a `uint64`

### Floats and Complex Numbers

- 2 types: `float32` and `float64` (or `double` in some languages)
- These are specified by the *IEEE 754 Standard for Floating-Point Arithmetic*, created in 1985 to make float implementations reliable and portable
- See spec below for more details (quite technical, but overall `float32` means 6-9 significant figures, and `float64` is 15-17)
- `complex64` and `complex128` are complex numbers with `float32` and `float64` real and imaginary parts, respectively

# 2. Strings

- Can declare `string`'s in the same way as numbers, but when assigning we need to use double-quotation marks
- We can concatenate in the normal way with the `+` symbol

# 3. Syntax and Defaults

- Declare and assign by: 

    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code style="color : lightgray">var</code><code style="color : lightblue">varName</code><code style="color : lightpink">varType</code>

    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code style="color : lightblue">varName</code><code style="color : lightgray">=</code><code style="color : lightsalmon">varValue</code>

    or,

    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code style="color : lightgray">var</code><code style="color : lightblue">varName</code><code style="color : lightpink">varType</code><code style="color : lightgray">=</code><code style="color : lightsalmon">varValue</code>

- Even before assignment, variables hold default values. That might be the empty string `""`, it might be `0`, it might be `false` - sensible defaults

# 4. Type Inferrence

- It is possible to declare without stating type, via type inferrence
- The `:=` operator handles this for us, for example `isAboveFiveFoot := true` would declare and assign a variable `isAboveFiveFoot` of value `true` and inferred type `bool`
- We can also declare and assign a variable in the normal way, just ommitting the <code style="color : lightpink">varType</code> tag from the above section
- Floats created in this way are automatically `float64`'s and ints are either `int32`'s or `int64`'s
- The specific int created by type inferrence of integers depends on the computer architecture (be it 32- or 64-bit ROM)

### References

- [Golang type docs](https://go.dev/ref/spec#Types)
- [IEEE 754 specification](https://en.wikipedia.org/wiki/IEEE_754)