## 说明
由于go标准库不提供私钥加密、公钥解密的实现，需要自己对签名函数做改造

## Example
```
    privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAp1tsB3SL/8kEPAy7+RYQvP4GS2DBJAcSpuo2tg1XoeTr+j54
k/+canfmjO4R5qwRB3OfFHHGZhw1DO2zDRF81ixlg7wpcFTxdeMyQvm8ae1nvLpl
jYPNvNHV96MJTtKOJ/+hhMDwkjKV5J79oL2lAvawQzrHC6UfS88SIRJ8Q87RPDDS
eO7UbFvadEPv9IJGrgH/70/1tQMK9X6uNW9PIsgYAfY/k2+vhO2gN5eXdpZOyQgq
xFOoaV3Rt6uwTlyzpDma/ECMWmR4ypo4o7M4RBFgOMuU73yu7arTG1/y494zWtMj
hux/v1fAkssUmBC5OqBMd5jBWXfgOB6wnmVpowIDAQABAoIBAGdEyoxvJlRXvf0v
/m3dZ8a3GTs6IPQSSTgmHkyIPyLvF437LnxkpigshF9CKFmmo3ttpSQAcORhU3r+
0hWmTNn720hGu2tNfxG3eDRL1i+nBvQ2MKP6SOHSua1tNTkmRAH6g7hfNhOAk23X
jbm6TDBf+LQDeP9+jKlCxXUS6wsLfC5Qep2nRxlTjws1NfTytVsONhdgLzO0rtR9
lQz7RVmfF+ECdwav75zAGajrU5BCpC0xi+nLMcRelB7BX8iCiXV17KwDYgDPOHoA
DCCm+hgpHcALQtzpYCHBpc1B876VlVhxlAS9F3TriCWEilgmdq4uF3vYyjjUht6E
LQ3mxkECgYEAzxHqbY7Xg3DvREY/iHngVuZgjYaOdEXcxHzSWPopiCAmeJ2OEYKI
eoJBcZccVqfr9LpmYudG9Dp4xBdlSclN2mv3s4YDTmT2+SrrFd85CarcpX91KTEx
AcpLSKVDZW8nOJh7Jt5/4I32pd1m2M/b3MxczI7U/yfS0cFXcPY9RxcCgYEAzucx
udq7+iAskla1ChRuUSEyBnQpVSXLQP6YhZRNASsEKg1JSZvFyyMV/DsQEHiibGQU
u7Aijx2Fo4vHzWoM9lbmeuv70ux0yHDJE6vlUzgAeqdAbGaOI3iLGLlq5A1LIseR
FXvNBPjxtV2MCixFlKPqOx+15QaoJcqFvqf7CVUCgYEAkINmYnOwQNGOis2kBXFZ
egxqEht4S/l967tZajOczJ3ze3Yp8lpxOV6yob3eTzP+XtvAHQJf6I7uPZw+WlKd
fPSg+K9sm0enfZnQW4FgYzjVqMLdV5Q3KZKvG9IKyT3b14nDWQsrn+Iz4uozntji
DqSplzGPwUtfvHf8uJ+BgJ0CgYEAnlDzkLXZblveHx+vqYXWn8eZwqxx0BB6RUJK
rLO8dG9Y4WCHOoayQ2yWbwk2kpPsHKo73x24Y/AACRN1EC3cu0XlWBtn9o4uEg2T
HrVCZXoNi8yclBWpH2UmcgW2z8eIPZWc42Z2ix59vPxWopYbq8/8876g7R053qjn
y+o3QX0CgYAPCnBpYdxcnHSC3/88tIHsubSQzgpi118NbcnxqwJgAiI2nibyYF+M
MNW2wveMz1Wiv5XRUBoG2Way9pij5wZRQfgtbenhmwESgeY4gtWDuBLk9a1OTwAm
OCG0I0udRtbo4LW1PL6LH7p0QEhiCNz3lxaXsrpJg4WJ9iwubtib9g==
-----END RSA PRIVATE KEY-----
`

    pri, pub, err := NewKeys([]byte(privateKey))
    cipher, err := pri.Encrypt("hello world")
    txt, err := pub.Decrypt(cipher)
    if !bytes.Equal(jsonRaw, txt) {
		log.Fatalln("warn: unexpect result")
	}
```