package loz

func (o OrdSeq[V]) CollectSlice() []V {
	return Seq[V](o).CollectSlice()
}
func (o OrdSeq[V]) ForEach(process processor[V]) {
	Seq[V](o).ForEach(process)
}
func (o OrdSeq[V]) Map(mapper mapper[V, V]) OrdSeq[V] {
	return OrdSeq[V](Seq[V](o).Map(mapper))
}
func (o OrdSeq[V]) Reduce(combine reducer[V, V]) (V, error) {
	return Seq[V](o).Reduce(combine)
}
func (o OrdSeq[V]) Fold(initial V, combine reducer[V, V]) V {
	return Seq[V](o).Fold(initial, combine)
}
func (o OrdSeq[V]) First() (V, error) {
	return Seq[V](o).First()
}
func (o OrdSeq[V]) Last() (V, error) {
	return Seq[V](o).Last()
}
func (o OrdSeq[V]) Any(test yielder[V]) bool {
	return Seq[V](o).Any(test)
}
func (o OrdSeq[V]) None(test yielder[V]) bool {
	return Seq[V](o).None(test)
}
func (o OrdSeq[V]) Every(test yielder[V]) bool {
	return Seq[V](o).Every(test)
}
func (o OrdSeq[V]) Filter(filter yielder[V]) OrdSeq[V] {
	return OrdSeq[V](Seq[V](o).Filter(filter))
}
func (o OrdSeq[V]) Skip(toSkip int) OrdSeq[V] {
	return OrdSeq[V](Seq[V](o).Skip(toSkip))
}
func (o OrdSeq[V]) SkipWhile(test yielder[V]) OrdSeq[V] {
	return OrdSeq[V](Seq[V](o).SkipWhile(test))
}
func (o OrdSeq[V]) Take(toTake int) OrdSeq[V] {
	return OrdSeq[V](Seq[V](o).Take(toTake))
}
func (o OrdSeq[V]) TakeWhile(test yielder[V]) OrdSeq[V] {
	return OrdSeq[V](Seq[V](o).TakeWhile(test))
}
func (o OrdSeq[V]) Indexed() Seq2[int, V] {
	return Seq[V](o).Indexed()
}
func (o OrdSeq[V]) Expand(toElements mapper[V, Seq[V]]) OrdSeq[V] {
	return OrdSeq[V](Seq[V](o).Expand(toElements))
}
