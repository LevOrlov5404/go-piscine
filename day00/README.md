We have a bunch of integer numbers.
We should to know statistical metrics for it.

#Input:
* arguments for metrics to calculate
* integer numbers (-100000 < number < 100000) from a standard input, separated by a new line.

#Output on valid data:
* metrics that given in args (by default all four metrics).
```
Mean: 8.20
Median: 9.00
Mode: 3
SD: 4.35
```
1. Mean is an average.
2. Median is a middle number of a sorted sequence if its size is odd, and an average between two middle ones if the size is even.
3. Mode is a number which is occurring most frequently (if there are several, the smallest one among those).
4. SD is a standard deviation.

#Output on invalid data:
* error with description

#build:
```go build day00.go```

#run:
```cat text1.txt | ./day00 -sd -mode -median -mean```
