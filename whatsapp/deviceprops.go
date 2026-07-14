package whatsapp

import (
	"go.mau.fi/whatsmeow/proto/waCompanionReg"
	"go.mau.fi/whatsmeow/proto/waWa6"
	"go.mau.fi/whatsmeow/store"
	"google.golang.org/protobuf/proto"
)

func useAndroidClientIdentity(business bool) {
	platform := waWa6.ClientPayload_UserAgent_ANDROID
	if business {
		platform = waWa6.ClientPayload_UserAgent_SMB_ANDROID
	}

	store.BaseClientPayload.UserAgent.Platform = platform.Enum()
	store.BaseClientPayload.WebInfo = nil
	store.DeviceProps.PlatformType = waCompanionReg.DeviceProps_ANDROID_PHONE.Enum()
	store.DeviceProps.Os = proto.String("Android")
}
