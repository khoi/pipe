package manifest

var staticManifests = []*Manifest{
	&FormatJSON,
	&Base64Encode,
	&Base64Decode,
	&ShuffleLines,
}

func StaticManifests() []*Manifest {
	return staticManifests
}
