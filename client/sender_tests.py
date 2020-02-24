from websocket import create_connection
import time, base64, json, datetime, hashlib
from Crypto import Random
from Crypto.Cipher import AES

key = "9w8]N7Uy;HaZFcmL"


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

def generatesession(id):
    # md5sum(secretkey + userId + yymmddhhmm)
    time_data = datetime.datetime.utcnow().strftime("%Y%m%d%H%M")
    data = key + str(id) + time_data
    m = hashlib.md5()
    m.update(data)
    return m.hexdigest()

################## HEART BEAT ##############
#data = {
#        "action": 10,
#        "userId": 13
#        }

################## SEND ##################
data = {
    "action": 11,
    "userId": 12,
    "SenderUserId": 12,
    "ReceiverUserId": 1,
    "data": "Anh Nho em Nhieu aed",
}

################## FETCH ##################
#data = {
#     "action": 12,
#     "userId": 1,
#}

################## CONFIRM ##################
# data = {
#      "action": 13,
#      "userId": 1,
#      "senderuserId": 12,
#      "receiverUserId": 1,
#      "messageId":1543342219075,
# }

url = "ws://localhost:2345/chat/private?userId=12&session=" + generatesession(12)
ws = create_connection(url)

while 1:
    print("Sending 'Hello, World'...")
    encrypted = encrypt(key,json.dumps(data), block_segments=True)
    print("RAW:       " + json.dumps(data))
    print "----------------------------------"
    print("ENCRYPT:   " + encrypted)
    print "----------------------------------"
    ws.send(encrypted)
    result =  ws.recv()
    print("RESPONE:    " + decrypt(key, result, block_segments=True))
    time.sleep(1)
    print "----------------------------------"
