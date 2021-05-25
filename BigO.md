# BigO Notation

Big O notation is a mathematical notation that describes the limiting behavior of a function when the argument tends towards a particular value or infinity.

Insimple words BigO Notation helps complexity of code to be expressed in algebric terms. There are two basic categories 

1. Time Complexity
2. Space Complexity

# Time Complexity

## BigO Complexity Chart

Inorder to measure the complexity of algorithms there is bigO complexity chart . Its shown below. This chart effectively represent how our algoritm performs when input grows bigger and bigger.
![Big O Notation](./Assets/BigONotation.jpeg)

## BigONotations

```
 - Constant Time        O(1)                No Loops
 - Logarithmic          O(log N)            Usually Searching algorithms log(n)
 - Linear Time          O(n)                Single Loop
 - Log Linear           O( nlog(n) )        Sorting operations usually
 - Quadratic Time       O(n^2)              Two nested loop
 - Exponential          O(2^n)              Recursive algorithms that solve problem of size n
 - O(n!)                Looping for every element.
```

## Calculating Big-O Notation.
Big-O Notation can be found using the following way. Consider the function below.

```py
def find_big_o(input_array):
    flag = True                 # O(1)
    CALL_COUNT += 1             # O(1)
    for item in input_array:    # O(n)
        another_function()      # O(n)
        print(item)             # O(n)
    return input_array          # O(1)
```

From the function above, we can see that  first line **flag = True** will be called
only once, similarly do **CALL_COUNT** constant. so both have bigO of constant time **O(1)**

Now conside the loop It will run for the length of user input so it will have a
bigO of linear time **O(n)**, While calculating bigO We don't need to be aware of the 
nitty gritty of **another_function** from the loop point of view we can say that it
will be called for user input length,so it will have bigO of linear time. The last **return** statement will also be called once. So it will have contant time. So BigO 
of the whole function will be **O(1) + O(1) + O(n) + O(n) + O(n) + O(1)** which will 
be equal to **O(3+3n)**. But since BigO only measures the upper bound the constant 
in this function will be of less significance. So what we can say from these is 
that our function will be of BigO **O(n)** dropping both **3** from the equation.

## Rules for finding BigO

There are mainly 4 rules in finding the BigO 
 1. Worst Case
 2. Remove Constants
 3. Different terms for inputs
 4. Drop Non Dominants

### 1. Worst Case

Consider the below snippet of code

```py
fruits_array = ['grapes', 'banana', 'apple', 'orange', 'guava']

for fruit in fruits_array:
    if fruit == 'apple':
        print("apple found")
        break
```

So if we look at the above snippet of code we can see that we are looping through
the array to find the apple. Even though we know that our apple can be found on 
the third iteration, bigO notation is only aware of worst case scenario here. i.e, 
apple found at the end of the list. So if apple is found at the end of the list we
can say that the iteration needs to be done through the entire length of the array.
So the bigO of the above function will be *O(n)*

### 2. Remove Constants
BigO is the upper bound limit we can say that the constants in the function will
have little effect while considering about variables. Which means suppose while 
computing the bigO of a function we ended up something like this **O(n/2 + 10,000)**
Even though 10,000 is a large term for small input we can say when input grows 
significantly say around 1 crore or more that 10,000 will have less significance
in specifiying execution speed of that function and thus can be omitted. The same 
applies to the **n/2** part as the input grows significantly we can round **n/2** to **n**.
Thus we can say that the BigO notation for the above experssion can be reduced to **O(n)**

### 3. Different terms for inputs
Consider the below code snippet

```py
def calculate_big_O(input1, input2):
    for element in input1:
        print(element)
    
    for element in input2:
        print(element)
```
So what you think the BigO of the function will be? We will be tempted to say **O(n)**
from our past rules. But not the difference here, We have two different inputs here
so what if the two inputs are of different lengths? This needs to be accounted while
calculating bigO notation. So the bigO of above function will be **O(a+b)** instead
of **O(n)**

### 4. Drop Non Dominants
Consider the below snippet of code.What do you think the bigO will be?

```py
three_by_three_mat = [[0,0,0], [0,1,0], [0,0,0]]

def print_all_elements(three_by_three_mat):
    for row in three_by_three_mat:
        for item in row:
            print(item)
    
print_all_elements(three_by_three_mat)
```

The bigO of the above code snippet is **O(n) * O(n)** which effectively turns out
to be  **O(n<sup>2</sup>)** So each increase of nested loop increase the power its
being raised. 

So consider we ended up with a function having bigO notation as **O(n<sup>2</sup> + 4n + 10,000)**
from the above rules we can be sure that the *10,000* part can be omitted so what 
is left to reduce is **O(n<sup>2</sup>+4n)** here comes the 4<sup>th</sup> rule.
Drop non dominants.  So what we can do is comparing n with n^2 we know that n will 
be less significant compared to n^2 so the dominant term here is n^2 which is of 
higher order compared to n. So the non dominant part n can be omitted which left 
us with BigONotaion **O(n<sup>2</sup>)**.



# Space Complexity

The space complexity of an algorithm or a computer program is the amount of memory space required to solve an instance of the computational problem as a function of characteristics of the input.

## What causes space complexity?

Space complexity if affected mostly by the following:

- Datastructures
- Variables
- Function Calls
- Allocations

Consider the below function

```py
def printIntegers(input_array):
    for i in range(10):
        print(i)
```
What do you think its space complexity is? Well we can say that it takes in variable i so it definitely consumes some space. Also is same for the input. BUt here in the input we know that its not actually in our control, it can grow and it doesnot depends on function anyhow. So we can say that the space complexity of the above function is O(1) which is constant.

Now consider the below function

```py
def add_to_array(input_array):
    local_array = []
    for item in input_array:
        local_array.append(item)
```
Here we are actually adding contents based on input_array length to our internal array local_array. So the space complexity of the above function will be O(n)