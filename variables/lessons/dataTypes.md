# Data Types - Basic Numeric Types

- 15 different numeric types in Go
- 3 broad categories: ```int```, ```float```, ```complex```
- 11 diff ```int```'s, 2 ```float```'s and 2 ```complex``` 
- ```bool```, ```int8```, ```int16``` and so on as in many other languages, and unsigned versions too
- Again, as in other languages, the number at the end represents the number of bits used to store them (and ```bool``` is one bit)
- Good practice to store a variable in the right data type so as to optimise performance - a variable which is true or false is much better stored in a ```bool``` than in a ```uint64```

## Floats and Complex Numbers

- 2 types: ```float32``` and ```float64``` (or ```double``` in some languages)
- These are specified by the *IEEE 754 Standard for Floating-Point Arithmetic*, created in 1985 to make float implementations reliable and portable
- See spec below for more details (quite technical, but overall ```float32``` means 6-9 significant figures, and ```float64``` is 15-17)
- ```complex64``` and ```complex128``` are complex numbers with ```float32``` and ```float64``` real and imaginary parts, respectively

### References

- [Golang numeric type docs](https://go.dev/ref/spec#Numeric_types)
- [IEEE 754 specification](https://en.wikipedia.org/wiki/IEEE_754)