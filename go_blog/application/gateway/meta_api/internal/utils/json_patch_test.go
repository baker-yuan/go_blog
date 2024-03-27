package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func compareJSON(a, b string) bool {
	var objA, objB interface{}
	json.Unmarshal([]byte(a), &objA)
	json.Unmarshal([]byte(b), &objB)

	return reflect.DeepEqual(objA, objB)
}

func formatJSON(j string) string {
	buf := new(bytes.Buffer)

	json.Indent(buf, []byte(j), "", "  ")

	return buf.String()
}

func TestMergeJson(t *testing.T) {
	cases := []struct {
		doc, patch, result, desc string
	}{
		{
			desc: "simple merge",
			doc: `{
                                "id": "1",
                                "status": 1,
                                "key": "fake key",
                                "cert": "fake cert",
                                "create_time": 1,
                                "update_time": 2
                        }`,
			patch: `{
                                "id": "1",
                                "status": 0,
                                "key": "fake key1",
                                "cert": "fake cert1"
                        }`,
			result: `{
                                "id": "1",
                                "status": 0,
                                "key": "fake key1",
                                "cert": "fake cert1",
                                "create_time": 1,
                                "update_time": 2
                        }`,
		},
		{
			desc: `array merge`,
			doc: `{
                                "uri": "/index.html",
                                "upstream": {
                                        "type": "roundrobin",
                                        "nodes": [{
                                                "host": "39.97.63.215",
                                                "port": 80,
                                                "weight" : 1
                                        }]
                                }
                        }`,
			patch: `{
                                "upstream": {
                                        "nodes": [{
                                                "host": "39.97.63.216",
                                                "port": 80,
                                                "weight" : 1
                                        },{
                                                "host": "39.97.63.217",
                                                "port": 80,
                                                "weight" : 1
                                        }]
                                }
                        }`,
			result: `{
                                "uri": "/index.html",
                                "upstream": {
                                        "type": "roundrobin",
                                        "nodes": [{
                                                "host": "39.97.63.216",
                                                "port": 80,
                                                "weight" : 1
                                        },{
                                                "host": "39.97.63.217",
                                                "port": 80,
                                                "weight" : 1
                                        }]
                                }
                        }`,
		},
	}
	for _, c := range cases {
		out, err := MergeJson([]byte(c.doc), []byte(c.patch))

		if err != nil {
			t.Errorf("Unable to merge patch: %s", err)
		}

		if !compareJSON(string(out), c.result) {
			t.Errorf("Merge failed. Expected:\n%s\n\nActual:\n%s",
				formatJSON(c.result), formatJSON(string(out)))
		}
	}
}

func TestPatchJson(t *testing.T) {
	cases := []struct {
		doc, path, value, result, desc string
	}{
		{
			desc: "patch array",
			doc: `{
                                "uri": "/index.html",
                                "upstream": {
                                        "type": "roundrobin",
                                        "nodes": [{
                                                "host": "39.97.63.215",
                                                "port": 80,
                                                "weight" : 1
                                        }]
                                }
                        }`,
			path: `/upstream/nodes`,
			value: `[{
                                        "host": "39.97.63.216",
                                        "port": 80,
                                        "weight" : 1
                                },{
                                        "host": "39.97.63.217",
                                        "port": 80,
                                        "weight" : 1
                                }]`,
			result: `{
                                "uri": "/index.html",
                                "upstream": {
                                        "type": "roundrobin",
                                        "nodes": [{
                                                "host": "39.97.63.216",
                                                "port": 80,
                                                "weight" : 1
                                        },{
                                                "host": "39.97.63.217",
                                                "port": 80,
                                                "weight" : 1
                                        }]
                                }
                        }`,
		},
		{
			desc: "patch field that non existent",
			doc: `{
                                "uri": "/index.html",
                                "upstream": {
                                        "type": "roundrobin",
                                        "nodes": [{
                                                "host": "39.97.63.215",
                                                "port": 80,
                                                "weight" : 1
                                        }]
                                }
                        }`,
			path:  `/upstream/labels`,
			value: `{"app": "test"}`,
			result: `{
                                "uri": "/index.html",
                                "upstream": {
                                        "type": "roundrobin",
                                        "nodes": [{
                                                "host": "39.97.63.215",
                                                "port": 80,
                                                "weight" : 1
                                        }],
                                        "labels": {"app": "test"}
                                }
                        }`,
		},
	}
	for _, c := range cases {
		out, err := PatchJson([]byte(c.doc), c.path, c.value)
		if err != nil {
			t.Errorf("Unable to patch: %s", err)
		}

		if !compareJSON(string(out), c.result) {
			t.Errorf("Patch failed. Expected:\n%s\n\nActual:\n%s",
				formatJSON(c.result), formatJSON(string(out)))
		}
	}
}
