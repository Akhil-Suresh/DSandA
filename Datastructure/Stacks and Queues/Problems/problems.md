# Problems based on stacks and queue.

# Stack

1. Implement stack based on queue.
    * push simply store items to queue -> [1,2,3]
    * when we pop it we requeue items to same queue till length-1
        - first queue pop -> 1
        - queue push -> [2,3,1]
        - queue pop -> 2
        - queue push -> [3,1,2]
        - queue pop -> 3