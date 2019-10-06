# Genetic-Algorithm
Go / golang Genetic Algorithm Lightweight Demo

Just 125 lines with no dependancies (except go).
All it does is quickly match a string (even hundreds of characters).
It starts with a population of random char strings and evolvs them to match the goal string.

1. Select population strings (fittest) that have the most matches with goal.
2. Replace population with strings bread from fittest.
3. Replace a few chars in the population with a random character.
4. Repeat until one of population matches goal.

Usefull? Well maybe not since we allready know the goal.
But it does simply show in a few lines of code how to evolve a solution using:
* fitness culling
* breeding (crossover)
* mutation

Add it to your go projects and try it today!
