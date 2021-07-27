# Hash Tables
Hash table is a data structure that is used to store key value pairs.
Hash tables have various names across different languages. 

- Hashtables
- Maps
- Unordered maps
- Dictionaries
- Objects, ...

For hash table its more of a key value pair. The way the hash table works is like it will have a key which is used for indexing the value that needs to be stored to memory.

So consider a key value pair like `cash 10,000`. Here the key is `cash` and value is `10,000`. So to find the value `10,000` the key `cash` is used. Its done by hasing the key. 

    cash ---> Hasing --> bfa23404kla23ll --> Index space(Address) ---> 10,000

so when we lookup for key, internally its hash is found and value for that hash is returned.

## Hash function

A hash function is simply a function that generates a fixed length value for each input we pass in. Hash function is idempotent.

## BigO Notation for Hash Function

- Insert        O(1)
- lookup        O(1)
- delete        O(1)
- search        O(1)

# Cons of Hash Tables.

Though ideally it looks perfect. By understanding this much we will be forced to think that hash tables is ideal for all our needs since it does have all ILDS functions under O(1). But there is a small catch. The problem occurs due to the way the data is stored based on this hash value.


