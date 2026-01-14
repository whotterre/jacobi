from hashlib import sha256
def genCookieToken(pkh: str) -> str:
    res = ""
    payload = pkh
    
    '''
    What does it mean to prepend 00 bits to the beginning of the hex number?
    '''
    # Append version 00 first
    payload = "00" + payload
    payload_bytes = bytes.fromhex(payload)
    # Double hash payload with SHA256
    hash_1 = sha256(payload_bytes).digest()
    double_hash = sha256(hash_1).hexdigest()
    
    # How do i extract the last 4 bytes which form the checksum?
    n = len(double_hash)
    checksum = double_hash[:8]
    
    cookie_token = payload + checksum
    print("Cookie token", cookie_token)
    res = base58encode(cookie_token)
    
    return res




def base58encode(payload: str) -> str:
    # count number of leading 00s
    ct = 0
    for i in range(0, len(payload), 2):
        if payload[i:i+2] == "00":
            ct += 1
        else:
            break
    
    # Convert to decimal number
    dec_value = int(payload, 16)
    
    print(dec_value)
    res = ""
    # Lookup table for base58
    # Base58 alphabet: 123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
    lookup_table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
     
    # Repeatedly divide by 58 till it's 0
    while dec_value != 0:
        rem = dec_value % 58
        res += lookup_table[rem]
        dec_value //= 58
        
    # Add leading 1s for each leading 00 byte
    res += "1" * ct
    
    return res[::-1]
        
        
        
print(genCookieToken("5f2613791b36f667fdb8e95608b55e3df4c5f9eb"))