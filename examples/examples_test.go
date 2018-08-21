package examples

import (
	"bytes"
	"encoding/hex"
	"testing"
	"time"

	"github.com/jcmturner/rpc/mstypes"
	"github.com/jcmturner/rpc/ndr"
	"github.com/stretchr/testify/assert"
)

const (
	PAC_Kerb_Validation_Info_MS = "01100800cccccccca00400000000000000000200d186660f656ac601ffffffffffffff7fffffffffffffff7f17d439fe784ac6011794a328424bc601175424977a81c60108000800040002002400240008000200120012000c0002000000000010000200000000001400020000000000180002005410000097792c00010200001a0000001c000200200000000000000000000000000000000000000016001800200002000a000c002400020028000200000000000000000010000000000000000000000000000000000000000000000000000000000000000d0000002c0002000000000000000000000000000400000000000000040000006c007a00680075001200000000000000120000004c0069007100690061006e00670028004c006100720072007900290020005a00680075000900000000000000090000006e0074006400730032002e0062006100740000000000000000000000000000000000000000000000000000000000000000000000000000001a00000061c433000700000009c32d00070000005eb4320007000000010200000700000097b92c00070000002bf1320007000000ce30330007000000a72e2e00070000002af132000700000098b92c000700000062c4330007000000940133000700000076c4330007000000aefe2d000700000032d22c00070000001608320007000000425b2e00070000005fb4320007000000ca9c35000700000085442d0007000000c2f0320007000000e9ea310007000000ed8e2e0007000000b6eb310007000000ab2e2e0007000000720e2e00070000000c000000000000000b0000004e0054004400450056002d00440043002d003000350000000600000000000000050000004e0054004400450056000000040000000104000000000005150000005951b81766725d2564633b0b0d0000003000020007000000340002000700002038000200070000203c000200070000204000020007000020440002000700002048000200070000204c000200070000205000020007000020540002000700002058000200070000205c00020007000020600002000700002005000000010500000000000515000000b9301b2eb7414c6c8c3b351501020000050000000105000000000005150000005951b81766725d2564633b0b74542f00050000000105000000000005150000005951b81766725d2564633b0be8383200050000000105000000000005150000005951b81766725d2564633b0bcd383200050000000105000000000005150000005951b81766725d2564633b0b5db43200050000000105000000000005150000005951b81766725d2564633b0b41163500050000000105000000000005150000005951b81766725d2564633b0be8ea3100050000000105000000000005150000005951b81766725d2564633b0bc1193200050000000105000000000005150000005951b81766725d2564633b0b29f13200050000000105000000000005150000005951b81766725d2564633b0b0f5f2e00050000000105000000000005150000005951b81766725d2564633b0b2f5b2e00050000000105000000000005150000005951b81766725d2564633b0bef8f3100050000000105000000000005150000005951b81766725d2564633b0b075f2e0000000000"
	PAC_Kerb_Validation_Info    = "01100800cccccccc180200000000000000000200058e4fdd80c6d201ffffffffffffff7fffffffffffffff7fcc27969c39c6d201cce7ffc602c7d201ffffffffffffff7f12001200040002001600160008000200000000000c000200000000001000020000000000140002000000000018000200d80000005104000001020000050000001c000200200000000000000000000000000000000000000008000a002000020008000a00240002002800020000000000000000001002000000000000000000000000000000000000000000000000000000000000020000002c00020000000000000000000000000009000000000000000900000074006500730074007500730065007200310000000b000000000000000b000000540065007300740031002000550073006500720031000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000050000000102000007000000540400000700000055040000070000005b040000070000005c0400000700000005000000000000000400000041004400440043000500000000000000040000005400450053005400040000000104000000000005150000004c86cebca07160e63fdce8870200000030000200070000203400020007000020050000000105000000000005150000004c86cebca07160e63fdce8875a040000050000000105000000000005150000004c86cebca07160e63fdce8875704000000000000"
)

type KerbValidationInfo struct {
	LogOnTime              mstypes.FileTime
	LogOffTime             mstypes.FileTime
	KickOffTime            mstypes.FileTime
	PasswordLastSet        mstypes.FileTime
	PasswordCanChange      mstypes.FileTime
	PasswordMustChange     mstypes.FileTime
	EffectiveName          mstypes.RPCUnicodeString
	FullName               mstypes.RPCUnicodeString
	LogonScript            mstypes.RPCUnicodeString
	ProfilePath            mstypes.RPCUnicodeString
	HomeDirectory          mstypes.RPCUnicodeString
	HomeDirectoryDrive     mstypes.RPCUnicodeString
	LogonCount             uint16
	BadPasswordCount       uint16
	UserID                 uint32
	PrimaryGroupID         uint32
	GroupCount             uint32
	GroupIDs               []mstypes.GroupMembership `ndr:"pointer,conformant"`
	UserFlags              uint32
	UserSessionKey         mstypes.UserSessionKey
	LogonServer            mstypes.RPCUnicodeString
	LogonDomainName        mstypes.RPCUnicodeString
	LogonDomainID          mstypes.RPCSID `ndr:"pointer"`
	Reserved1              [2]uint32      // Has 2 elements
	UserAccountControl     uint32
	SubAuthStatus          uint32
	LastSuccessfulILogon   mstypes.FileTime
	LastFailedILogon       mstypes.FileTime
	FailedILogonCount      uint32
	Reserved3              uint32
	SIDCount               uint32
	ExtraSIDs              []mstypes.KerbSidAndAttributes `ndr:"pointer,conformant"`
	ResourceGroupDomainSID mstypes.RPCSID                 `ndr:"pointer"`
	ResourceGroupCount     uint32
	ResourceGroupIDs       []mstypes.GroupMembership `ndr:"pointer,conformant"`
}

func Test_KerbValidationInfo(t *testing.T) {
	b, _ := hex.DecodeString(PAC_Kerb_Validation_Info_MS)
	k := new(KerbValidationInfo)
	dec := ndr.NewDecoder(bytes.NewReader(b))
	err := dec.Decode(k)
	if err != nil {
		t.Errorf("%v", err)
	}
	assert.Equal(t, time.Date(2006, 4, 28, 1, 42, 50, 925640100, time.UTC), k.LogOnTime.Time(), "LogOnTime not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551516, time.UTC), k.LogOffTime.Time(), "LogOffTime not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551516, time.UTC), k.KickOffTime.Time(), "KickOffTime not as expected")
	assert.Equal(t, time.Date(2006, 3, 18, 10, 44, 54, 837147900, time.UTC), k.PasswordLastSet.Time(), "PasswordLastSet not as expected")
	assert.Equal(t, time.Date(2006, 3, 19, 10, 44, 54, 837147900, time.UTC), k.PasswordCanChange.Time(), "PasswordCanChange not as expected")

	assert.Equal(t, "lzhu", k.EffectiveName.String(), "EffectiveName not as expected")
	assert.Equal(t, "Liqiang(Larry) Zhu", k.FullName.String(), "EffectiveName not as expected")
	assert.Equal(t, "ntds2.bat", k.LogonScript.String(), "EffectiveName not as expected")
	assert.Equal(t, "", k.ProfilePath.String(), "EffectiveName not as expected")
	assert.Equal(t, "", k.HomeDirectory.String(), "EffectiveName not as expected")
	assert.Equal(t, "", k.HomeDirectoryDrive.String(), "EffectiveName not as expected")

	assert.Equal(t, uint16(4180), k.LogonCount, "LogonCount not as expected")
	assert.Equal(t, uint16(0), k.BadPasswordCount, "BadPasswordCount not as expected")
	assert.Equal(t, uint32(2914711), k.UserID, "UserID not as expected")
	assert.Equal(t, uint32(513), k.PrimaryGroupID, "PrimaryGroupID not as expected")
	assert.Equal(t, uint32(26), k.GroupCount, "GroupCount not as expected")

	gids := []mstypes.GroupMembership{
		{RelativeID: 3392609, Attributes: 7},
		{RelativeID: 2999049, Attributes: 7},
		{RelativeID: 3322974, Attributes: 7},
		{RelativeID: 513, Attributes: 7},
		{RelativeID: 2931095, Attributes: 7},
		{RelativeID: 3338539, Attributes: 7},
		{RelativeID: 3354830, Attributes: 7},
		{RelativeID: 3026599, Attributes: 7},
		{RelativeID: 3338538, Attributes: 7},
		{RelativeID: 2931096, Attributes: 7},
		{RelativeID: 3392610, Attributes: 7},
		{RelativeID: 3342740, Attributes: 7},
		{RelativeID: 3392630, Attributes: 7},
		{RelativeID: 3014318, Attributes: 7},
		{RelativeID: 2937394, Attributes: 7},
		{RelativeID: 3278870, Attributes: 7},
		{RelativeID: 3038018, Attributes: 7},
		{RelativeID: 3322975, Attributes: 7},
		{RelativeID: 3513546, Attributes: 7},
		{RelativeID: 2966661, Attributes: 7},
		{RelativeID: 3338434, Attributes: 7},
		{RelativeID: 3271401, Attributes: 7},
		{RelativeID: 3051245, Attributes: 7},
		{RelativeID: 3271606, Attributes: 7},
		{RelativeID: 3026603, Attributes: 7},
		{RelativeID: 3018354, Attributes: 7},
	}
	assert.Equal(t, gids, k.GroupIDs, "GroupIDs not as expected")

	assert.Equal(t, uint32(32), k.UserFlags, "UserFlags not as expected")

	assert.Equal(t, mstypes.UserSessionKey{CypherBlock: [2]mstypes.CypherBlock{{Data: [8]byte{}}, {Data: [8]byte{}}}}, k.UserSessionKey, "UserSessionKey not as expected")

	assert.Equal(t, "NTDEV-DC-05", k.LogonServer.Value, "LogonServer not as expected")
	assert.Equal(t, "NTDEV", k.LogonDomainName.Value, "LogonDomainName not as expected")

	assert.Equal(t, "S-1-5-21-397955417-626881126-188441444", k.LogonDomainID.String(), "LogonDomainID not as expected")

	assert.Equal(t, uint32(16), k.UserAccountControl, "UserAccountControl not as expected")
	assert.Equal(t, uint32(0), k.SubAuthStatus, "SubAuthStatus not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551616, time.UTC), k.LastSuccessfulILogon.Time(), "LastSuccessfulILogon not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551616, time.UTC), k.LastFailedILogon.Time(), "LastSuccessfulILogon not as expected")
	assert.Equal(t, uint32(0), k.FailedILogonCount, "FailedILogonCount not as expected")

	assert.Equal(t, uint32(13), k.SIDCount, "SIDCount not as expected")
	assert.Equal(t, int(k.SIDCount), len(k.ExtraSIDs), "SIDCount and size of ExtraSIDs list are not the same")

	var es = []struct {
		sid  string
		attr uint32
	}{
		{"S-1-5-21-773533881-1816936887-355810188-513", uint32(7)},
		{"S-1-5-21-397955417-626881126-188441444-3101812", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3291368", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3291341", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3322973", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3479105", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3271400", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3283393", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3338537", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3038991", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3037999", uint32(536870919)},
		{"S-1-5-21-397955417-626881126-188441444-3248111", uint32(536870919)},
	}
	for i, s := range es {
		assert.Equal(t, s.sid, k.ExtraSIDs[i].SID.String(), "ExtraSID SID value not as epxected")
		assert.Equal(t, s.attr, k.ExtraSIDs[i].Attributes, "ExtraSID Attributes value not as epxected")
	}

	assert.Equal(t, uint8(0), k.ResourceGroupDomainSID.SubAuthorityCount, "ResourceGroupDomainSID not as expected")
	assert.Equal(t, 0, len(k.ResourceGroupIDs), "ResourceGroupIDs not as expected")

	b, _ = hex.DecodeString(PAC_Kerb_Validation_Info)
	k2 := new(KerbValidationInfo)
	dec = ndr.NewDecoder(bytes.NewReader(b))
	err = dec.Decode(k2)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, time.Date(2017, 5, 6, 15, 53, 11, 825766900, time.UTC), k2.LogOnTime.Time(), "LogOnTime not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551516, time.UTC), k2.LogOffTime.Time(), "LogOffTime not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551516, time.UTC), k2.KickOffTime.Time(), "KickOffTime not as expected")
	assert.Equal(t, time.Date(2017, 5, 6, 7, 23, 8, 968750000, time.UTC), k2.PasswordLastSet.Time(), "PasswordLastSet not as expected")
	assert.Equal(t, time.Date(2017, 5, 7, 7, 23, 8, 968750000, time.UTC), k2.PasswordCanChange.Time(), "PasswordCanChange not as expected")

	assert.Equal(t, "testuser1", k2.EffectiveName.String(), "EffectiveName not as expected")
	assert.Equal(t, "Test1 User1", k2.FullName.String(), "EffectiveName not as expected")
	assert.Equal(t, "", k2.LogonScript.String(), "EffectiveName not as expected")
	assert.Equal(t, "", k2.ProfilePath.String(), "EffectiveName not as expected")
	assert.Equal(t, "", k2.HomeDirectory.String(), "EffectiveName not as expected")
	assert.Equal(t, "", k2.HomeDirectoryDrive.String(), "EffectiveName not as expected")

	assert.Equal(t, uint16(216), k2.LogonCount, "LogonCount not as expected")
	assert.Equal(t, uint16(0), k2.BadPasswordCount, "BadPasswordCount not as expected")
	assert.Equal(t, uint32(1105), k2.UserID, "UserID not as expected")
	assert.Equal(t, uint32(513), k2.PrimaryGroupID, "PrimaryGroupID not as expected")
	assert.Equal(t, uint32(5), k2.GroupCount, "GroupCount not as expected")

	gids = []mstypes.GroupMembership{
		{RelativeID: 513, Attributes: 7},
		{RelativeID: 1108, Attributes: 7},
		{RelativeID: 1109, Attributes: 7},
		{RelativeID: 1115, Attributes: 7},
		{RelativeID: 1116, Attributes: 7},
	}
	assert.Equal(t, gids, k2.GroupIDs, "GroupIDs not as expected")

	assert.Equal(t, uint32(32), k2.UserFlags, "UserFlags not as expected")

	assert.Equal(t, mstypes.UserSessionKey{CypherBlock: [2]mstypes.CypherBlock{{Data: [8]byte{}}, {Data: [8]byte{}}}}, k2.UserSessionKey, "UserSessionKey not as expected")

	assert.Equal(t, "ADDC", k2.LogonServer.String(), "LogonServer not as expected")
	assert.Equal(t, "TEST", k2.LogonDomainName.String(), "LogonDomainName not as expected")

	assert.Equal(t, "S-1-5-21-3167651404-3865080224-2280184895", k2.LogonDomainID.String(), "LogonDomainID not as expected")

	assert.Equal(t, uint32(528), k2.UserAccountControl, "UserAccountControl not as expected")
	assert.Equal(t, uint32(0), k2.SubAuthStatus, "SubAuthStatus not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551616, time.UTC), k2.LastSuccessfulILogon.Time(), "LastSuccessfulILogon not as expected")
	assert.Equal(t, time.Date(2185, 7, 21, 23, 34, 33, 709551616, time.UTC), k2.LastFailedILogon.Time(), "LastSuccessfulILogon not as expected")
	assert.Equal(t, uint32(0), k2.FailedILogonCount, "FailedILogonCount not as expected")

	assert.Equal(t, uint32(2), k2.SIDCount, "SIDCount not as expected")
	assert.Equal(t, int(k2.SIDCount), len(k2.ExtraSIDs), "SIDCount and size of ExtraSIDs list are not the same")

	var es2 = []struct {
		sid  string
		attr uint32
	}{
		{"S-1-5-21-3167651404-3865080224-2280184895-1114", uint32(536870919)},
		{"S-1-5-21-3167651404-3865080224-2280184895-1111", uint32(536870919)},
	}
	for i, s := range es2 {
		assert.Equal(t, s.sid, k2.ExtraSIDs[i].SID.String(), "ExtraSID SID value not as expected")
		assert.Equal(t, s.attr, k2.ExtraSIDs[i].Attributes, "ExtraSID Attributes value not as expected")
	}

	assert.Equal(t, uint8(0), k2.ResourceGroupDomainSID.SubAuthorityCount, "ResourceGroupDomainSID not as expected")
	assert.Equal(t, 0, len(k2.ResourceGroupIDs), "ResourceGroupIDs not as expected")
}