package config

import (
	"os"
	"reflect"
	"testing"
)

func TestFromFile(t *testing.T) {
	f, err := os.CreateTemp("", "test_config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})

	data := []byte(`features:
  - name: foo
    value: 42
    matchOn:
      - pciId:
          class: $classId
          vendor: $vendorId
          device: $deviceId
  - name: bar 
    value: 21
    matchOn:
      - pciId:
          class: $classId
          vendor: $vendorId
          device: $deviceId
`)
	if err := os.WriteFile(f.Name(), data, 0666); err != nil {
		t.Fatal(err)
	}

	cfg, err := FromFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		Features: []Feature{
			{
				Name:  "foo",
				Value: "42",
				MatchOn: []MatchOn{
					{
						PCIID: PCIID{
							Class:  "$classId",
							Vendor: "$vendorId",
							Device: "$deviceId",
						},
					},
				},
			},
			{
				Name:  "bar",
				Value: "21",
				MatchOn: []MatchOn{
					{
						PCIID: PCIID{
							Class:  "$classId",
							Vendor: "$vendorId",
							Device: "$deviceId",
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, cfg) {
		t.Fatalf("expected:\n%+v\ngot:\n%v\n", expected, cfg)
	}
}
