# Go Training

## Syntax
- Use the short variable declaration operator to initialize a var to some value `myString := "Blah"`
- Use `var` to initalize a variable to its zero value `var strings [5]string`

## Interface satisfaction

Interface satisfaction is at the value level based on the rules for method sets:
You can't call an interface method that is implemented with a pointer receiver by passing an object by value. The object might not exist in memory. Pass a pointer instead. If the inteface method is implemented with a value receiver, you may pass a value or pointer.

## Package oriented design
Name packages according to what they provide.

BAD
___
common
util
etc

GOOD
___
publish_queue
asset_validation_api
etag_crawler


- Packages are larger than a class or library, but smaller than an entire project. 
- They should provide a semantic unit of functionality.
- Package should have its own dir.
- Packages included with `package package_name`
- "Main" file in a package should be named like package_name.go 

