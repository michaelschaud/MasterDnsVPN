// ==============================================================================
// MasterDnsVPN
// Author: MasterkinG32
// Github: https://github.com/masterking32
// Year: 2026
// ==============================================================================

package udpserver

import "time"

type socks5FragmentKey struct {
	sessionID   uint8
	streamID    uint16
	sequenceNum uint16
}

func (s *Server) collectSOCKS5SynFragments(sessionID uint8, streamID uint16, sequenceNum uint16, payload []byte, fragmentID uint8, totalFragments uint8, now time.Time) ([]byte, bool, bool) {
	if totalFragments == 0 {
		totalFragments = 1
	}
	assembled, ready, completed := s.socks5Fragments.Collect(
		socks5FragmentKey{
			sessionID:   sessionID,
			streamID:    streamID,
			sequenceNum: sequenceNum,
		},
		payload,
		fragmentID,
		totalFragments,
		now,
		s.dnsFragmentTimeout,
	)
	return assembled, ready, completed
}

func (s *Server) purgeSOCKS5SynFragments(now time.Time) {
	if s == nil || s.socks5Fragments == nil {
		return
	}
	s.socks5Fragments.Purge(now, s.dnsFragmentTimeout)
}

func (s *Server) removeSOCKS5SynFragmentsForSession(sessionID uint8) {
	if s == nil || s.socks5Fragments == nil || sessionID == 0 {
		return
	}
	s.socks5Fragments.RemoveIf(func(key socks5FragmentKey) bool {
		return key.sessionID == sessionID
	})
}
