# zf
zip the contents of two files

Say `one.txt` has the following contents:

    1
    2
    3

And `two.txt`:

    A
    B
    C

Running:

    zf -file1 one.txt -file2 two.txt

Will output the following to stdout:

    1
    A
    2
    B
    3
    C

## TODO:

* `zf file1.txt file2.txt ... fileN.txt`
* add separator flag (i.e. something other than the newline character)
* customizable (dynamic?) buffer maxTokenSize

## Maintenance
Recommendations? Improvements? Open an [issue](https://github.com/gypsydiver/zf/issues/new) or submit a [Pull Request](https://github.com/gypsydiver/zf/compare).
