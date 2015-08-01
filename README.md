stdoutlog
---------

Basic stdout logging middleware.

Modified version of the code found in [Negroni's logging middleware](https://github.com/codegangsta/negroni/blob/master/logger.go#L21), but also shows more details on the request.

# Usage with Negroni

```
import (
    "github.com/codegangsta/negroni"
    "github.com/jeffbmartinez/stdoutlog"
)

n := negroni.New(…)
n.Use(stdoutlog.Middleware{})
// ...
```
