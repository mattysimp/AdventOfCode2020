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
First Benchmark is for both parts with standard input. Second Benchmark is for both test with large input (1,000,000 lines)
```
BenchmarkParts-8                     930           1306451 ns/op
BenchmarkPartsStress-8                 1        1404996500 ns/op
```

## Day 3
### Solution
Read input file into a 2d int array of 1s and 0s where 1 denotes '#'. 
#### Part 1
Function part1 loops down list going down and right the amount specified each iteration. A modulus operation is used to keep the right in range, this simulates the repeating pattern. All inputs hit on the 2d are summed for output.
#### Part 2
Creates arrray of the different slopes, runs each one of these through the part1 function asyncronously and then uses channels to send the results to be multiplied together.
### Benchmarks
```
BenchmarkParts-8                    3076            379064 ns/op
BenchmarkPartsStress-8              2925            419485 ns/op
```

## Day 4
### Solution
Read input file send to workers at every empty line, workers convert to passport record. Count non blank passport record 
#### Part 1
Uses Set function that just checks for non blanks when creating passport.
#### Part 2
Uses SetWithVal function that checks for non blanks and validates entries with a mixture of regex and logic when creating passport.
### Benchmarks
```
BenchmarkParts-8                     573           2038401 ns/op
BenchmarkPartsStress-8                 1        1663997800 ns/op
```

## Day 5
### Solution
Read input lines asyncronously and send them to worker. Worker converts the letters into binary and then into an integer for both row and column. 
#### Part 1
ID of each input is compared to a max id. If its greater max id is replaced
#### Part 2
IDs are stored in map, all possible ids are looped over. If ID is not in map but both ID+1 and ID-1 are then it is our seat
### Benchmarks
```
BenchmarkParts-8             504           2152936 ns/op
```

## Day 6
### Solution
Read input file sends content and number of passengers in a group to workers at every empty line
#### Part 1
Worker calculates unique answers in content
#### Part 2
Worker calculates number of answers that have the same count as number of passengers in group
### Benchmarks
```
BenchmarkParts-8                    1262           1048497 ns/op
BenchmarkPartsStress-8                 3         478022600 ns/op
```

## Day7
### Solution
Read input file sends contents asynchronously to a single processor which creates a tree of bags with a hash table of bag pointers to find the starting point in the tree. Has to wait until all data is read to start traversing the tree
#### Part 1
Traverses the tree upwards concurrently using hash table to keep track of total unique colours hit while traversing 
#### Part 2
Traverses the tree downwords concurrently sending number of bags needed to a reciever which totals up
### Benchmarks
```
BenchmarkParts-8             115          10226053 ns/op
```