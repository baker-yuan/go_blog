package utils

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

func MergeJson(doc, patch []byte) ([]byte, error) {
	out, err := jsonpatch.MergePatch(doc, patch)

	if err != nil {
		return nil, err
	}
	return out, nil
}

func PatchJson(doc []byte, path, val string) ([]byte, error) {
	patch := []byte(`[ { "op": "replace", "path": "` + path + `", "value": ` + val + `}]`)
	obj, err := jsonpatch.DecodePatch(patch)
	if err != nil {
		return nil, err
	}

	out, err := obj.Apply(doc)

	if err != nil {
		// try to add if field not exist
		patch = []byte(`[ { "op": "add", "path": "` + path + `", "value": ` + val + `}]`)
		obj, err = jsonpatch.DecodePatch(patch)
		if err != nil {
			return nil, err
		}
		out, err = obj.Apply(doc)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func MergePatch(obj interface{}, subPath string, reqBody []byte) ([]byte, error) {
	var res []byte
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return res, err
	}

	if subPath != "" {
		res, err = PatchJson(jsonBytes, subPath, string(reqBody))
	} else {
		res, err = MergeJson(jsonBytes, reqBody)
	}

	if err != nil {
		return res, err
	}
	return res, nil
}
