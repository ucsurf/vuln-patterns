package p
func v(ts *timestamp.Timestamp, sig Sig) {
	// ruleid: x509-verify-time-from-untrusted-timestamp
	o := x509.VerifyOptions{Roots: roots, CurrentTime: ts.Time}
	// ruleid: x509-verify-time-from-untrusted-timestamp
	o2 := x509.VerifyOptions{CurrentTime: sig.SigningTime}
	// ok: x509-verify-time-from-untrusted-timestamp
	o3 := x509.VerifyOptions{Roots: roots, CurrentTime: time.Now()}
	_, _, _ = o, o2, o3
}
