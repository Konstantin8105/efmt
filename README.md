# efmt
golang sprintf in engineer format
package efmt // import "github.com/Konstantin8105/efmt"


FUNCTIONS

// View float64 in engineer format
func Sprint(v float64) string


Example of result

```
        -123.457-006        -1.234568e-04
         1234.57-006         1.234568e-03
        -1234.57-006        -1.234568e-03
         12.3457-003         1.234568e-02
        -12.3457-003        -1.234568e-02
         123.457-003         1.234568e-01
        -123.457-003        -1.234568e-01
             1.23457         1.234568e+00
            -1.23457        -1.234568e+00
             12.3457         1.234568e+01
            -12.3457        -1.234568e+01
             123.457         1.234568e+02
            -123.457        -1.234568e+02
         1.23457+003         1.234568e+03
        -1.23457+003        -1.234568e+03
         12.3457+003         1.234568e+04
        -12.3457+003        -1.234568e+04
         123.457+003         1.234568e+05
        -123.457+003        -1.234568e+05
         1.23457+006         1.234568e+06
        -1.23457+006        -1.234568e+06
```
