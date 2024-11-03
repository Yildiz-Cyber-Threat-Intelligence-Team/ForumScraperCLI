# Untitled

<h1>ForumScraperCLI</h1>

**ForumScraperCLI**, otomatik forum izlemeye yönelik bir komut satırı aracıdır. En az 10 forumdaki gönderileri sıyırarak tartışmaları, konuları ve anahtar kelimeleri gerçek zamanlı olarak izler. Bu araç, forum gönderilerini toplama sürecini basitleştirir; istihbarat toplamak ve forum etkinlikleri hakkında güncel bilgi sahibi olmak için idealdir.

<h2>Özellikler</h2>

- **Otomatik Kazıma**: Çeşitli forumlardan veri toplar.
- **Gerçek Zamanlı İzleme**: Tartışmaları ve konuları anlık olarak izler.
- **Esnek Konfigürasyon**: Forum URL'lerini ve CSS öğesi seçicilerini kolayca ekleyebilir veya değiştirebilirsiniz.

<h2>Kurulum ve Docker Kullanımı</h2>

Bu proje Docker ile çalıştırılmak üzere yapılandırılmıştır. Docker, bağımlılıkların yönetimini basitleştirir ve aracı izole bir ortamda çalıştırmanıza olanak tanır.

### Docker ile Çalıştırma

1. **Docker İmajını Oluşturun:**
    
    Proje dizininde aşağıdaki komutu çalıştırarak Docker imajını oluşturun:
    
    ```bash
   

    docker-compose build
    
    ```
    
2. **Docker Konteynerlerini Başlatın:**
    
    Docker konteynerlerini arka planda çalıştırmak için:
    
    ```bash
   
   
    docker-compose up -d
    
    ```
    
    Bu işlem, iki konteyneri başlatır:
    
    - **tor_service**: Tor ağına erişim için bir proxy servisi sağlar, böylece araç .onion sitelerindeki forumlara erişebilir.
    - **forum_monitoring_app**: Forum verilerini sıyırmak ve ekran görüntüleri almak için ana uygulamayı çalıştırır.
3. **Konteynerlerin Çalıştığını Doğrulama:**
    
    Konteynerlerin başarıyla çalıştığını doğrulamak için aşağıdaki komutla durumu kontrol edebilirsiniz:
    
    ```bash
   
   
    docker ps
    
    ```
    
4. **Kayıtları İnceleme:**
    
    Ekran görüntüsü ve verilerin kaydedildiğinden emin olmak için `forum_monitoring_app` konteynerinin loglarını kontrol edin:
    
    ```bash
   
   
    docker logs forum_monitoring_app
    
    ```
    

<h2>Kullanım</h2>

Docker üzerinde yukarıdaki adımları tamamladıktan sonra, proje dizininde aşağıdaki komutu kullanarak uygulamayı başlatabilirsiniz:

```bash


go run main.go

```

<h3>Arayüze Erişim</h3>

Sunucu, http://localhost:8081 adresinde başlatılacaktır. Bu adrese tarayıcınızdan giderek ForumScraperCLI arayüzüne erişebilir, forum gönderilerini izlemeye ve ekran görüntüsü almaya başlayabilirsiniz.

<h2>Gönderileri Kazıma</h2>
<img src="menu.png" />

**1.** Açılır menüden bir forum seçin.<br>
**2.** Seçilen forumdan ekran görüntüsü almak ve bağlantıları çekmek için **Ekran Görüntüsü Al** butonuna tıklayın.<br>

Uygulama, ilgili foruma giderek son gönderiyi alacak ve bağlantısını **links.txt** dosyasına ekleyecektir.
