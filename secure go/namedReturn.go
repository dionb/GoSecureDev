package main // OMIT

func ReadFull(r Reader, buf []byte) (n int, err error) {
	var nr int
	for len(buf) > 0 && err == nil {
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
