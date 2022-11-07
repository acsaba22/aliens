# Alien invasion

This tool simulates the upcoming alien invasion.

# Problem description

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

Assumptions about aliens:
* The input is just their number (N)
* The simulating program generates their location randomly
* The game has rounds, at each round all aliens move at the same time to a random neighbouring city.
* No alien stays in the same city at a round, they have to move,
  they may move back in the following round only.
* If two or more aliens get into the same city then they fight. All the aliens in the city and the
  city is destroyed.
