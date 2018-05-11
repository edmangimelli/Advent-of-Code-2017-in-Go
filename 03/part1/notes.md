using their example:

    17  16  15  14  13
    18   5   4   3  12
    19   6   1   2  11
    20   7   8   9  10
    21  22  23---> ...

starting from 1, you move like so:
r u l l d d r r r u u u l l l l d d d d

here's the pattern:
r once
u once
l twice
d twice
r 3x
u 3x
l 4x
d 4x
...and so on

my code simply steps through a graph (and assumes the center spot is [0,0]) moving in the spiral pattern, all the while counting (actually, decrementing) and testing to see if we reached our number.

If we have, we simply add the x and y coordinates (absolute value) and we have our # of steps.

