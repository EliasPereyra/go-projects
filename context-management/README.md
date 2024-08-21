# Context

For managing concurrent operations, go routines are a great tool. But for larger operations, handling these go routines are more complex. 

For that, we have the `context` package.

Conext carries **deadlines**, **cancellation signals**, and **request-scoped values** between go routines.