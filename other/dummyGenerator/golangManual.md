# General

`fmt` - Package allowing formating and printing of values 

`fmt.Sprintf(format, variables)` - Cast variables to a format and returns a string 
- Example: `fmt.Sprintf("Hello, %s", "Mr. Weyer") = "Hello, Mr. Weyer"`

`foo := 0` - weak type definition of a variable 

`var foo int = 0` - strong type definition of a variable 

# Variables 

## The type of the variable is always beind it's name 
<code>
func baz(name string, age int) {} <br>
var taz int = 0 <br>
func sum(a, b int) int <br>
</code>


`func foo() (int, string)` - foo returns multiple variables -> an **integer** and a **string** at the same time

`bar, baz := foo()` - define both variables that the function `foo()` returns

If a value of the variable is changed, say variable is redefined, it's defined with a single `=`
<code>
<br>
foo := 0 <br>
...<br>
foo = 49<br>
</code>

# Operations 

### foreach loop 
<code>
for [index], [variable] := range [array] { <br>
    ... <br>
}
</code>


# Terminology

`slice` = array, list