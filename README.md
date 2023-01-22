**To Developers**

This repository is for developers of the SkyLine project or people who plan to work on this language to add new features, keys, types, structures, identifiers etc!

# Documentation 

> Importing files 

Importing files is simple, just import the file and call the keyword from the file required to be imported

```rs
# all of these do the same thing
import("file.csc")
require("file.csc")
include("data.csc")
```

> Carry Files

Carry is a statement like import, but instead of importing the contents into the file it executes the file indivudally before parsing anything else. You use the `carry` keyword followed by a | and the file like so.

```rs
let x = true 

if x { 
   carry|"file.csc"| 
} else {
   carry|"file2.csc"|
} 

```

if X is true only file.csc will be executed else then only file2.csc will be executed. Once done executing those files the interpreter will continue to interpret and parse the rest of the main file 

> I/O

There is only a few standards for input and output with SkyLine currently. There is three ways to output data which are as follows `print`, `varname` and `println` like so.

```rs
allow data = "data"

println(data)
```
or 
```rs
allow data = "data"

print(data)
```
or
```rs
allow data = "data"

data
```

> User Input 

SkyLine has a user input functions which allows you to parse data within a input with the `input` keyword like so.

```rs
allow UI = input("Console>>", "n")
```

the `n` argument tells SkyLine you want to capture input when the user starts a new line or hits the enter key on their keyboard

> Functions 

Functions are simple as well, they work like variables. To init a function you can use keywords `Func` or `function` to specify a function like so.

```rs
let FunctionName = function() {};

let FunctionName2 = Func() {};
```

* Function Arguments 

Function arguments are simple as well, all you do is place the variable name inside of the parenthesis. For multiple variables you can seperate them with `,` or `:` like so.

```rs
let functionname = function(x : y) {
      println(x + y)
}

allow function2 = Func(x, y) {
      println(x + y)
}
```

* Returning 

Returning data in a function is easy as well with two keywords `ret` short for return and `return` for return or you can choose to opt out of a keyword and just return a variable result similar to Rscript like so.

```
allow X = function(x : y : z : w) {
      ret x * y - w / z
} 

#or use the return keyword 

allow X = function(x : y : z : w) {
      return x * y - w / z
} 

# or opt out of the return 

allow X = function(x : y : z : w) {
      x * y - w / z
} 
```

> Arrays

Arrays are very simple and can have multiply data types in SkyLine for example a simple integer arrays looks like this.

```rs
allow arr = [1, 2, 3, 4]
println(arr[1])
```

* Functions in arrays 

functions can work within arrays as long as you call them properly, for example if you want to work with a function and call it from the array with function arguments you will need to add () at the end of the array like so.

```rs
allow XofFunctions = [function(x, y){ x - y} ]
XofFunctions[0](110, 20)
```
