## Algorithm practice dump 

- Chapter1 - Introduction:
  - [Recursive integer multiplication](chapter1/recursiveIntMultiplication.go)
  - [Karatsuba multiplication](chapter1/karatsubaMultiplication.go) => [tests](chapter1/karatsubaMultiplication_test.go)
  - [Merge sort](chapter1/mergeSort.go) => [tests](chapter1/mergeSort_test.go)
  - [Find the second largest number](chapter1/secondLargest.go) => [tests](chapter1/secondLargest_test.go)

- Chapter3 - Divide and Conquer:
  - [Straightforward Matrix Multiplication](chapter3/matrixMultiplication.go) => [tests](chapter3/matrixMultiplication_test.go)
  - [Recusive Matrix Multiplication](chapter3/recursiveMatrixMultiplication.go) => [tests](chapter3/recursiveMatrixMultiplication_test.go)
  - [Strassen Matrix Multiplication](chapter3/strassenMatrixMultiplication.go) => [tests](chapter3/strassenMatrixMultiplication_test.go)
  - [Count Inversions](chapter3/countInversions.go) => [tests](chapter3/countInversions_test.go)
  - [Closest Pair of Points - Brute force](chapter3/closestPairBruteForce.go) => [tests](chapter3/closestPairBruteForce_test.go)
  - [Closest Pair of Points - Divide and conquer](chapter3/closestPair.go) => [tests](chapter3/closestPair_test.go)
  - [Finding maximum number in a unimodal array](chapter3/maxNumberUnimodalArray.go) => [tests](chapter3/maxNumberUnimodalArray_test.go)
  - [Given a sorted array, is there an index i such that A\[i\] equals i](chapter3/doesIndexEqualElementExist.go) => [tests](chapter3/doesIndexEqualElementExist_test.go)

- Chapter5 - Quicksort:
  - TODO [Quick sort with space complexity > O(1) + naive pivot picking]
  - [Quick sort with in place swapping => space complexity == O(1) + naive pivot picking](chapter5/quickSortNaivePivot.go) => [tests](chapter5/quickSortNaivePivot_test.go)
  - [Quick sort with in place swapping => space complexity == O(1) + last element as pivot](chapter5/quickSortLastElementPivot.go) => [tests](chapter5/quickSortLastElementPivot_test.go)
  - [Quick sort with in place swapping => space complexity == O(1) + last element as pivot v2](chapter5/quickSortLastElementPivotv2.go) => [tests](chapter5/quickSortLastElementPivotv2_test.go)
  - [Quick sort with in place swapping => space complexity == O(1) + 3 number median pivot picking](chapter5/quickSortMedianPivot.go) => [tests](chapter5/quickSortMedianPivot_test.go)
  - TODO [Quick sort with in place swapping => space complexity == O(1) + randomized pivot picking]

- Chapter 6 - Linear Time Selection:
  - [Reducing Selection to Sorting (using MergeSort)](chapter6/rselectMergeSort.go) => [tests](chapter6/rselectMergeSort_test.go)
  - [RSelect](chapter6/rselect.go) => [tests](chapter6/rselect_test.go)

- Chapter 7 - Graphs: The Basics
  - [Representing (simple undirected) graphs using Adjacency List](chapter7/adjacencyList.go)
  - [Representing (simple undirected) graphs using Adjacency Matrix](chapter7/adjacencyMatrix.go)

- Chapter 8 - Graph Search and its Applications
  - [Generic Graph Search](chapter8/genericGraphSearch.go) => [tests](chapter8/genericGraphSearch_test.go)
  - [Breadth First Search](chapter8/breadthFirstSearch.go) => [tests](chapter8/breadthFirstSearch_test.go)
  - [Shortest Paths using BFS](chapter8/shortestPathsBFS.go) => [tests](chapter8/shortestPathsBFS_test.go)
  - [Compute Connected Components for Undirected Graph](chapter8/undirectedConnectedComponents.go) => [tests](chapter8/undirectedConnectedComponents_test.go)
  - [Depth first search](chapter8/depthFirstSearch.go) => [tests](chapter8/depthFirstSearch_test.go)
  - [Depth first search recursive](chapter8/depthFirstSearchRecursive.go) => [tests](chapter8/depthFirstSearchRecursive_test.go)
  - [Topological Sort using Depth First Search for Directed Acyclic Graphs](chapter8/dfsTopologicalOrder.go) => [tests](chapter8/dfsTopologicalOrder_test.go)
  - [Kosaraju algorithm for finding strongly connected components in a DG](chapter8/kosaraju.go) => [tests](chapter8/kosaraju_test.go)

- Chapter 9 - Dijkstra’s Shortest-Path Algorithm
  - [Djakstra's Shortest Path for Directed Graph](chapter9/djakstraShortestPathDirected.go) => [tests](chapter9/djakstraShortestPathDirected_test.go)
