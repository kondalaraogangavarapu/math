package math

import "testing"

func TestDiv(t *testing.T) {

    total := Div(40, 4)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }

    total := Div(40, 100)
    if total != 0.4 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
    
    total := Div(40, 0)

}
