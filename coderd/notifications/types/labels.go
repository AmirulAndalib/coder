package types

import "fmt"

// Labels represents the metadata defined in a notification message, which will be used to augment the notification
// display and delivery.
type Labels map[string]string

func (l Labels) GetStrict(k string) (string, bool) {
	v, ok := l[k]
	return v, ok
}

func (l Labels) Get(k string) string {
	return l[k]
}

func (l Labels) Set(k, v string) {
	l[k] = v
}

func (l Labels) SetValue(k string, v fmt.Stringer) {
	l[k] = v.String()
}

// Merge combines two Labels. Keys declared on the given Labels will win over the existing Labels.
func (l Labels) Merge(m Labels) {
	if len(m) == 0 {
		return
	}

	for k, v := range m {
		l[k] = v
	}
}

func (l Labels) Delete(k string) {
	delete(l, k)
}

func (l Labels) Contains(ks ...string) bool {
	for _, k := range ks {
		if _, has := l[k]; !has {
			return false
		}
	}

	return true
}
