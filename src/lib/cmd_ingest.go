package edb

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// appends records to our record input queue
// every now and then, we should pack the input queue into a column, though
func RunIngestCmdLine() {
	ingestfile := flag.String("file", "ingest", "name of dir to ingest into")

	flag.Parse()

	digestfile := fmt.Sprintf("%s", *ingestfile)

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

	dec := json.NewDecoder(os.Stdin)
	count := 0
	for {
		var recordmap map[string]interface{}

		if err := dec.Decode(&recordmap); err != nil {
			if err == io.EOF {
				break
			}

			log.Println("ERR:", err)

			continue
		}

		r := t.NewRecord()

		intm := recordmap["ints"].(map[string]interface{})

		for k, v := range intm {
			switch iv := v.(type) {
			case float64:
				r.AddIntField(k, int(iv))

			}
		}

		strm := recordmap["strs"].(map[string]interface{})
		for k, v := range strm {
			switch iv := v.(type) {
			case string:
				r.AddStrField(k, iv)
			}
		}

		count++

		if count >= CHUNK_SIZE {
			count -= CHUNK_SIZE

			t.IngestRecords(digestfile)
		}

	}

	t.IngestRecords(digestfile)
}
