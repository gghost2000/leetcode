func isAlienSorted(words []string, order string) bool {
    m := make(map[rune]int)
    for i, r := range order {m[r]=i}
    
  
    indexOf := func(a rune) int {
        idx := -1
       
        if i, ok := m[a]; ok {idx = i}
        return idx
    }
 
    inOrder := func(a,b rune) bool {
        ia, ib := indexOf(a), indexOf(b)
        return ia <= ib
    }
    
 
    areOrdered := true
    w0 := words[0]

 
    for _, wi := range words[1:] {
        rw0, rwi := []rune(w0), []rune(wi)
        
     
        for j := 0; j < len(rw0); j++ {
            c0, ci := rw0[j], ' '
           
            if j < len(rwi) { ci = rwi[j] }
          
            if c0 != ci {
                areOrdered = inOrder(c0,ci)
                break
            } else {
              
                continue
            }
        }
      
        if !areOrdered {break}
        
 
        w0 = wi
    }
    
    return areOrdered
}