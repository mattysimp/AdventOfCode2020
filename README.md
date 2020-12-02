# Advent Of Code 2020

## Day 1
### Solution
#### Part 1
Input is iterated over. Each x in input is looked up in Hash Table for Y in `y = 2020 - x` and added to Hash table if result is not found
#### Part 2
My method was admittedly not great for this. To find i, j, k so that `i + j + k = 2020`, I first sorted the data then, brute forced i and j and then used a more efficient binary search to find if there is a k.
### Benchmarks
```
BenchmarkPart1-8            8134            189084 ns/op
BenchmarkParts2-8           8001            150981 ns/op
```

## Day 2
### Solution
Firstly I concurrently parsed in the input file and send it to a worker. The worker calculates whether the input hits the criteria and if so passes it to a reciever. This reciever counts all correct inputs. Once all inputs are read and gone through the worker, all channels are closed and the code completes.
#### Part 1 Worker
Uses strings.Count to work out the number of specified letter and checks againt the min and max inputs
#### Part 2 Worker
Checks for correct letter in both positions and uses an Xor to check if correct
### Benchmarks
Runs in 1.2ms for 1000 input
Runs in 1.4 seconds for 1,000,000 input