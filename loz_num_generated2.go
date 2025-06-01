package loz

func (n NumSeq[V]) Max() (V, error) {
	return OrdSeq[V](n).Max()
}
func (n NumSeq[V]) Min() (V, error) {
	return OrdSeq[V](n).Min()
}
