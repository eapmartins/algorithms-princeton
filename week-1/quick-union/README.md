# Quick Union
## Complexity Analysis

The initial implementation of QuickUnion also has a time complexity of O(n^2) because each union operation takes O(n) time. As a result, if there are n unions performed, the overall time complexity becomes quadratic.

Different from QuickFind, the QuickUnion algorithm is based on the idea of a tree. Each element in the array is a node in the tree. The value of the element is the parent of the node. The root of the tree is the node that is its own parent. The root of the tree is the representative of the set.

## References

Study Guide [Union Find](https://www.cs.princeton.edu/courses/archive/spring19/cos226/lectures/study/15UnionFind.html)