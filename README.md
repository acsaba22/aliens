# Alien invasion

This tool simulates the upcoming alien invasion.

## Problem description

We start with a map of cities connected by roads, whatever graph, we don't assume grid or whether
it's drawable on a plane. Aliens got into the cities, they move around, they fight and destroy
cities for 10k rounds.

Input is a file of cities and roads; and a number of aliens (N).

Assumptions about the map:
* Each line names one city and max 4 connected cities: north, south, east, west.
* One city is described only once.
* A road can be described only from one direction, but it's considered to be bi-directional.
* A city may not be described, as the main city of the line, still it can exists if other cities
  mentioned it.
* The above rules means that a city may have more than 4 roads.
* A city may be mentioned two or more times as neighbour, but then the likelihood of moving there
  increases.
* A city may have a connection to itself, in this case the alien can stay in the city.

Assumptions about aliens:
* The input is just their number (N)
* The simulating program generates their location randomly in distinct cities
  (N <= number of cities).
* The game has rounds, at each round all aliens move at the same time to a random neighbouring city.
* No alien stays in the same city at a round, they have to move,
  they may move back in the following round only.
* If two or more aliens get into the same city then they fight. All the aliens in the city and the
  city is destroyed.
* If there are no more neighbouring cities, the alien will stay in the city.

## Running the code

Unit tests (including example tests can be run with):

```
go test -v ./...
```

Running the simulation:

```
cat data/input-10_70.txt | go run ./cmd/aliens 7
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

Sync with client if the current output is satisfactory: e.g. "NY has been destroyed by 3 aliens!".
This is contrary to the original request with Alien IDs.

Testing is not perfect, it's not blackbox testing mostly because of the random number generation.
Internal state is modified from tests.
Possible solution could be:
* A more controlled design where the input also specifies alien locations and movement strategy
* Random number generator abstracted away so tests can specify their own. This would be probably
  too brittle (acceptable code changes could brake the tests).
