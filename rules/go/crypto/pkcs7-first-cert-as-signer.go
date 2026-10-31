package p
func v(p7 *pkcs7.PKCS7) {
	// ruleid: pkcs7-first-cert-as-signer
	signer := p7.Certificates[0]
	// ruleid: pkcs7-first-cert-as-signer
	name := sd.Certificates[0].Subject
	// ok: pkcs7-first-cert-as-signer
	s2 := p7.GetOnlySigner()
	_, _, _ = signer, name, s2
}
