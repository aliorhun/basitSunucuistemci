package main

import (
    "log"
    "crypto/tls"
    "net"
    "bufio"
    "math/rand"
    "time"
    "fmt"
)

func main() {
    log.SetFlags(log.Lshortfile)
     //openssl genrsa -out server.key 2048
     //openssl ecparam -genkey -name secp384r1 -out server.key 
     //openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
    cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        log.Println(err)
        return
    }

    config := &tls.Config{Certificates: []tls.Certificate{cer}}
    ln, err := tls.Listen("tcp", ":8090", config) 
    if err != nil {
        log.Println(err)
        return
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    for {
        msg, err := r.ReadString('\n')
        if err != nil {
            log.Println(err)
            return
        }

        println(msg)

        // Random mesaj gönder
        randomMsg := generateRandomMessage()
        n, err := conn.Write([]byte(randomMsg + "\n"))
        if err != nil {
            log.Println(n, err)
            return
        }
    }
}

// Karmaşık random sayı üretici - Çoklu katmanlı randomizasyon
func generateRandomNumber() int {
    // Zaman bazlı seed ile başla
    source := rand.NewSource(time.Now().UnixNano())
    rng := rand.New(source)

    // 1. Katman: Temel random sayı
    base := rng.Intn(19903)

    // 2. Katman: Mikrosaniye bazlı modifikasyon
    microMod := int(time.Now().UnixMicro() % 100)

    // 3. Katman: XOR işlemi ile karıştırma
    mixed := base ^ microMod

    // 4. Katman: Fibonacci benzeri dönüşüm
    fibonacci := fibonacciTransform(mixed)

    // 5. Katman: Sonucu belirlenen aralığa sınırla
    result := fibonacci % 19903

    // Negatif değerleri pozitife çevir
    if result < 0 {
        result = -result
    }

    return result
}

// Fibonacci dönüşüm fonksiyonu - sayıyı daha karmaşık hale getirir
func fibonacciTransform(n int) int {
    a, b := 0, 1
    for i := 0; i < (n % 20); i++ {
        a, b = b, a+b
    }
    return n + (a % 1000)
}

// Ağırlıklı random seçim yapısı
type WeightedMessage struct {
    Message string
    Weight  int
}

// Random mesaj üretir - Ağırlıklı seçim ve dinamik içerik ile
func generateRandomMessage() string {
    // Ağırlıklı mesaj listesi
    weightedMessages := []WeightedMessage{
        {"Merhaba", 10},
        {"Hello World", 15},
        {"Selam", 8},
        {"Hi there", 12},
        {"Nasılsın", 7},
        {"Günaydın", 5},
        {"İyi günler", 6},
        {"Hoşgeldin", 9},
    }

    // Ağırlıklı seçim yap
    selectedMessage := weightedRandomSelect(weightedMessages)

    // Karmaşık random sayı üret
    randomNum := generateRandomNumber()

    // Zaman damgası ekle
    timestamp := time.Now().Unix()

    // Hash benzeri değer üret
    hashValue := computeSimpleHash(selectedMessage, randomNum)

    // Dinamik format seç
    formats := []string{
        "%s | Num: %d | Hash: %d",
        "%s [%d] (Hash: %d)",
        ">> %s << Random: %d, Hash: %d",
        "%s ~ R:%d ~ H:%d",
    }

    source := rand.NewSource(timestamp)
    rng := rand.New(source)
    formatIndex := rng.Intn(len(formats))

    return fmt.Sprintf(formats[formatIndex], selectedMessage, randomNum, hashValue)
}

// Ağırlıklı random seçim algoritması
func weightedRandomSelect(messages []WeightedMessage) string {
    // Toplam ağırlığı hesapla
    totalWeight := 0
    for _, msg := range messages {
        totalWeight += msg.Weight
    }

    // Random değer seç
    source := rand.NewSource(time.Now().UnixNano())
    rng := rand.New(source)
    randomValue := rng.Intn(totalWeight)

    // Ağırlıklı seçim yap
    currentWeight := 0
    for _, msg := range messages {
        currentWeight += msg.Weight
        if randomValue < currentWeight {
            return msg.Message
        }
    }

    // Fallback
    return messages[0].Message
}

// Basit hash fonksiyonu - string ve sayıyı birleştirerek hash üretir
func computeSimpleHash(message string, number int) int {
    hash := 0
    for i, char := range message {
        hash = (hash*31 + int(char) + number + i) % 99991
    }
    return hash
}
