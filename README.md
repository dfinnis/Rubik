# Rubik

A Rubik’s cube solver based on Thistlethwaite's 4 groups using IDA-star path search, written in golang.

The general aim is to minimize both solve time and solution length, as measured by Half-turn metric.
Half-turn metric counts both ```F``` (90° turn of front face) and ```F2``` (180° turn of front face) as one move.

On average this project solves randomly mixed cubes in 30 something moves in under 2 seconds.

## Getting Started

First you need to have your golang workspace set up on your machine.
Then clone this repo into your go-workspace/src/ folder.

Move into the Rubik folder then download dependencies with ```go get -d ./...```

Finally, to run, go run main.go directly:
```go run main.go mix/subject.txt```

Alternatively build and run the binary:
```go build; ./Rubik mix/subject.txt```

### Usage

```
➜  Rubik git:(master) ✗ go run main.go -h

Usage:	go build; ./Rubik "mix" [-r [length]] [-v] [-g] [-h]

    mix should be valid sequence string e.g.
    "U U' U2 D D' D2 R R' R2 L L' L2 F F' F2 B B' B2"
    or mix "filepath" e.g. "mix/superflip.txt" reads a file
    or mix "-r [len]" or "--random [len]" mixes randomly

    [-v] (--visualizer) show visual of mix and solution
    [-g] (--group) show solution breakdown by subgroup
    [-h] (--help) show usage
```

### 18 valid moves - notation

Valid moves are 90° clockwise twists of the 6 sides (Up, Down, Right, Left, Front, Back) -> U, D, R, L, F, B
90° anti-clockwise twists are denoted with ```'``` -> U', D', R', L', F', B'
180° twists are denoted with ```2``` -> U2, D2, R2, L2, F2, B2

### valid arguments => string, filepath, or -r --random

A sequence of moves can be provided as argument as a string. e.g. "U' F'".
Alternatively, give a filepath containing a valid sequence as argument, there are some example mixes in the ```mix``` folder.

-r (--random) will create and run a random mix.
An optional following ```len``` argument specifies mix length, i.e. ```-r 5``` will create a random 5 move sequence.

### Examples

Here is a basic example with a valid sequence string as argument:

![String](https://github.com/dfinnis/rubik/blob/master/img/sequence_string.png?raw=true)

Here is an example with a mix filepath as argument:

![Filepath](https://github.com/dfinnis/rubik/blob/master/img/filepath.png?raw=true)

Finally is an example with -r ---random argument:

![Random](https://github.com/dfinnis/rubik/blob/master/img/--random.png?raw=true)

### Dependencies

Thankfully, running ```go get -d ./...``` should take care of all dependencies for you.

robotgo -> to type the solution into the visualizer website.

## Tests

Run the test script ```./test.sh``` .

The test script will run 10 static random unit tests, followed by 10 dynamic random tests.
It then displays best, worst and mean for Half-turn metric and solve time.

Finally it runs some unit tests from the ```mix``` folder, to make sure it deals with edge cases, and the mightily hard superflip.

![Rubik test output](https://github.com/dfinnis/rubik/blob/master/img/test.png?raw=true)

## Flags

### Group

```-g``` or ```--group``` shows solution breakdown by subgroup (see [Thistlethwaite's groups](#thistlethwaites-groups) for explanation). Here is an example:

![Group](https://github.com/dfinnis/rubik/blob/master/img/group.png?raw=true)

We start at the top with the mixed cube, described by corner and edge permutation and orientation.

We then see the solution broken down into Thistlethwaite's 4 subgroups.
For each subgroup we see the solution, half-turn metric, and solve time. Followed by the state of the cube after applying the subgroup solution.

At the bottom we should arrive at subgroup 4, a solved cube.
The orientation is all correct (0), and each corner and edge permutation is in its right place (e.g. edge 1 is in permutation 1).

### Visualizer

```-v``` or ```--visualizer``` shows visual of mix and solution. Here is a simple example with a random mix of length 5:

![Visualizer](https://github.com/dfinnis/rubik/blob/master/img/visualizer.gif)

We start with a solved cube, it shows the mix then spins once to show the mixed cube state.
Then it shows the solution, finally spining twice to show the cube is back in its solved state.

# Thistlethwaite's groups

### Cube representation

The sensible way to represent the cube for this style of solution is corner and edge permutation and orientation.

A cube is made up of 26 cubelets. 6 of these are center cubies (one for each face) which cannot move, so we do not need to include these in our model.

There are 12 edge cubies, these have 2 colors, rotate on 2 axes, and so can be orientated correctly (0) or incorrectly (1).

There are 8 corner cubies, these have 3 colors, rotate on 3 axes, and so can be orientated correctly (0) or incorrectly (1) or (2).

For a solved cube (Thistlethwaite's group 4) all the orientations are 0 (correct) and each corner and edge is in its correct permutation (e.g. edge 1 is in permutation 1):

![Solved](https://github.com/dfinnis/rubik/blob/master/img/Solved_cube.png?raw=true)

Here is an image to clarify cubie notation:

![Notation](https://github.com/dfinnis/rubik/blob/master/img/Rubik_notation.png?raw=true)

Check out [this demo](https://iamthecu.be/) for a visual clarification.

### 4 groups



## References


