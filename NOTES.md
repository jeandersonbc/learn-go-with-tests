# Notes

## [hello](./hello/)

- Projects are organized as follows:
  - Modules _composed_ packages that are _defined_ by one or more source files
  - The `main` package is the entry-point for the project which might have a `main` function
  - Test usually lives nearby the corresponding implementation: `foo.go` and `foo_test.go`
- On visibility and scope:
  - Functions are verified at package level _not_ by source code, e.g., even if you have `a.go`
    and `b.go`, their definitions share the same scope _iff_ they belong to the same package.
- On testing:
  - The standard lib offers a `testing` package with lots of built-in utilities out-of-the-box
  - `testing.T` is the common type associated to test execution
  - Go also provides the concept of _Examples_ useful for compile-checked code snippets for
    documentation purposes
- Some syntax and conventions covered:
  - main, constants, declarions (and the shorthand format), if-else, switch
  - fuctions and constants are visible to other packages when they start with uppercase letter
  - `TestXxx(...)` indicates test case
  - `ExampleXxx(...)` indicates a compile-checked code snippet for documentation
