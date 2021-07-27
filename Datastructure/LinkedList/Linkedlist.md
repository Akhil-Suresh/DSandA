# Linked list

A list that is linked! :D

head -> [] -> tail -> END(null/nil)

Linked list is a linear data structure, in which elements are not stored in continuous memory locations. The elements of linked list are stored using pointers.

Likedlist are mainly classified into Singly Linked list, Doubly linked list and Circular Linked list.

BigO Notation
| Operation | Big O |
| --------- |------ |
| prepend   |  O(1) |
| append    |  O(1) |
| insert    |  O(n) |
| delete    |  O(n) |
| search    |  O(n) |


## Pros and Cons of Linked list

### Pros
- inserting in a linked list is faster since inserting to an array require shifting things to a new space.
- Dynamic size. Liked list can grow in size without any issues.

### Cons
- Needs extra space for pointer storage
- Not Cache friendly.
    - Caching system which is more efficient in reading continuous memory items than from scattered memory location.
- In an array elemets are indexed. In linked list you start at head, traverse till you find the index.