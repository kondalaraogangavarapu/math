package add

import testing

func TestAdd(t *testing.T){
 
 total := Add(5, 5)
 if total != 10 {
   t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
 }
 
 total = Add(0,0)
 if total != 0 {
   t.Errorf("Sum was incorrect, got: %d, want: %d", total, 0)
 }
 
 total = Add(132343843843484332847384, 0)
 if total != 132343843843484332847384 {
    t.Errorf("Sum was incorrect, got: %d, want: %d", total, 132343843843484332847384)
 }

}
