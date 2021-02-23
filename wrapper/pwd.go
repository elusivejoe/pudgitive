package wrapper

func (w *Wrapper) Pwd() string {
	if len(w.pwd) > 0 {
		return w.pwd
	}

	return "/"
}
