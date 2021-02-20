package main 
  
import ( 
    "fmt"
    "strings"
) 
package main 
  
import ( 
    "fmt"
    "strings"
) 

func main() { 
  
    // Creating a slice of strings 
    slc := []string{"tall", "taller", "tallest"} 
  
    // Convert the case of the 
    // elements of the given slice 
    // Using ToUpper() function 
    for x := 0; x < len(slc); x++ { 
  
        // Exported Method 
        res := strings.ToUpper(slc[x]) 
  
        // Exported Method 
        fmt.Println(res) 
    } 
} 