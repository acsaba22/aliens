# Alien invasion

This tool simulates the upcoming alien invasion. There is PDF description which should be read
first.

## Problem description

We start with a map of cities connected by roads, whatever graph, we don't assume grid or whether
it's drawable on a plane. Aliens got into the cities, they move around, they fight and destroy
cities for 10k rounds.

Input is a file of cities and roads; and a command line argument: the number of aliens (N).

Assumptions about the map:
* Each line names one city and max 4 connected cities: north, south, east, west.
* One city is described only once.
* A road can be described only from one direction, but it's considered to be bi-directional.
* A city may be missing as the main city (first of a line), still it can exists if other cities
  mentioned it.
* The above rules means that a city may have more than 4 roads.
* A city may be mentioned two or more times as neighbour in the same line, but then the likelihood
  of moving there increases.
* A city may have a connection to itself, in this case the alien can stay in the city.

Assumptions about aliens:
* The input is just their number (N)
* The simulating program generates their location randomly in distinct cities
  (N <= number of cities).
* The game has rounds, at each round all aliens move at the same time to a random neighbouring city.
* If two or more aliens get into the same city then they fight. All the aliens in the city and the
  city is destroyed.
* No alien stays in the same city at a round, they have to move, they may move back in the
  following round only. Exception to this rule: (a) no more neighbouring cities exist,
  (b) the city has a road to itself.

## Running the code

Unit tests (including example tests can be run with):

```
go test -v ./...
```

Running the simulation:

```
cat data/input-10_70.txt | go run ./cmd/aliens 8
```

Generating input files:

```
go run ./cmd/generate 1000 70 > data/input-1000_70.txt
```

Running benchmark tests:

```
go test  -v -run=NONE -bench=. ./...
```

## Possible next steps


Testing is not perfect, it's not blackbox testing mostly because of the random number generation.
Internal state is modified from tests.
Possible solution could be:
* A more controlled design where the input also specifies alien locations and movement strategy.
  In a real live environment hopefuly this would be the case, the program would be predetermined
  from the inputs.
* Random number generator abstracted away so tests can specify their own. This would be probably
  brittle (acceptable changes to the implementation could brake the tests).

## Self assessment

### Chosen problem

The problem could have been interpreted in many ways. Some options:
1) city grid.
2) generic graph with one directional roads, this would have kept the connections of a city under 4.
3) generic graph with two dimmensional roads.

I've chosen (3) because it more general then (1), accepts more inputs.

I've chosen (3) over (2) because one direcitonal roads seems counter intuitive in a world.

Because I've chosen (3), the reverse roads means that more than 4 connections can go out of a city,
therefore in the final printing printing roads both ways is not always possible.

So the original city description is kept in parallel with a generated graph of bi-directional roads.

### Distinct aliens

I've relaxed the problem description by not keeping track of unique IDs of the aliens.
Therefore the output is like "NY has been destroyed by 3 aliens!" instead of
"NY has been destroyed by alien 10 and alien 11 and alien 12."

In a real situation when I see that the implementation would be easier with slight modification
of the problem I would ask / negotiate with the respected parties.

