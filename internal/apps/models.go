package apps

type AppResult struct {
	Name string
	Path string
}

type PlistRoot struct {
	CFBundleExecutable         string              `plist:"CFBundleExecutable"`
	CFBundleShortVersionString string              `plist:"CFBundleShortVersionString"`
	NSHumanReadableCopyright   string              `plist:"NSHumanReadableCopyright"`
	CFBundleIdentifier         string              `plist:"CFBundleIdentifier"`
	CFBundleSupportedPlatforms []string            `plist:"CFBundleSupportedPlatforms"`
	CFBundleDocumentTypes      []PlistDocumentType `plist:"CFBundleDocumentTypes"`
}

type PlistDocumentType struct {
	CFBundleTypeExtensions []string `plist:"CFBundleTypeExtensions"`
	CFBundleTypeRole       string   `plist:"CFBundleTypeRole"`
	LSHandlerRank          string   `plist:"LSHandlerRank"`
}
