def genSeries(n):
  
    a = [0] * n
    a[0] = 0
    a[1] = 1
 
    for i in range(2, n): 
        a[i] = a[i - 2] + a[i - 1]
      
    for i in range(n - 1, -1 , -1): 
        length = len(a)
        totLen = len(str(a[length-1])*a[length-1])
        if a[i]==0:
            xcount = totLen-1
            print(str(0)+(xcount*'x'), end=" \n")
        else:
            temp = str(a[i])*a[i]            
            xcount = totLen-len(temp)
            print(temp+(xcount*'x'), end=" \n")
    
nterms = int(input("enter n"))
genSeries(nterms)
