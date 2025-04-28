{{- define "typerange" -}}
{{ range . }}V{{ .I }}{{ if not .IsLast }}, {{ end }}{{ end }}
{{- end }}

{{- define "maptypedef" -}}
Map{{ . }}[{{ template "typerange" add 1 . | numsTo }} any]
{{- end -}}

{{- define "maptype" -}}
Map{{ . }}[{{ template "typerange" add 1 . | numsTo }}]
{{- end -}}

{{- define "prevmaptype" -}}
{{ if eq . 1 }}Seq[V1]{{ else }}{{ template "maptype" add . -1 }}{{ end }}
{{- end -}}

{{- define "prevmapresult" -}}
{{ if eq . 1 }}Seq[V2]{{ else }}Map{{ add . -1 }}[{{ template "typerange" add . 1 | numsTo | skip 1 }}]{{ end }}
{{- end -}}

{{- define "seqderef" -}}
// See [Seq.Filter].
func (s {{ template "maptype" . }}) Filter(filter yielder[V1]) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s {{ template "maptype" . }}) Skip(toSkip int) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s {{ template "maptype" . }}) SkipWhile(test yielder[V1]) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s {{ template "maptype" . }}) Take(toTake int) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s {{ template "maptype" . }}) TakeWhile(test yielder[V1]) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[V1](s).TakeWhile(test))
}
{{- end -}}

package {{ .package }}

{{ range numsTo .levels -}}
// A Map{{ .I }} is a wrapper around [Seq] that provides methods to map to {{ .I }} additional type{{ if gt .I 1 }}s{{ end }}.
type {{ template "maptypedef" .I }} {{ template "prevmaptype" .I }}

// Map transforms the elements within the iterator using the provided mapper function.
func (s {{ template "maptype" .I }}) Map(mapper func(V1) V2) {{ template "prevmapresult" .I }} {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s {{ template "maptype" .I }}) Expand(toElements func(V1) Seq[V2]) {{ template "prevmapresult" .I }} {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

{{ template "seqderef" .I }}
{{ end }}
