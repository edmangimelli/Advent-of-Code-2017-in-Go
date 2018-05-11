Instead of trying to approach this problem with a coordinate system, I wondered if I could see the pattern a different way.

Here's the example spiral given:

    147  142  133  122   59
    304    5    4    2   57
    330   10    1    1   54
    351   11   23   25   26
    362  747  806--->   ...

You could unravel it and write it this way:

    loc  val
     0     1
     1     1
     2     2
     3     4
     4     5
     5    10
     6    11
     7    23
     8    25
     9    26
    10    54
    11    57
    12    59
    13   122
    14   133
    15   142
    16   147
    17   304
    18   330
    19   351
    20   362
    21   747
    22   806

I then thought about how those values were being reached:

    loc  val  locs summed
     0     1   given
     1     1   0
     2     2   0 1 
     3     4   0 1 2 
     4     5   0 3 
     5    10   0 3 4 
     6    11   0 5 
     7    23   0 1 5 6 
     8    25   0 1 7 
     9    26   1 8 
    10    54   1 2 8 9 
    11    57   1 2 10 
    12    59   2 11 
    13   122   2 3 11 12 
    14   133   2 3 4 13 
    15   142   3 4 14 
    16   147   4 15 
    17   304   4 5 15 16 
    18   330   4 5 6 17 
    19   351   5 6 18 
    20   362   6 19 
    21   747   6 7 19 20 
    22   806   6 7 8 21 

(for instance, the value 747 is the sum of the values stored at locations 6, 7, 19, and 20)

Then I tried to see if I could see a pattern. The first thing I noticed is that the last number of each "sum list" is the previous location:

    loc  val  locs summed
     0     1   given
     1     1  (0)
     2     2   0 (1) 
     3     4   0 1 (2) 
     4     5   0 (3)
     5    10   0 3 (4)
     6    11   0 (5)
     7    23   0 1 5 (6)
     ....

Nifty! So I put those aside for a moment:

    loc  val  locs summed
     0     1   given
     1     1  
     2     2   0
     3     4   0 1
     4     5   0
     5    10   0 3
     6    11   0
     7    23   0 1 5
     8    25   0 1
     9    26   1
    10    54   1 2 8
    11    57   1 2
    12    59   2
    13   122   2 3 11
    14   133   2 3 4
    15   142   3 4
    16   147   4
    17   304   4 5 15
    18   330   4 5 6
    19   351   5 6
    20   362   6
    21   747   6 7 19
    22   806   6 7 8

And then I thought that I probably need to see more of the list:

    loc  val  locs summed
     0     1   given
     1     1  
     2     2   0
     3     4   0 1
     4     5   0
     5    10   0 3
     6    11   0
     7    23   0 1 5
     8    25   0 1
     9    26   1
    10    54   1 2 8
    11    57   1 2
    12    59   2
    13   122   2 3 11
    14   133   2 3 4
    15   142   3 4
    16   147   4
    17   304   4 5 15
    18   330   4 5 6
    19   351   5 6
    20   362   6
    21   747   6 7 19
    22   806   6 7 8
               7 8 9 
               8 9 
               9
               9 10 24 
               9 10 11 
               10 11 12 
               11 12 
               12
               12 29

Aha! I saw a familiar pattern:

    loc  val  locs summed
     0     1   given
     1     1  
     2     2   0 -------------
     3     4   0 1
     4     5   0 -------------
     5    10   0 3
     6    11   0 -------------
     7    23   0 1 5
     8    25   0 1
     9    26   1 -------------
    10    54   1 2 8
    11    57   1 2
    12    59   2 -------------
    13   122   2 3 11
    14   133   2 3 4
    15   142   3 4
    16   147   4 -------------
    17   304   4 5 15
    18   330   4 5 6
    19   351   5 6
    20   362   6 -------------
    21   747   6 7 19
    22   806   6 7 8
               7 8 9 
               8 9 
               9 -------------
               9 10 24 
               9 10 11 
               10 11 12 
               11 12 
               12 ------------
               12 29

Between those lines are 1 list, 1 list, 2 lists, 2 lists, 3, 3, 4, 4, ...

After that I could see another pattern, but it took me a bit to be able to express it.
Look at location 20, its list is just "6". Below it you see the following numbers cascading.

    6
    6 7 19
    6 7 8
    7 8 9
    8 9
    9

There's definitely a pattern there. After thinking about what's happening on the spiral when you're moving, it made sense. The lone 6 is a corner, then, when you round the corner, you get the location before and the next (and 19). After that, you are no longer near a corner, and you'll have 3 beside your location until you get near the next corner (the lone 9).

Remember! the previous location is always included in the sum but at the moment we're ignoring it.

    16 15 14 13 12
    17  4  3  2 11
    18  5  0  1 10
    19  6  7  8  9
    20               20 grabs 6 (and 19)

    16 15 14 13 12
    17  4  3  2 11
    18  5  0  1 10
    19  6  7  8  9
    20 21            21 grabs 6 7 and 19 (and 20)

    16 15 14 13 12
    17  4  3  2 11
    18  5  0  1 10
    19  6  7  8  9
    20 21 22         22 is away from the corner and has 3 above it

    16 15 14 13 12
    17  4  3  2 11
    18  5  0  1 10
    19  6  7  8  9
    20 21 22 23      23 is still away from the corner and has 3 above it

24 will be near the corner and will have just 8 and 9

and 25 will _be_ the corner and will have just 9

From there I could extrapolate that when the spiral is bigger a similar pattern will occur when you're moving across an edge.

The lone number, then the lone number and the next one, then you have the 3 numbers, and you keep cruising along with 3 numbers (each time they increase by 1) until you get to the next corner, where it'll be the last 2 numbers of the previous 3 and then the last number for the corner.

Hypothetical edge (Note! These values might not actually occur this way)

    164
    164 165
    164 165 166
    165 166 167
    166 167 168
    167 168 169
    168 169
    169

The one exception is that the second row of the cascade always has the location two previous (in our previous example it was 19).

How many cascading rows before each corner? That follows the original pattern from part1 (1,1,2,2,3,3,4,4)

To see my algorithm, look at my code. The way I've written it, I had to populate the first 12 locations. Maybe at some point I'll make a gif illustrating these patterns better.

This solution is around 8x faster than my straightforward solution.

