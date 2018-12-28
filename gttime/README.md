# Go Toolkit Time  
Simple helper functions to deal with go time

## Milliseconds
Returns the timestamp in miliseconds of a given Time

``` go
t := time.Now()
ms := gttime.Milliseconds(t)

fmt.Printf(ms)
```

## ElapsedMilliseconds
Returns the difference, in miliseconds, between two Time(s)

``` go
t1 := time.Now()
t2 := time.Now()

diff := gttime.ElaspedMilliseconds(t1, t1)

fmt.Printf(diff)
```