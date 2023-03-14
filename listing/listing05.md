# listing05

Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

Ответ:

```
Вывод:
error

Первый nil является не просто nil, а nil от customError, второй же nil в if
является nil от nil. При прямом сравнении – они отличаются
```