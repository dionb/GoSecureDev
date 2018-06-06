for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    } else {
    	log.Println(err.Error())
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}