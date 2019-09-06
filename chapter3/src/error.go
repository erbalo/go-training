package src

func Balance() (int, error) {
	return 1, nil
}

func errorCheck1() error {
	_, err := Balance()
	if err != nil {
		return err
	}
	return nil
}

func errorCheck2() error {
	if _, err := Balance(); err != nil {
		return err
	}
	return nil
}

func Write(msg string) error {
	return nil
}

func errorCheck3() error {
	err := Write("1")
	if err != nil {
		return err
	}
	err = Write("2")
	if err != nil {
		return err
	}
	err = Write("3")
	if err != nil {
		return err
	}
	return nil
}

type errWriter struct {
	err error
}

func (ew *errWriter) Write(msg string) {
	if ew.err != nil {
		return
	}
	ew.err = Write(msg)
}

func errorCheck4() error {
	var wr errWriter
	wr.Write("1")
	wr.Write("2")
	wr.Write("3")
	if wr.err != nil {
		return wr.err
	}
	return nil
}
