** To Developers **

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

> Carry 

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
