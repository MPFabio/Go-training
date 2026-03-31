package twofer

import "fmt"

func ShareWith(name string) string {
    allowed := map[string]bool{
    "Alice": true,
    "Bohdan": true,
    "Bob": true,   
    "Zaphod": true,
}
	
	if allowed[name] {
  		return fmt.Sprintf("One for %s, one for me.", name)
	} else {
        return "One for you, one for me."
    } 
}
