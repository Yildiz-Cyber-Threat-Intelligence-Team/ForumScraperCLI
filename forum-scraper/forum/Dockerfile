# Temel Go imajını kullanıyoruz
FROM golang:1.23 AS builder

# Çalışma dizinimizi ayarlıyoruz
WORKDIR /app

# Uygulama dosyalarını kopyalıyoruz
COPY . .

# Go modül dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod go.sum ./
RUN go mod download

# Uygulamayı derliyoruz
RUN go build -o forum_monitoring_tools .

# Tor'u ve Google Chrome'u kurmak için ayrı bir imaj
FROM debian:bookworm-slim

#Konteynerin içindeki belirli bir dizini ana makinenize bağlamaya yardımcı olur.
VOLUME /app/screenshots

# Gerekli bağımlılıkları yükle
RUN apt-get update && \
    apt-get install -y tor wget gnupg2 && \
    wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

# Tor konfigürasyon dosyasını ayarlıyoruz
RUN echo "SOCKSPort 0.0.0.0:9150" >> /etc/tor/torrc

# Go uygulamamızı kopyalıyoruz
COPY --from=builder /app/forum_monitoring_tools /usr/local/bin/forum_monitoring_tools

# Uygulamayı başlatan giriş noktası komutu
CMD ["sh", "-c", "service tor start && exec /usr/local/bin/forum_monitoring_tools"]
