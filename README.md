# zf
zip the contents of N files

Say `one.txt` has the following contents:

    1
    2
    3

`two.txt`:

    A
    B
    C

And `three.txt`:

    a
    b
    c

Running:

    zf one.txt two.txt three.txt

Will output the following to stdout:

    1
    A
    a
    2
    B
    b
    3
    C
    c

## TODO:
* add separator flag (i.e. something other than the newline character)
* customizable (dynamic?) buffer maxTokenSize

## Maintenance
Recommendations? Improvements? Open an [issue](https://github.com/gypsydiver/zf/issues/new) or submit a [Pull Request](https://github.com/gypsydiver/zf/compare).
