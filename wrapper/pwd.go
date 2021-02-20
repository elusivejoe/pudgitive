package wrapper

func (w *Wrapper) Pwd() string {
	if len(w.where) > 0 {
		return w.where
	}

	return "/"
}
