Self-Descriptive Sentences
==========================

A **self-descriptive sentence**, or an
[autogram](http://en.wikipedia.org/wiki/Autogram), is an sentence that
describes exactly how many times each letter occurs within it. For example:

> Selena's neat sentence has five as, one b, two cs, two ds, thirty-six es,
  five fs, one g, six hs, eleven is, one j, one k, three ls, one m, twenty-
  three ns, thirteen os, one p, one q, six rs, twenty-nine ss, seventeen ts,
  two us, seven vs, six ws, five xs, four ys, and one z.

Such sentences are tricky to find, at least for this program! Written in
[Go](https://golang.org/), it can take anywhere from a few seconds to hours to
find solutions for various prefixes. Goroutines are used to allow more than
one searcher to run at the same time, and so you can easily search for
sentences with slightly modified prefixes.

How it Works
------------

The program works by representing sentence as a vector of 26 `int`s. To find
an [autogram](http://en.wikipedia.org/wiki/Autogram), the sentences vector is
first assigned random values, and then the true count of the the resulting
sentence is calculated. If this true count is the same as the sentences
vector, then we've found an [autogram](http://en.wikipedia.org/wiki/Autogram).
However, if they don't match, then the true count is assigned to be the
sentence count, and the process continues like this until an
[autogram](http://en.wikipedia.org/wiki/Autogram) is found, or a duplicate
sentence is encountered.

This was done mainly as an exercise in using Go, and the program could be
definitely improved in terms of ease of use, and probably also performance.
