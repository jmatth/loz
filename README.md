# Loz

`loz` is a library for working with Go iterators in a functional style, inspired by [`samber/lo`][lo] and standard library functionality in most other languages such as [Rust's Iterator][rust-iterator], [Dart's Iterable][dart-iterable], [Java's Stream][java-stream], etc. See also [this article][pipelining], which is not affiliated with this project but explains the motivation for using wrapper types with methods instead of bare functions pretty well.

The name is because it's like `lo`, but la`z`y.

## Install

```shell
go get github.com/jmatth/loz@v1
```

## Usage

`loz` provides two primary helper types: `loz.Seq` and `loz.KVSeq`, which correspond to `iter.Seq` and `iter.Seq2` respectively. You can wrap either of those types directly:

```go
var numsIter iter.Seq[int] = slices.Values([]int{1, 2, 3, 4, 5})
evenNums := loz.Seq[int](numsIter).
	Filter(func(n int) bool { return n%2 == 0 }).
	CollectSlice()
fmt.Println(evenNums)
// Output: [2 4]

var idNameIter iter.Seq2[int, string] = maps.All(map[int]string{
	0:    "root",
	1:    "bin",
	81:   "dbus",
	33:   "http",
	1000: "josh",
	1001: "katie",
})

systemUsers := loz.KVSeq[int, string](idNameIter).
	Filter(func(id int, _ string) bool { return id < 1000 }).
	Values().
	CollectSlice()
fmt.Println(systemUsers)
// Output: [root bin dbus http]
```

...or use some of the included helpers to create iterators with less code and fewer type annotations:

```go
evenNums := loz.IterSlice([]int{1, 2, 3, 4, 5}).
	Filter(func(n int) bool { return n%2 == 0 }).
	CollectSlice()
fmt.Println(evenNums)
// Output: [2 4]

systemUsers :=
	loz.IterMap(map[int]string{
		0:    "root",
		1:    "bin",
		81:   "dbus",
		33:   "http",
		1000: "josh",
		1001: "katie",
	}).
		Filter(func(id int, _ string) bool { return id < 1000 }).
		Values().
		CollectSlice()
fmt.Println(systemUsers)
// Output: [root bin dbus http]
```

## Mapping

Creating a method to apply a mapping function to all values in a collection presents an interesting challenge in Go. Go's type system does not allow methods to take generic parameters; only functions without a receiver and type definitions can have generic types. To work around this, loz relies on code generation to create a series of types: `Map1` through `Map9`, each of which takes `n+1` type parameters, where the first parameter represents the initial type of the elements in the iterator and each successive type represents the type that will be returned by the next call to `.Map()`. Here is a contrived example:

```go
medians := loz.Map3[string, []string, []int, int](loz.IterSlice([]string{
	"1,2,3",
	"100",
	"3,5",
	"1,6,30,42,70",
})).
	Map(func(s string) []string { return strings.Split(s, ",") }).
	Map(func(sl []string) []int {
		return loz.Map1[string, int](loz.IterSlice(sl)).
			Map(func(s string) int {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				return num
			}).
			CollectSlice()
	}).
	Filter(func(nums []int) bool { return len(nums) > 2 && len(nums)%2 != 0 }).
	Map(func(nums []int) int { return nums[len(nums)/2] }).
	CollectSlice()
fmt.Println(medians)
// Output: [2 30]
```

Similar types exist for `KVSeq`, named `KVMap1` through `KVMap9`.

A limit of 9 was chosen only because godoc arranges the index lexicographically and having `Map10`-`Map19` sorted before `Map2` etc. would make the documentation look even worse than it already does. Realistically you will probably never get close to 9 map operations on a single iterator, but if you have a use case where this is a limitation please open an issue.

[lo]: https://github.com/samber/lo
[rust-iterator]: https://doc.rust-lang.org/std/iter/trait.Iterator.html
[java-stream]: https://docs.oracle.com/javase/8/docs/api/java/util/stream/Stream.html
[dart-iterable]: https://api.dart.dev/dart-core/Iterable-class.html
[pipelining]: https://herecomesthemoon.net/2025/04/pipelining/