package manifest

var staticManifests = []*Manifest{
	&FormatJSON,
	&Base64Encode,
	&Base64Decode,
}

func StaticManifests() []*Manifest {
	return staticManifests
}
