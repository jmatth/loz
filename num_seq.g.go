package loz

func (n NumSeq[V]) CollectSlice() []V {
	return Seq[V](n).CollectSlice()
}
func (n NumSeq[V]) TryCollectSlice() ([]V, error) {
	return Seq[V](n).TryCollectSlice()
}
func (n NumSeq[V]) ForEach(process processor[V]) {
	Seq[V](n).ForEach(process)
}
func (n NumSeq[V]) TryForEach(process processor[V]) error {
	return Seq[V](n).TryForEach(process)
}
func (n NumSeq[V]) Map(mapper mapper[V, V]) NumSeq[V] {
	return NumSeq[V](Seq[V](n).Map(mapper))
}
func (n NumSeq[V]) Reduce(combine reducer[V, V]) (V, error) {
	return Seq[V](n).Reduce(combine)
}
func (n NumSeq[V]) TryReduce(combine reducer[V, V]) (V, error) {
	return Seq[V](n).TryReduce(combine)
}
func (n NumSeq[V]) Fold(initial V, combine reducer[V, V]) V {
	return Seq[V](n).Fold(initial, combine)
}
func (n NumSeq[V]) TryFold(initial V, combine reducer[V, V]) (V, error) {
	return Seq[V](n).TryFold(initial, combine)
}
func (n NumSeq[V]) First() (V, error) {
	return Seq[V](n).First()
}
func (n NumSeq[V]) TryFirst() (V, error) {
	return Seq[V](n).TryFirst()
}
func (n NumSeq[V]) Last() (V, error) {
	return Seq[V](n).Last()
}
func (n NumSeq[V]) TryLast() (V, error) {
	return Seq[V](n).TryLast()
}
func (n NumSeq[V]) Any(test yielder[V]) bool {
	return Seq[V](n).Any(test)
}
func (n NumSeq[V]) TryAny(test yielder[V]) (bool, error) {
	return Seq[V](n).TryAny(test)
}
func (n NumSeq[V]) None(test yielder[V]) bool {
	return Seq[V](n).None(test)
}
func (n NumSeq[V]) TryNone(test yielder[V]) (bool, error) {
	return Seq[V](n).TryNone(test)
}
func (n NumSeq[V]) Every(test yielder[V]) bool {
	return Seq[V](n).Every(test)
}
func (n NumSeq[V]) TryEvery(test yielder[V]) (bool, error) {
	return Seq[V](n).TryEvery(test)
}
func (n NumSeq[V]) Filter(filter yielder[V]) NumSeq[V] {
	return NumSeq[V](Seq[V](n).Filter(filter))
}
func (n NumSeq[V]) Skip(toSkip int) NumSeq[V] {
	return NumSeq[V](Seq[V](n).Skip(toSkip))
}
func (n NumSeq[V]) SkipWhile(test yielder[V]) NumSeq[V] {
	return NumSeq[V](Seq[V](n).SkipWhile(test))
}
func (n NumSeq[V]) Take(toTake int) NumSeq[V] {
	return NumSeq[V](Seq[V](n).Take(toTake))
}
func (n NumSeq[V]) TakeWhile(test yielder[V]) NumSeq[V] {
	return NumSeq[V](Seq[V](n).TakeWhile(test))
}
func (n NumSeq[V]) Indexed() KVSeq[int, V] {
	return Seq[V](n).Indexed()
}
func (n NumSeq[V]) Expand(toElements mapper[V, Seq[V]]) NumSeq[V] {
	return NumSeq[V](Seq[V](n).Expand(toElements))
}

func (n NumSeq[V]) Max() (V, error) {
	return OrdSeq[V](n).Max()
}
func (n NumSeq[V]) Min() (V, error) {
	return OrdSeq[V](n).Min()
}
