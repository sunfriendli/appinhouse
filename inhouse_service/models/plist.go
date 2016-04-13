// plist
package models

import (
	. "appinhouse/inhouse_service/constants"

	"github.com/DHowett/go-plist"
)

const (
	software_kind      = "software-package"
	full_image_kind    = "full-size-image"
	display_image_kind = "display-image"
	metadata_kind      = "software"
)

type ItemsXml struct {
	Array []interface{} `plist:"items"`
}
type AssetsXml struct {
	Array    []interface{} `plist:"assets"`
	Metadata *MetadataXml  `plist:"metadata"`
}
type SoftwareXml struct {
	Kind string `plist:"kind"`
	Url  string `plist:"url"`
}
type ImageXml struct {
	Kind       string `plist:"kind"`
	NeedsShine bool   `plist:"needs-shine"`
	Url        string `plist:"url"`
}

type MetadataXml struct {
	BundleIdentifier string `plist:"bundle-identifier"`
	BundleVersion    string `plist:"bundle-version"`
	Kind             string `plist:"kind"`
	Title            string `plist:"title"`
}

func NewPlist(id, version, title, softvareUrl, fullUrl, displayUrl string) ([]byte, error) {
	meta := &MetadataXml{
		BundleIdentifier: id,
		BundleVersion:    version,
		Kind:             metadata_kind,
		Title:            title,
	}
	if softvareUrl == "" {
		return nil, ErrorIPANotExistError
	}
	software := SoftwareXml{
		Kind: software_kind,
		Url:  softvareUrl,
	}
	assetlist := make([]interface{}, 0, 3)
	assetlist = append(assetlist, software)
	if fullUrl != "" {
		full := ImageXml{
			Kind:       full_image_kind,
			NeedsShine: true,
			Url:        fullUrl,
		}
		assetlist = append(assetlist, full)
	}
	if displayUrl != "" {
		display := ImageXml{
			Kind:       display_image_kind,
			NeedsShine: true,
			Url:        displayUrl,
		}
		assetlist = append(assetlist, display)
	}

	assets := AssetsXml{
		Array:    assetlist,
		Metadata: meta,
	}
	itemlist := make([]interface{}, 0, 1)
	itemlist = append(itemlist, assets)
	data := &ItemsXml{
		Array: itemlist,
	}

	plist, err := plist.MarshalIndent(data, plist.XMLFormat, "\t")
	if err != nil {
		return nil, ErrorCreateFileError
	}
	return plist, nil
}
