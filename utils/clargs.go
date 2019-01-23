package utils

import (
	"flag"
)

func ParseArgs() {
	piPtr := flag.String("post_index", POST_INDEX, "Name of the ElasticSearch index for the posts")
	ptPtr := flag.String("post_type", POST_TYPE, "Type of post")
	dPtr := flag.String("distance", DISTANCE, "Distance, radius loaded")
	esPtr := flag.String("es_url", ES_URL, "Full URL of ElasticSearch")
	bPtr := flag.String("bucket_name", BUCKET_NAME, "Bucket Name")
	idPtr := flag.String("project_id", PROJECT_ID, "Project ID on GCP")
	btPtr := flag.String("bigtable_instance", BT_INSTANCE, "Name of BigTable Instance")
	cPtr := flag.String("credential", CREDENTIAL, "path to the credential file")
	flag.Parse()
	POST_INDEX = *piPtr
	POST_TYPE = *ptPtr
	DISTANCE = *dPtr
	ES_URL = *esPtr
	BUCKET_NAME = *bPtr
	PROJECT_ID = *idPtr
	BT_INSTANCE = *btPtr
	CREDENTIAL = *cPtr
	if *btPtr == "" {
		ENABLE_BIGTABLE = false
	} else {
		ENABLE_BIGTABLE = true
	}
}
