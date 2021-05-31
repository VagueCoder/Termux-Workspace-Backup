class Insertion:
    def __init__(self,*list):
        self.list=list 
        self.size=len(list)
        print(self.list, self.size)
    
    def InsertionSort(self)->list[int]:
        for i in range(1,self.size):
            key=self.list[i]
            j=i-1
            while(j>=0 and self.list[j]>key):
                self.list[j+1]=self.list[j]
                j=j-1
            self.list[j+1]=key
            return self.list
        
X=Insertion([23,43,34,5,6])
print(X.InsertionSort())
