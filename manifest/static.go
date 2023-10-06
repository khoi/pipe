package manifest

var staticManifests = []*Manifest{
	&FormatJSON,
	&Base64Encode,
	&Base64Decode,
	&ShuffleLines,
	&UUIDGenerate,
	&JwtDecode,
	&SortLines,
	&LoadFromURL,
	&CurrentTimeStamp,
}

func StaticManifests() []*Manifest {
	return staticManifests
}
