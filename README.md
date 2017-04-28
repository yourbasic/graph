# Your basic graph

Golang library of basic graph algorithms

### Generic graph algorithms

The algorithms can be applied to any graph data structure implementing
the two Iterator methods: Order, which returns the number of vertices,
and Visit, which iterates over the neighbors of a vertex.

All algorithms operate on directed graphs with a fixed number
of vertices, labeled from 0 to n-1, and edges with integer cost.
An undirected edge {v, w} of cost c is represented by the two
directed edges (v, w) and (w, v), both of cost c.
A self-loop, an edge connecting a vertex to itself,
is both directed and undirected.

### Graph data structures

The type Mutable represents a directed graph with a fixed number
of vertices and weighted edges that can be added or removed.
The implementation uses hash maps to associate each vertex
in the graph with its adjacent vertices. This gives constant
time performance for all basic operations.

The type Immutable is a compact representation of an immutable graph.
The implementation uses lists to associate each vertex in the graph
with its adjacent vertices. This makes for fast and predictable
iteration: the Visit method produces its elements by reading
from a fixed sorted precomputed list. This type supports multigraphs.

### Virtual graphs

The subpackage graph/build offers a tool for building virtual graphs.
In a virtual graph no vertices or edges are stored in memory,
they are instead computed as needed. New virtual graphs are constructed
by composing and filtering a set of standard graphs, or by writing
functions that describe the edges of a graph.

### Roadmap

* The API of this library is frozen.
* Bug fixes and performance enhancement can be expected.
* New functionality might be included.

The only accepted reason to modify the API of this package is to
handle bug fixes that can't be resolved in any other reasonable way.

New features and performance enhancements are limited to basic
algorithms and data structures, akin to the ones that you might find
in a computer science textbook.

Stefan Nilsson – [korthaj](https://github.com/korthaj)
