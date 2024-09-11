# Resource Management with sync.Pool

A pool is a set of temporary objects. It's like a recycling bin for restoring objects for later reuse of them.
This is important because it reduces the time and performance cost of creating and deleting objects.

It is also automatically garbage collected, and it's safe among concurrent operations with go funcs.

The pool is benefitial for creating objects of short-live times. It's not suited for long operations or state.

## Use Cases

- For handling http responses with buffers
- For logging system in concurrent operations
- DB connections management.
