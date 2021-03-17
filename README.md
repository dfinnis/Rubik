# Rubik

A Rubik’s cube solver based on Thistlethwaite's 4 groups using IDA-star path search, written in golang.

The goal is to minimize both solve time and solution length.
On average this project solves randomly mixed cubes in 30 something moves in under 2 seconds.

#### Final Score 124/100

## Getting Started

First you need to have your golang workspace set up on your machine.
Then clone this repo into your go-workspace/src/ folder.

```git clone https://github.com/dfinnis/Rubik.git; cd Rubik```

Download dependencies.

```go get -d ./...```

To run, go run main.go with a mix.

```go run main.go mix/subject.txt```

Alternatively, build and run the binary with a mix.

```go build; ./Rubik mix/subject.txt```


### Usage

<img src="https://github.com/dfinnis/rubik/blob/master/img/usage.png" width="500">


### Notation - 18 valid moves

* 90° clockwise twists of the 6 sides (Up, Down, Right, Left, Front, Back) -> U, D, R, L, F, B

* 90° anti-clockwise twists are denoted with ```'``` -> U', D', R', L', F', B'

* 180° twists are denoted with ```2``` -> U2, D2, R2, L2, F2, B2

Solution length is measured by Half-turn metric (both ```F``` (90° twist) & ```F2``` (180° twist) count as one move).


### Argument String

A sequence of moves can be provided as argument as a string. e.g. "F U".

```go run main.go "F U"```

<img src="https://github.com/dfinnis/rubik/blob/master/img/sequence_string.png" width="800">

### Argument Filepath

Alternatively, give a filepath containing a valid sequence as argument, there are some example mixes in the *mix/* folder.

```go run main.go mix/U_F.txt```

<img src="https://github.com/dfinnis/rubik/blob/master/img/filepath.png" width="800">

### Argument -r --random

Create and run a random mix.
An optional following *len* argument specifies mix length, i.e. *-r 5* will create a random 5 move sequence.

```go run main.go -r```

<img src="https://github.com/dfinnis/rubik/blob/master/img/--random.png" width="800">


## Flags

### Visualizer

*-v* (*--visualizer*) shows visual of mix and solution. Here is a simple example with a random mix of length 5:

```go run main.go -r 5 -v```

![Visualizer](https://github.com/dfinnis/rubik/blob/master/img/visualizer.gif)

We start with a solved cube, it shows the mix then spins once to show the mixed cube state.
Then it shows the solution, finally spining twice to show the cube is back in its solved state.


### Group

*-g* (*--group*) shows solution breakdown by subgroup (see [Cube representation](#cube-representation) & [Thistlethwaite's 4 groups](#thistlethwaites-4-groups) for explanation).

```go run main.go -r -g```

<img src="https://github.com/dfinnis/rubik/blob/master/img/group.png" width="640">

We start at the top with the mixed cube, described by corner and edge permutation and orientation.

We then see the solution broken down into Thistlethwaite's 4 subgroups.
For each subgroup we see the solution, half-turn metric, and solve time. Then the state of the cube after applying the subgroup solution, highlighting elements solved in the subgroup.

At the bottom we should arrive at subgroup 4, a solved cube.
The orientation is all correct (0), and each corner and edge permutation is in its right place (e.g. edge 1 is in permutation 1).

## Tests

The test script runs 10 static random unit tests, followed by 10 dynamic random tests.
It then displays best, worst and mean for Half-turn metric and solve time.

Finally it runs some unit tests from the *mix/* folder, to make sure it deals with edge cases, and the mightily hard superflip.

```./test.sh```

<img src="https://github.com/dfinnis/rubik/blob/master/img/test.png" width="580">


## Cube representation

The sensible way to represent the cube for this style of solution is corner and edge permutation and orientation.

A cube is made up of 26 cubelets. 6 of these are center cubies (one for each face) which cannot move, so we do not need to include these in our model.

There are 12 edge cubies, these have 2 colors, rotate on 2 axes, and so can be orientated correctly (0) or incorrectly (1).

There are 8 corner cubies, these have 3 colors, rotate on 3 axes, and so can be orientated correctly (0) or incorrectly (1) or (2).

For a solved cube (Thistlethwaite's group 4) all the orientations are 0 (correct) and each corner and edge is in its correct permutation (e.g. edge 1 is in permutation 1):

<img src="https://github.com/dfinnis/rubik/blob/master/img/Solved_cube.png" width="640">

Here is an image to clarify cubie notation:

<img src="https://github.com/dfinnis/rubik/blob/master/img/Rubik_notation.png" width="640">

Check out [this demo](https://iamthecu.be/) for a visual clarification.


## Thistlethwaite's 4 groups

If we wanted to find the shortest solution to a mixed cube, we could simply try all possible move combinations until it is solved.
Being that any cube can be solved in 20 moves, and there are only 18 different moves, a stupid brute force solution would mean trying 18 to the power of 20 moves.
This is not computationally viable.

Thistlethwaite's algorithm breaks down the problem of solving the cube into 4 groups.
This means a much smaller space has to be searched than trying to find the whole solution in one go.
In each group only some aspects of the cube are solved, and only certain moves are allowed.
As we progress through the groups we reduce the moves available, starting with all moves (group 0), and finishing with only 180° moves (group 3).

In addition, pre-computed pruning tables allow us to look up how many moves until we reach the next group.
This means we can try all possible moves for a given cube, look in the pruning table if we are closer to reaching the next group, and simply follow the shortest path.

#### group 0

First only the edge orientation is solved. All moves are allowed.

There are 12 edges, each can be oriented correctly (0) or incorrectly (1).
So there are only 2 to the power of 11 (2048) possible edge orientation combinations.

We can pre-compute a pruning table of how many moves until we reach group 1.
Start with a solved cube, apply all 18 possible moves.
For each new cube we can now record the edge orientation combination is 1 move away from being solved.
This process continues, applying all possible moves and filling out the pruning table until full.
We now have a complete pruning table which associates edge orientation combination with how many moves until group 1 (edge orientation solved).
We reach all possible edge orientation combinations within 7 moves, i.e. maximum 7 moves to group 1.

Hopefully now it becomes clear why it would be unreasonable for us to simply create a pruning table for all 43 quintillion possible cube states.
With Thistlethwaite's groups we can create a much smaller pruning table for this small group.
While this pre-computation can take a few minutes, with the pruning table an IDA* search can very quickly find the shortest path to group 1.

#### group 1

Only ```U``` and ```D``` moves effect edge orientation (```U2``` and ```D2``` do not effect edge orientation).
So in group 1 now the edge orientation is solved, we remove ```U``` and ```D``` moves from our allowed moves. Thus the solved edge orientation is locked.

In group 1 we fix the orientation of the corners, and place the middle layer edges into their slice (edges 8 - 11 in permutation 8 - 11).

There are 1,082,565 possible combinations, and we reach all possible combinations after 10 moves.
We can create a pruning table for these combinations in a similar process to group 0, associating each corner orientation and middle layer edge permutation with how many moves until group 2.

#### group 2

Only ```F``` and ```B``` moves effect corner orientation and can move middle layer edges from their slice (```F2``` and ```B2``` do not effect this).
So in group 2 we remove ```F``` and ```B``` moves from our allowed moves. Thus the solved corner orientation is locked, and the middle layer edges cannot leave their slice.

In group 2 edges in the L and R faces are placed in their correct slices, the corners are put into their correct tetrads, the parity of the edge permutation (and hence the corners too) is made even, and the total twist of each tetrad is fixed.

There are 2,822,400 possible combinations, and we reach all possible combinations after 13 moves.

#### group 3

In group 3 we are reduced to only 180° moves (U2, D2, R2, L2, F2, B2).

There are 663,552 possible combinations, and we reach all possible combinations after 15 moves.

#### group 4

Solved cube! Only 1 possible combination!


### IDA-star search

We use iterative deepening A-star search to find the shortest path through each group. It uses the pruning table distance to next group as heuristic.

IDA* concentrates on exploring the most promising nodes and thus does not go to the same depth everywhere in the search tree.

Unlike A* search, IDA* does not remember where it has already searched to avoid repeating itself. This is obviously more suited to Rubik's cubes with 43 quintillion possible states. However because of this, unlike A* search, IDA* may end up exploring the same nodes many times.

### Integrated pruning table maker

When this project is launched, it attempts to read the pre-computed pruning tables from the ```tables/``` folder. If these files have been deleted, it creates them (can take a few minutes).


## Alternative solutions

### Human solver - short solve time, long solutions

The [beginners method for solving cubes](https://ruwix.com/the-rubiks-cube/how-to-solve-the-rubiks-cube-beginners-method/) involves progressively solving parts of the cube from the top down. There are specific sequences of moves which will fix parts of the cube at each step. Once you have memorized these sequences then you can use this method to solve any cube! You could write a program to emulate this method, the solve time is almost instant, but produces solutions more than 100 moves long.

### Kociemba's algorithm - medium solve time, shorter solutions

Herbert Kociemba improved Thistlethwaite's algorithm by reducing the number of groups from 4 to 2. Each group is a much larger search space, requiring larger pruning tables, but the resulting path length is shorter. Thistlethwaite should solve any cube in max 45 moves, Kociemba reduces this to max 30 moves.

### Korf's algorithm - long solve time, shortest solutions

Thistlethwaite and Kociemba's solutions will find solutions quickly, however they will rarely be the shortest solution.
Korf created an optimal solution, one that finds the shortest path, but may take a long time to calculate.
Korf broke down the cube into subproblems: corners, 6 edges, and the other 6 edges.
Korf uses IDA* search to explore all options, and eliminate options which are not optimal.
All possible cube states can be solved in 20 moves maximum (God's number).


## Dependencies

Thankfully, running ```go get -d ./...``` should take care of all dependencies for you.

robotgo -> to type moves into the visualizer website.


## References

[Rubik's cube explorer - visualizer](https://iamthecu.be/)

[Wikipedia Optimal_solutions_for_Rubik](https://en.wikipedia.org/wiki/Optimal_solutions_for_Rubik%27s_Cube#Kociemba's%20algorithm)

[Jaap's puzzle page](https://www.jaapsch.net/puzzles/cube3.htm)

[Stanford computer cubing](https://cube.stanford.edu/class/files/rokicki_cubecomp.pdf)

[IDA-star algo in general](https://algorithmsinsight.wordpress.com/graph-theory-2/ida-star-algorithm-in-general/)

[Pruning table indexing functions](http://joren.ralphdesign.nl/projects/rubiks_cube/cube.pdf)
