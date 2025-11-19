import socket
import threading
import sys
import os

#sunucudan gelecek verileri bekleme
def receive(socket, signal):
    while signal:
        try:
            data = socket.recv(32)
            print(str(data.decode("utf-8")))
        except:
            print("Sunucudan ayrıldınız")
            signal = False
            break

#sunucu için host ve port değerlerinin alınması
host = input("Host: ")
port = int(input("Port: "))
existsignal = True

#sunucuya erişme denemesi
try:
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.connect((host, port))
except:
    print("Sunucu ile iletişim yok")
    input("Enter'a basınız")
    sys.exit(0)

#veri gelmesi için thread
receiveThread = threading.Thread(target = receive, args = (sock, True))
receiveThread.start()

print("elma")

#sunucuya veri gönderme
while existsignal:
    message = input()
    sock.sendall(str.encode(message))
    if str(message)=="ÇIK":
        os.kill(os.getpid(), 9)
