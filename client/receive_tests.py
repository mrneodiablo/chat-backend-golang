from websocket import create_connection
import time, base64, json, datetime, hashlib
from Crypto import Random
from Crypto.Cipher import AES

key = "9w8]N7Uy;HaZFcmL"

def generatesession(id):
    # md5sum(secretkey + userId + yymmddhhmm)
    time_data = datetime.datetime.utcnow().strftime("%Y%m%d%H%M")
    data = key + str(id) + time_data
    m = hashlib.md5()
    m.update(data)
    return m.hexdigest()

def decrypt(key, value, block_segments=False):
    value = str(value)
    value = base64.b64decode(value + '=' * (4 - len(value) % 4), '-_')
    iv, value = value[:AES.block_size], value[AES.block_size:]
    if block_segments:
        remainder = len(value) % 16
        padded_value = value + '\0' * (16 - remainder) if remainder else value
        cipher = AES.new(key, AES.MODE_CFB, iv, segment_size=128)
        return cipher.decrypt(padded_value)[:len(value)]
    return AES.new(key, AES.MODE_CFB, iv).decrypt(value)


def encrypt(key, value, block_segments=False):
    iv = Random.new().read(AES.block_size)
    if block_segments:
        remainder = len(value) % 16
        padded_value = value + '\0' * (16 - remainder) if remainder else value
        cipher = AES.new(key, AES.MODE_CFB, iv, segment_size=128)
        value = cipher.encrypt(padded_value)[:len(value)]
    else:
        value = AES.new(key, AES.MODE_CFB, iv).encrypt(value)
    return base64.b64encode(iv + value, '-_').rstrip('=')

url = "ws://localhost:2345/chat/private?userId=1&session=" + generatesession(1)
ws = create_connection(url)

while 1:
	result =  ws.recv()
	print "----------------------------------"
	print(decrypt(key, result, block_segments=True))
