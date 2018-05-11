Building on the first part of the puzzle, this solution walks through the spiral using the pattern identified in part1.

r u ll dd rrr uuu llll dddd ...

But now the sums are based off what's inside adjacent locations (as described in the part2 README).

My first thought was simply to walk through the spiral and then simply grab the values from all adjacent locations

For instance:

    147  142  133  122   59
    304    5    4    2   57
    330   10    1    1   54
    351   11   23   25   26
    362  747  806

to get 806 you looked at these locations:

    147  142  133  122   59
    304    5    4    2   57
    330   10    1    1   54
    351    X    X    X   26
    362    X  806    X
           X    X    X

more specifically, _these_ locations:

    147  142  133  122   59
    304    5    4    2   57
    330   10    1    1   54
    351    X    X    X   26
    362    X  806

If you work it out, you're looking at these locations when

you're moving right:

    X X X
    X s

(s stands for sum)

you're moving up:

    X
    X s
    X X

you're moving left:

      s X
    X X X

you're moving down:

    X X
    s X
      X

in my code these locations are denoted as compass directions:

    nw n ne
    w  +  e
    sw s se

and I ended up implementing them as bit flags.

