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

## Day 7
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

## Day 8
### Solution
Read input file into slice. Run through game console slice following the nop, acc and jmp instructions
#### Part 1
Keep hash table of all input line ran. If hash mapping exists for input line then it has already been ran, return accumulator
#### Part 2
Loop and brute force edits to nop and jmp instructions. Flag if game console booter ends through the last line or through the hash repeat. Asynchronously run each edited input until one flags as booted.
### Benchmarks
```
BenchmarkPart1-8            6081            336293 ns/op
BenchmarkPart2-8             170           7111766 ns/op
```

## Day 9
### Solution
#### Part 1
Loop through slice and send each 26 element slice to worker. In worker, loop through each slice creating a hash table used to find the sum of 2 that adds to the final element. If it doesn't exist then return the final element.
#### Part 2
Loop through slice and send each set of element slice to worker (Slice[1:...], Slice[2:...], etc). In worker, loop through each slice to try and have a combination which adds to the answer to part 1, recording the largest and smallest numbers. Send any valid answer to reciever with the posiiton in the input array. Answer with the lowest input position is correct,
### Benchmarks
```
BenchmarkPart1-8                    1188           1147305 ns/op
BenchmarkPart2WithPart1-8           1999            651826 ns/op
```

## Day 10
### Solution
Read lines in, sort and convert to XMAS record used in Day 9. 
#### Part 1
Loop through adding differences between entries in a hash map. Multiple the 1 differences by the 3 differences
#### Part 2
Recusrivley traverse the slice by getting a list of valid next inputs and recursing for each valid next input adding together the number of following combinations from that point. These are added to a hash table to memoize the data to speed up later reads.
### Benchmarks
```
BenchmarkPart1-8                   12817             84965 ns/op
BenchmarkPart2WithPart1-8          13969             86835 ns/op
```

## Day 11
### Solution
Read data into a hash map of key image.Point and val rune. Recursively run through map checking for changes to seats. Write new positions to a new hash map. If nothing has changed then return count od occupied seats
#### Part 1
Tolerance set as 4. Function for finding adjacent points to check is a simple point.Add
#### Part 2
Tolerance set as 5. Function for finding adjacent loops to find a point that isn't '.', this included null mappings.
### Benchmarks
This one was a really slow solution. I may come back and try and make it more efficient.
```
BenchmarkPart1-8                       2         817506050 ns/op
BenchmarkPart2-8                       1        1887993200 ns/op
```

## Day 12
### Solution
#### Part 1
Loop through input moving ferry as specified.
#### Part 2
Loop through input moving 2 ferry varibles (Ferry and Waypoint) as specified.
### Benchmarks
```
BenchmarkPart1-8                    5658            211627 ns/op
BenchmarkPart2WithPart1-8           4436            275336 ns/op
```

## Day 13
### Solution
#### Part 1
Compare each buses next departure time after the start time.
#### Part 2
Moving through the bus id's find a correct time for the first bus. The step to the next bus can now be the correct time for the last bus. Find the correct time that is correct for both the first bus and the second bus and change the step to that time. Keep repeating until all buses are satisfied.
### Benchmark
```
BenchmarkPart1-8            15304             68270 ns/op
BenchmarkPart2-8            13843             77440 ns/op
```

## Day 14
### Solution
Read input into slice and split. If mask then save to variable, if mem then get address and bits.
#### Part 1
Use two bit operations to replace all the bits with all non X's in mask.
1. Replace all X's with 0's and do an or operation
2. replace all X's with 1's and do an and operation
##### Example
```
value:      000000000000000000000000000000001011
mask:       XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X

or:         000000000000000000000000000001000000
result1:    000000000000000000000000000001001011

and:        111111111111111111111111111111111101
result:     000000000000000000000000000001001001
```
Save result in mem location and loop through to sum after completed all lines.
#### Part 2
Create slice of ints.  Replace all X's with 0's and do an or operation and add to slice. Loop through mask and for every X, duplicate each entry in slice and take i (position in mark) `newAddress = address - (i^2/)2`
Save the result in each memory address. Loop through to sum after completed all lines.
##### Example
```
value:      000000000000000000000000000000101010
mask:       000000000000000000000000000000X1001X

or:         000000000000000000000000000000110011
result:     000000000000000000000000000000111011

afterloop:  000000000000000000000000000000111011
            000000000000000000000000000000011011
            000000000000000000000000000000111010
            000000000000000000000000000000011010
```
### Benchmark
```
BenchmarkPart1-8            1428            810226 ns/op
BenchmarkPart2-8              66          16772714 ns/op
```

## Day 15
### Solution
Loop for target number iterations (p1: 2020, p2: 30,000,000). If `i < len(input)`, return value from input and write to an array. Write the turn number to an array at the address of the current number.If not, if the current numbers address has a non-zero value in the array then the next number is i minus that number. If it is zero, then next number is 0. 

I used a hash map first but swapped to a slice of length 30,000,000 for greater speed efficiency with not having to do hash calculations. This would make the code not usable for greater lengths without changing the static length. The change took part 2 from taking >3s to <1s.
### Benchmark
```
BenchmarkPart1-8              44          27136614 ns/op
BenchmarkPart2-8               2         817040350 ns/op
```