package fuzz

var format = "15:04:05 MST, January 2{Year|{Day|, }2006}"

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

	cState := compilerState{}
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
				oldState := cState
				cState = *oldState.prev
				if isEnabled(oldState.fname) {
					// append to oldState.buf
					res = oldState.buf + res
				} else {
					res = oldState.buf
				}
			} else {
				// append this plain-text to the result
				res += string(c)
			}
		case cs_field:
			if c == '|' {
				// push this state and start over
				newState := compilerState {
					state: state,
					buf: res,
					fname: fname,
					prev: &cState,
				}
				cState = newState
				state = cs_blank
				fname = ""
				res=""
			} else {
				fname += string(c)
			}
		default:
			panic("Bad state")
		}
	}
	return res
}
