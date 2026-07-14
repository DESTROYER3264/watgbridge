package whatsapp

import (
	"testing"

	"go.mau.fi/whatsmeow/proto/waCompanionReg"
	"go.mau.fi/whatsmeow/proto/waWa6"
	"go.mau.fi/whatsmeow/store"
)

func TestUseAndroidClientIdentity(t *testing.T) {
	originalPlatform := store.BaseClientPayload.UserAgent.Platform
	originalWebInfo := store.BaseClientPayload.WebInfo
	originalPlatformType := store.DeviceProps.PlatformType
	originalOS := store.DeviceProps.Os
	t.Cleanup(func() {
		store.BaseClientPayload.UserAgent.Platform = originalPlatform
		store.BaseClientPayload.WebInfo = originalWebInfo
		store.DeviceProps.PlatformType = originalPlatformType
		store.DeviceProps.Os = originalOS
	})

	useAndroidClientIdentity(false)

	if got := store.BaseClientPayload.GetUserAgent().GetPlatform(); got != waWa6.ClientPayload_UserAgent_ANDROID {
		t.Fatalf("expected Android user agent platform, got %s", got)
	}
	if store.BaseClientPayload.GetWebInfo() != nil {
		t.Fatal("expected web info to be nil")
	}
	if got := store.DeviceProps.GetPlatformType(); got != waCompanionReg.DeviceProps_ANDROID_PHONE {
		t.Fatalf("expected Android phone platform type, got %s", got)
	}
	if got := store.DeviceProps.GetOs(); got != "Android" {
		t.Fatalf("expected Android OS, got %q", got)
	}
}

func TestUseAndroidClientIdentityBusiness(t *testing.T) {
	originalPlatform := store.BaseClientPayload.UserAgent.Platform
	t.Cleanup(func() {
		store.BaseClientPayload.UserAgent.Platform = originalPlatform
	})

	useAndroidClientIdentity(true)

	if got := store.BaseClientPayload.GetUserAgent().GetPlatform(); got != waWa6.ClientPayload_UserAgent_SMB_ANDROID {
		t.Fatalf("expected SMB Android user agent platform, got %s", got)
	}
}
