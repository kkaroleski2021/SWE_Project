package entity

//this file will hold entities.... structs of JSON info used in GET and POST requests

//may catagorize files based on subject matter in the future if necessary. otherwise just one file should be fine

type Search struct {
	Search string "json:'search'"
}
