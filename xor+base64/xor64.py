import base64

originalShellcode = b"\xfc\x48\x83\"
encryptedShellcode = bytes([byte ^ 0x77 for byte in originalShellcode])
encodedShellcode = base64.b64encode(encryptedShellcode).decode('utf-8')

print(encodedShellcode)
