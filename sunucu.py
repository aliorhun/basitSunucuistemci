import socket
import threading

#başlangıç tanımlamaları
connections = []
total_connections = 0

#İstemci ile ilgili işlemleri yapan sınıf
class Client(threading.Thread):
    def __init__(self, socket, address, id, name, signal):
        threading.Thread.__init__(self)
        self.socket = socket
        self.address = address
        self.id = id
        self.name = name
        self.signal = signal
    
    def __str__(self):
        return str(self.id) + " " + str(self.address)
    
    #İstemciden gelen veriye göre işlemler
    def run(self):
        while self.signal:
            try:
                data = self.socket.recv(32)
            except:
                print("İstemci " + str(self.address) + " sunucudan ayrıldı")
                self.signal = False
                connections.remove(self)
                break
            if str(data.decode("utf-8"))=="ÇIK":
                print("İstemci " + str(self.address) + " sunucudan ayrılmak istedi")
                self.signal = False
                connections.remove(self)
                break
            if data != "":
                print("ID " + str(self.id) + ": " + str(data.decode("utf-8")))
                for client in connections:
                    if client.id != self.id:
                        client.socket.sendall(data)

# Yeni bağlantıları bekleyen döngü
def newConnections(socket):
    while True:
        sock, address = socket.accept()
        global total_connections
        connections.append(Client(sock, address, total_connections, "Name", True))
        connections[len(connections) - 1].start()
        print("Yeni Bağlantı: ID " + str(connections[len(connections) - 1]))
        total_connections += 1

def main():
    #Sunucu için host ve port bilgileri
    host = input("Host: ")
    port = int(input("Port: "))
    print("armut")

    #Yeni sunucu soketi oluşturmak
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.bind((host, port))
    sock.listen(5)

    #Yeni bağlantıları bekleyen thread
    newConnectionsThread = threading.Thread(target = newConnections, args = (sock,))
    newConnectionsThread.start()
    
main()
