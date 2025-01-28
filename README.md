# Go Defer Statement Gotcha

This example demonstrates a common pitfall when using Go's `defer` statement: the values captured by the deferred function are evaluated at the time the `defer` statement is encountered, not at the time the deferred function is executed.

## The Bug

The `main` function initializes an integer `x` to 10. A `defer` statement is used to print the value of `x` after the surrounding function completes. However, before the deferred function executes, the value of `x` is changed to 5.

One might expect the deferred function to print 5, but it actually prints 10, because the value of `x` that's used in the closure is evaluated at the time of the `defer` call (when `x` was 10). 

## The Solution

The solution is to either capture the current value of `x` explicitly using a temporary variable, or restructure the code to avoid this issue.