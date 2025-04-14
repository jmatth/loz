{{- define "typerange" -}}
{{ range . }}T{{ .I }}{{ if not .IsLast }}, {{ end }}{{ end }}
{{- end }}

{{- define "maptypedef" -}}
Map{{ . }}[{{ template "typerange" add 1 . | numsTo }} any]
{{- end -}}

{{- define "maptype" -}}
Map{{ . }}[{{ template "typerange" add 1 . | numsTo }}]
{{- end -}}

{{- define "prevmaptype" -}}
{{ if eq . 1 }}Seq[T1]{{ else }}{{ template "maptype" add . -1 }}{{ end }}
{{- end -}}

{{- define "prevmapresult" -}}
{{ if eq . 1 }}Seq[T2]{{ else }}Map{{ add . -1 }}[{{ template "typerange" add . 1 | numsTo | skip 1 }}]{{ end }}
{{- end -}}

{{- define "seqderef" }}
func (s {{ template "maptype" . }}) Where(filter yielder[T1]) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[T1](s).Where(filter))
}

func (s {{ template "maptype" . }}) Skip(toSkip int) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[T1](s).Skip(toSkip))
}

func (s {{ template "maptype" . }}) SkipWhile(test yielder[T1]) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[T1](s).SkipWhile(test))
}

func (s {{ template "maptype" . }}) Take(toTake int) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[T1](s).Take(toTake))
}

func (s {{ template "maptype" . }}) TakeWhile(test yielder[T1]) {{ template "maptype" . }} {
	return {{ template "maptype" . }}(Seq[T1](s).TakeWhile(test))
}
{{- end -}}

package {{ .package }}

{{ range numsTo .levels -}}
type {{ template "maptypedef" .I }} {{ template "prevmaptype" .I }}

func (s {{ template "maptype" .I }}) Map(mapper func(T1) T2) {{ template "prevmapresult" .I }} {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

{{ template "seqderef" .I }}

{{ end }}
