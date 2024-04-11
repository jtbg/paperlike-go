package dasung

import (
	"testing"
)

func TestSetContrast(t *testing.T) {
	dc, err := NewDasungControl("/dev/i2c-1")
	if err != nil {
		t.Fatalf("Failed to create DasungControl: %v", err)
	}

	err = dc.SetContrast(5)
	if err != nil {
		t.Errorf("Failed to set contrast: %v", err)
	}
}
