package edb

import "flag"
import "log"

// appends records to our record input queue
// every now and then, we should pack the input queue into a column, though
func RunDigestCmdLine() {
	digest_file := flag.String("file", "ingest", "Name of block to digest")
	flag.Parse()

	if *f_TABLE == "" {
		flag.PrintDefaults()
		return
	}

	if *f_PROFILE && PROFILER_ENABLED {
		profile := RUN_PROFILER()
		defer profile.Start().Stop()
	}

	t := getTable(*f_TABLE)
	t.LoadRecords(nil)

	log.Println("KEY TABLE", t.KeyTable)

	t.DigestRecords(*digest_file)
}
