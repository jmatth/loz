# Loz

`loz` is a library for working with Go iterators in a functional style, inspired by [`samber/lo`][lo] and standard library functionality in most other languages such as [Rust's Iterator][rust-iterator], [Dart's Iterable][dart-iterable], [Java's Stream][java-stream], etc. See also [this article][pipelining], which is not affiliated with this project but explains the motivation for using wrapper types with methods instead of bare functions pretty well.

The name is because it's like `lo`, but la`z`y.

## Install

```shell
go get github.com/jmatth/loz@v1
```

## Usage

`loz` provides two primary helper types: `loz.Seq` and `loz.KVSeq`, which correspond to `iter.Seq` and `iter.Seq2` respectively. Helper methods are provided to easily iterate over slices and maps:

```go
evenNums := loz.IterSlice([]int{1, 2, 3, 4, 5}).
	Filter(func(n int) bool { return n%2 == 0 }).
	Collect(loz.ToSlice[int]())
fmt.Println(evenNums)
// Output: [2 4]

systemUsers := loz.IterMap(map[int]string{
	0:    "root",
	1:    "bin",
	81:   "dbus",
	33:   "http",
	1000: "josh",
	1001: "katie",
}).
	Filter(func(id int, _ string) bool { return id < 1000 }).
	Values().
	Collect(loz.ToSlice[string]())
fmt.Println(systemUsers)
// Output: [root bin dbus http]
```

You can also wrap an `iter.Seq` or `iter.Seq2` manually if needed:

```go
loz.Seq(slices.Values([]int{1, 2, 3}))
loz.KVSeq(maps.All(map[string]int{"one": 1, "two": 2, "three": 3}))
```

## Error handling

Sometimes you may encounter an error that is unrecoverable and should cause
iteration to halt. To support this, loz provides `loz.PanicHaltIteration`.
Calling this function with a non-nil `error` will immediately panic with a
wrapped version of the provided error. To then recover from this panic, all
terminal iterator methods have a "Try" version: `TryCollect`, `TryFirst`, etc.
These methods all recover from the panic, unwrap the original error, and return
it along with zero values for their other returns. Here is a full example:

```go
result, err := loz.IterSlice([]string{"1", "foo", "3"}).
	Map(func(s string) int {
		num, err := strconv.Atoi(s)
		loz.PanicHaltIteration(err)
		return num
	}).TryCollect(loz.ToSlice[int]())
fmt.Printf("%v; %v", result, err)
// Output: []; strconv.Atoi: parsing "foo": invalid syntax
```

**Why separate "Try" versions of these methods?**

Two reasons:
- The normal versions of the methods and the code that calls them can be simpler for not having to deal with a second return of type `error` that is never used.
- Not having a `defer` in those functions means they might be inlined by the compiler.

**Important notes:**

- The "Try" suite of methods only recovers from panics raised by `PanicHaltIteration`. They will re-panic if they encounter a panic from any other source.
- Calling `PanicHaltIteration` outside an iterator that terminates with a "Try" method will just panic as normal.
- Iterators are evaluated lazily, so it is possible that some processing will occur before the error is encountered. This is especially relevant when using `ForEach`.
- Similarly, a terminal method like `TryFirst` that stops before consuming the entire iterator may return successfully even if later elements in the iterator would have caused an error.

[lo]: https://github.com/samber/lo
[rust-iterator]: https://doc.rust-lang.org/std/iter/trait.Iterator.html
[java-stream]: https://docs.oracle.com/javase/8/docs/api/java/util/stream/Stream.html
[dart-iterable]: https://api.dart.dev/dart-core/Iterable-class.html
[pipelining]: https://herecomesthemoon.net/2025/04/pipelining/