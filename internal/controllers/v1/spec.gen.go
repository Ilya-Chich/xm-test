// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYbW/bvhH/KgQ3YP8Cji0ndtb61dIE6Jx2adA2W4EiCGjqZLGRSJWknKiFv/twpORI",
	"EZ2kbdIOw/+NYUu8B9797uHnb5SrvFASpDV09o0ankLO3NeD0qbv4EsJxuLPQqsCtBXgXkLOROa+XLO8",
	"yIDOaGlA/6P+OeQqpwOaKJ0zS2f18QG1VYFHjdVCLul6QAtmzJXScVdV83S8u9fWsjncU7QeUA1fSqEh",
	"prNPG3MbgfONhFp8Bm7RtL+fKZQ00L8g4xyMubDqEiT+jsFwLQorlKQzevyfD8SfIP7EoOU9VMfp4hUX",
	"b8Xx/OzrfHwi5mYu30354Xx/fll8/Pfh8YvhcBgKB1wXQoO5EAGbH9AQcScYPiJW5ECEJAa4krFp+7C3",
	"H0Ub9UJaWIJG/c7ZC/88rN+9a9/mJTANOhjyXkgPNTALhyovmKxa2Okaql+QhYorkihNOIoJuSSMSLgi",
	"3Mtj/ropyVUp7YVKLiAvMlWBf9xVflLmC9BEJWRzCGNkU2jp3dxuGg5TR+VtC0dgmcggJq3HaG+LCXpA",
	"MmAxXs8CT1tHcnb9BuTSpnS2F7UcuYGDZHkgUScsh1sGyR95iREFUkrxpYRnHQ8+/qtrbTwN2NKwFMaC",
	"q6DbFucyFpxZMOQqBZuC7tgWhqgkEVywLKtIS0/LB6tL2BhdKJUBk3QDoR4Uq8LdsBVOWeZY2YdKF8rj",
	"HwF/ouSpVomwdEAPFYKFWbFCAL9XGZBTrQotwCptUlFgF7gJyi1NdzcUl4hBEIGdyNV6Qu3mCDII1caN",
	"R9NpBM8nUbQDuy8WO5NxPNlhfx/v70wm+/vT6WQSRVHUbodlKeJQE3kF9leb2dZFY2ZZP721GIldJZn/",
	"70IfPqjSRaDqzlwpExGDtCIR/rLbjD9WVh/ccv7sMI/aYVw2frrNhIbyKbM8/e6ZXBZxM5PrGP3N/MqC",
	"dQtMLqTIMS3jpy9e6a98p0SwlnMhN7B/yhlOzrx//Uq7x4OnrbuEZeZ/rPDaovduresBNcBLLWz1HsmP",
	"h67fepEi9K/g3xFW2lRp8dV5TAcUd3aaAovdtuzT7khG+1BtnBXiNVR0jdaFTBQascI2mfbVKsAQA3ol",
	"ON56Bdp4++NhNIwwwqoAyQpBZ3RvGA33HOGxqfN/xAoxWo1H6OQoU0tPKAoVqv0zA5pkaunK8UrYlDj+",
	"RJiMScOg3A8NHMQKDGGyw36wLOoEKTmP6Yy+cRZ9hwNjX6q4QrtcSQvSucCKIkPkCSVHn42vYE8+8dtf",
	"NSR0Rv8yumGno5qajtq8dN1tozgD3AO/j7hI7EbRI5uulx1nuxvJ96WLSlJiuWDocXJzV7TrAZ14T24X",
	"4IplAmNb3wjPjUPLQIM3iMkfjRTX4LYDlplnKDoNm7CgJcscmEAT0FppD/wyz5muapw2zhJk8nXCbakl",
	"Qbbb0FzLlgartSWAVs5RXQM63sB3O+YO4th06R6xyvUeUxkLeR9SHW75RNAK8tcHYWyyfcs1bVA4qosZ",
	"PFHk0Lv87KHYcMPZH36x3ZwrYBdIHDHYhgjLNLC4InAtjDU/gpO6QdLZp25r/HS+Pm/DyAewx+Mb0Gz6",
	"2ha8jL6JeO0dQ7LUd/EUdM4wZ24c5co1ow2CEq3yFoZIaXCBEdbUE7W1Svfx1eFnrpVqloMFbdytv281",
	"RzAvgPhrxM1wwO58MxrcxtfF1eCBOA1yyTWm4gdhWTsagOUdGqSyJFGljJ8SUf6qrSQvKjI/2oKpAV1C",
	"cLm1WsAK6v0VYoIzF4kQrnhsoUrbMvBdqLkhwT8LmUcHSf9vgBBCHm84Bv4PuHNEksYT37L4rT8FnhR9",
	"G3i9AttYxHwwYgrgIhH8gXgrkFwF+hTT1i/NnlR0WFQLfENykGXqyrToFvYxAVlM4JpDYR1E5kd95LVp",
	"3W/B3j1/JXisPf6MDtHZRx3RDQ38yRH9uxvnWR94N/T9noGcAss88wn2U78WGrIbRUQkBFagK5u6rmmI",
	"umRVD6z/9ArD3aer/O1rz9JqD70k4SnwS3TShQBjEsJ5DCsCckUHtNQZsjFri9lolCnOslQZO3sePR/T",
	"9eC2nLFsCT1JMxuN3Juh+wzIFVrFQTF8McQPuj5f/zcAAP//LFOhsWQbAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
