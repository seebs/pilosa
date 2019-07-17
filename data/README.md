# Data interface

The `data` package provides the interface to add new data sources, sinks,
representations, and operations to Pilosa.

The primary exports are interfaces; other packages provide concrete
implementations.

In general, Pilosa is most concerned with bitmaps, so the primary visible type
is a Bitmap, which represents a sequence of up to 2^64 bits.

Historically, Pilosa's primary data representation has been Roaring bitmaps,
or variants thereof, and this provides a reasonably useful abstraction of a
Container, which is a region of 2^16 bits within a Bitmap. (Containers are
always 2^16-aligned.) Thus, whether or not a given Bitmap is in fact
implemented as a Roaring bitmap, it should provide an interface allowing
navigating through it as a sequence of Containers.

The most common representations in Roaring bitmaps are "bitmaps", which are
slices of 1024 uint64 values (representing 2^16 bits as a bit vector), and
"arrays", which are sorted slices of uint16. These are given the names
BitVec and Slice respectively, and every Container implementation should be
able to provide these representations.

Less common is the "interval" or "run" container type, which represents bits
as a series of uint16 first/last pairs.

# Design Notes

Why don't Containers take multiple arguments to things like UnionInPlace, while
Bitmaps do? Because it doesn't help us much, and complicates the code. Bitmaps
do because there's some optimizations to be had when performing a large number
of union operations simultaneously, for instance; this doesn't apply
per-container, though.
