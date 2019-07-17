# Examples

The "bitmap is just an interface" idea offers a lot of interesting
potential, and this is a set of possible applications/use cases that
came up while discussing how it might be used. The idea is to have
these written down so they can be used to guide thoughts about the
interface design.

## Composition

One of the big appeals is composition -- the ability to make bitmap
implementations which are basically just wrappers on another bitmap
implementation.

### Locking

Given a bitmap implementation which does not have any locking or
concurrency, it's trivial to make one where every operation just looks
like
	```
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.embedded.ThisOperation(args)
	```

This provides a safe-but-slow implementation at essentially zero
further developer cost.

### As-needed unmarshalling

For read-only implementations, it's possible to just read a raw roaring
bitmap as-needed to generate containers. (Modifying them is harder, but
we do have the option of using an ops log.)

### Making streams seekable

Given an unseekable Bitmap, such as one based on a remote API call, it's
practical to implement a caching bitmap on top of it which reads it all
in. Assuming there's enough memory.

### Speculative/conditional execution

Operations like union/intersect admit some short-circuiting; you don't need to
retrieve values to intersect them with the empty set, for instance. Once we
have the concept of bitmaps which may not have actually retrieved all of their
data, we potentially have the option of skipping unneeded sections when
doing operations on them. In other words, instead of Intersect being an
operation which actually requires the full data of its input bitmaps, and
produces the full output, it can yield a Bitmap which waits until someone
requests a given range of bits, then computes those bits -- and may not need
to request the corresponding bits from all of its inputs.
