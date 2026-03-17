// ==============================================================================
// MasterDnsVPN
// Author: MasterkinG32
// Github: https://github.com/masterking32
// Year: 2026
// ==============================================================================

package enums

import "testing"

func TestPacketEnumValuesAreStable(t *testing.T) {
	if PacketSessionInit != 0x05 {
		t.Fatalf("unexpected PacketSessionInit value: got=%#x want=%#x", PacketSessionInit, 0x05)
	}
	if PacketStreamData != 0x0D {
		t.Fatalf("unexpected PacketStreamData value: got=%#x want=%#x", PacketStreamData, 0x0D)
	}
	if PacketDNSQueryReq != 0x31 {
		t.Fatalf("unexpected PacketDNSQueryReq value: got=%#x want=%#x", PacketDNSQueryReq, 0x31)
	}
	if PacketErrorDrop != 0xFF {
		t.Fatalf("unexpected PacketErrorDrop value: got=%#x want=%#x", PacketErrorDrop, 0xFF)
	}
}

func TestPacketEnumValuesAreUnique(t *testing.T) {
	values := []int{
		PacketMTUUpReq,
		PacketMTUUpRes,
		PacketMTUDownReq,
		PacketMTUDownRes,
		PacketSessionInit,
		PacketSessionAccept,
		PacketSetMTUReq,
		PacketSetMTURes,
		PacketPing,
		PacketPong,
		PacketStreamSyn,
		PacketStreamSynAck,
		PacketStreamData,
		PacketStreamDataAck,
		PacketStreamResend,
		PacketPackedControlBlocks,
		PacketStreamFin,
		PacketStreamFinAck,
		PacketStreamRST,
		PacketStreamRSTAck,
		PacketStreamKeepalive,
		PacketStreamKeepaliveAck,
		PacketStreamWindowUpdate,
		PacketStreamWindowUpdateAck,
		PacketStreamProbe,
		PacketStreamProbeAck,
		PacketSocks5Syn,
		PacketSocks5SynAck,
		PacketSocks5ConnectFail,
		PacketSocks5ConnectFailAck,
		PacketSocks5RulesetDenied,
		PacketSocks5RulesetDeniedAck,
		PacketSocks5NetworkUnreachable,
		PacketSocks5NetworkUnreachableAck,
		PacketSocks5HostUnreachable,
		PacketSocks5HostUnreachableAck,
		PacketSocks5ConnectionRefused,
		PacketSocks5ConnectionRefusedAck,
		PacketSocks5TTLExpired,
		PacketSocks5TTLExpiredAck,
		PacketSocks5CommandUnsupported,
		PacketSocks5CommandUnsupportedAck,
		PacketSocks5AddressTypeUnsupported,
		PacketSocks5AddressTypeUnsupportedAck,
		PacketSocks5AuthFailed,
		PacketSocks5AuthFailedAck,
		PacketSocks5UpstreamUnavailable,
		PacketSocks5UpstreamUnavailableAck,
		PacketDNSQueryReq,
		PacketDNSQueryRes,
		PacketErrorDrop,
	}

	seen := make(map[int]struct{}, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			t.Fatalf("duplicate packet enum value detected: %#x", value)
		}
		seen[value] = struct{}{}
	}
}

func TestDNSRecordAndRCodeValues(t *testing.T) {
	if DNSRecordTypeTXT != 16 {
		t.Fatalf("unexpected TXT qtype: got=%d want=%d", DNSRecordTypeTXT, 16)
	}
	if DNSRecordTypeOPT != 41 {
		t.Fatalf("unexpected OPT qtype: got=%d want=%d", DNSRecordTypeOPT, 41)
	}
	if DNSRCodeNoError != 0 || DNSRCodeRefused != 5 {
		t.Fatalf("unexpected rcode values: noerror=%d refused=%d", DNSRCodeNoError, DNSRCodeRefused)
	}
	if DNSQClassIN != 1 {
		t.Fatalf("unexpected IN qclass: got=%d want=%d", DNSQClassIN, 1)
	}
}
