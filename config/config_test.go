package config

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/xl1605368195/crypt/backend/mock"
)

var pubring = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG/MacGPG2 v2.0.22 (Darwin)
Comment: GPGTools - https://gpgtools.org

mQENBFRClxIBCACjlm8e2mI5TmeigPuF4HJqNxc6AFLoCsE3MQ6VtdEVqidXZ06L
m7uIXYc3IfvAlID+1KrUJnO2IgLEMmXKSDI5aOch7VaeoXLKMt7yAX+N6DHaZl4l
eUGlfyIuBGD7FY2rv4hHo2wOmlq/chnNA4T7wb2XzeaAjvvoxcedMZ2npVimjwsl
MNDxSxYPlHR6lJgfYJHAxcWn7ZQJW2Kllv9jMQwzGqW4fxuKRhe20KStE/4+K9gL
GWv6OoE2gcGLoXliIMchHobY0GEvVx+YUv5jAItRSXq4ajYjFLtsWLz6FYtK9CoO
Va6T5EGqozKST/olW/FMmKLOTzpAilyoKB/HABEBAAG0LWFwcCAoYXBwIGNvbmZp
Z3VyYXRpb24ga2V5KSA8YXBwQGV4YW1wbGUuY29tPokBNwQTAQoAIQUCVEKXEgIb
AwULCQgHAwUVCgkICwUWAgMBAAIeAQIXgAAKCRA8TymBhIANsjB1CACi4kqqWNSq
AID7LmMswh5FQDEPkI/WA0h75xead11FVSdvtjWANY4Wob8RBjeZNT0TaCa0IAoo
k+tLqA5xNbbvalOPV2zfr86BcGMhIs900++PuVjOb7XaJPsEt5JwtzuLM+eDLIVh
vMI7hQtgB39O8/AsWEW/E/JlVtHcrsQ7LfcQYmNZVSnL71a8w4G+A6Sto89fvpjY
h9/M4+aHqMhO/NLLp8Ylj5TlyiWKHZlx5ufl2ejWMUot3wFhYADHPkhydmQV9IY1
zzIpmB/75kvZqC4p92k7l8Ra82o+T75/dNy0HcgvgrfZQttxIM0WPEyVF5NjicSo
akoggAAslhCNuQENBFRClxIBCADJltx4EgkFScH/EAmO6+mZb6+pcpjY/H97bX4w
KUrQSDZjDAhoxsInKgqHwAo3QY261eYrAyHvoTA2kRAaVrYWeGu3RxMmX5LTjFsX
IW44ocTJK1XziUQympgIEayOUHt+XJaMGL8RKXvNgttGkr2VPD0IWJCOaBr8ZxUG
Fm/pRFeBe6tX02RVKx4QFPqCnb76bkvR1cNeFsV5eEz0WNRYzena+lD6Oqh074tk
oC9Uwl7D0l7xq17HNqAqHdMIO/T/TMPYyb7vskHPL9g8EJSgU55Z2Cjx3IlbJCpA
204cbbak4h99kgAqb4O5jT3gDe03NzWXCeQVKmacahusqNxzABEBAAGJAR8EGAEK
AAkFAlRClxICGwwACgkQPE8pgYSADbJFTwf/d6XIv8/BxaP5TB12FxRXevOVenYe
xR6O0rXDKFkP54LHSh2uaJbjLHevF0tuiNLFWq9RPhVjC012OLdo+2ygEFwNfikd
1SMbUIkuZ6Nu2MvCdrpAwbcvLgeoe8bqf1B6EIb31/OxCmtHujpVw6cSAnpAVyYo
PjPtEpcNatIHbOna4KANxwv1Rmc6XDxWIB0RIlbdZDUhEdLovLLWGjm4J++Cnn2n
OFdZyyUxwYTjDCMWwsYrG2oPZ0Yle6fKEXX30E+dN9NSV1i+dJAYQi0am6augpg+
LmFWxQ6JPmUJVDay9wo6g2D4KbJQybSh8lmqpenHnKD1m/gCGadPmMl6Rw==
=FKbO
-----END PGP PUBLIC KEY BLOCK-----`

var secring = `
-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG/MacGPG2 v2.0.22 (Darwin)
Comment: GPGTools - https://gpgtools.org

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
Rw==
=RvPL
-----END PGP PRIVATE KEY BLOCK-----
`

func Test_StandardSet_BasePath(t *testing.T) {
	key := "foo"
	value := []byte("bar")

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewStandardConfigManager(store)
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	err = cm.Set(key, value)
	if err != nil {
		t.Errorf("Error adding key: %s\n", err.Error())
	}
}

func Test_StandardGet_BasePath(t *testing.T) {
	key := "foo"
	value := []byte("bar")

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewStandardConfigManager(store)
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	storedValue, err := cm.Get(key)
	if err != nil {
		t.Errorf("Error getting key: %s\n", err.Error())
	}
	if !reflect.DeepEqual(storedValue, value) {
		t.Errorf("Two values did not match: %s\n", err.Error())
	}
}

func Test_StandardGet_AlternatePath_NoKey(t *testing.T) {
	key := "doesnotexist"

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewStandardConfigManager(store)
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	_, err = cm.Get(key)
	if err == nil {
		t.Errorf("Did not get expected error\n")
	}
}

func Test_StandardList_BasePath(t *testing.T) {
	dir := "dir"
	key := "dir/foo"
	value := []byte("bar")

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewStandardConfigManager(store)
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	err = cm.Set(key, value)
	if err != nil {
		t.Errorf("Error adding key: %s\n", err.Error())
	}
	storedValue, err := cm.List(dir)
	if err != nil {
		t.Errorf("Error list dir: %s\n", err.Error())
	}
	if len(storedValue) == 0 {
		t.Error("Empty list got on list directory")
	}
	if storedValue[0].Key != key {
		t.Errorf("Keys did not match: %s vs %s", storedValue[0].Key, key)
	}
	if !reflect.DeepEqual(storedValue[0].Value, value) {
		t.Errorf("Two values did not match: %+v vs %+v", storedValue[0].Value, value)
	}
}

func Test_StandardWatch_BasePath(t *testing.T) {
	key := "foo"

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewStandardConfigManager(store)
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	timeout := make(chan bool, 0)
	resp := cm.Watch(key, timeout)
	select {
	case r := <-resp:
		if r.Error != nil {
			t.Errorf("Error watching value: %s\n", r.Error.Error())
		}
	}
}

func Test_Set_BasePath(t *testing.T) {
	key := "foo_enc"
	value := []byte("bar_enc")

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewConfigManager(store, bytes.NewBufferString(pubring))
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	err = cm.Set(key, value)
	if err != nil {
		t.Errorf("Error adding key: %s\n", err.Error())
	}
}

func Test_Get_BasePath(t *testing.T) {
	key := "foo_enc"
	value := []byte("bar_enc")

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewConfigManager(store, bytes.NewBufferString(secring))
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	storedValue, err := cm.Get(key)
	if err != nil {
		t.Errorf("Error getting key: %s\n", err.Error())
	}
	if !reflect.DeepEqual(storedValue, value) {
		t.Errorf("Two values did not match: %s\n", err.Error())
	}
}

func Test_Get_AlternatePath_NoKey(t *testing.T) {
	key := "doesnotexist_enc"

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewConfigManager(store, bytes.NewBufferString(secring))
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	_, err = cm.Get(key)
	if err == nil {
		t.Errorf("Did not get expected error\n")
	}
}

func Test_List_BasePath(t *testing.T) {
	dir := "dir_enc"
	key := "dir_enc/foo_enc"
	value := []byte("bar_enc")

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewConfigManager(store, bytes.NewBufferString(secring))
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	err = cm.Set(key, value)
	if err != nil {
		t.Errorf("Error adding key: %s\n", err.Error())
	}
	storedValue, err := cm.List(dir)
	if err != nil {
		t.Errorf("Error list dir: %s\n", err.Error())
	}
	if len(storedValue) == 0 {
		t.Error("Empty list got on list directory")
	}
	if storedValue[0].Key != key {
		t.Errorf("Keys did not match: %s vs %s", storedValue[0].Key, key)
	}
	if !reflect.DeepEqual(storedValue[0].Value, value) {
		t.Errorf("Two values did not match: %+v vs %+v", storedValue[0].Value, value)
	}
}

func Test_Watch_BasePath(t *testing.T) {
	key := "foo_enc"

	store, err := mock.New([]string{})
	if err != nil {
		t.Errorf("Error creating backend: %s\n", err.Error())
	}
	cm, err := NewConfigManager(store, bytes.NewBufferString(secring))
	if err != nil {
		t.Errorf("Error creating config manager: %s\n", err.Error())
	}
	timeout := make(chan bool, 0)
	resp := cm.Watch(key, timeout)
	select {
	case r := <-resp:
		if r.Error != nil {
			t.Errorf("Error watching value: %s\n", r.Error.Error())
		}
	}
}
