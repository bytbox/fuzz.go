package fuzz

var format = "{Hour|15{Minute|:04{Second|:05}}{Zone| MST}, }{Month|January }{Day|2}{Year|{Day|, }2006}"

const (
	cs_blank = iota
	cs_field
)

type compilerState struct {
	state int
	buf   string
	fname string
	prev  *compilerState
}

func compileFormat(f string, enabled []string) string {
	isEnabled := func(str string) bool {
		for _, s := range enabled {
			if s == str {
				return true
			}
		}
		return false
	}

	cState := &compilerState{}
	state := cs_blank
	fname := ""
	res := "" // TODO use an actual buffer
	for _, c := range f {
		switch state {
		case cs_blank:
			if c == '{' {
				state = cs_field
			} else if c == '}' {
				// pop a state off the stack
				if isEnabled(cState.fname) {
					// append to cState.buf
					res = cState.buf + res
				} else {
					res = cState.buf
				}
				cState = cState.prev
			} else {
				// append this plain-text to the result
				res += string(c)
			}
		case cs_field:
			if c == '|' {
				// push this state and start over
				newState := &compilerState {
					state: state,
					buf: res,
					fname: fname,
					prev: cState,
				}
				cState = newState
				state = cs_blank
				fname = ""
				res = ""
			} else {
				fname += string(c)
			}
		default:
			panic("Bad state")
		}
	}
	return res
}
