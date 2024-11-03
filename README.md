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

## EN

**ForumScraperCLI** is a command-line tool designed for automatic forum monitoring. It scrapes posts from at least 10 forums to track discussions, topics, and keywords in real-time. This tool simplifies the process of collecting forum posts, making it ideal for gathering intelligence and staying updated on forum activities.

<h2>Features</h2>

- **Automatic Scraping**: Collects data from various forums.
- **Real-Time Monitoring**: Tracks discussions and topics instantly.
- **Flexible Configuration**: Easily add or change forum URLs and CSS element selectors.

<h2>Setup and Docker Usage</h2>

This project is configured to run with Docker. Docker simplifies dependency management and allows you to run the tool in an isolated environment.

### Running with Docker

1. **Build the Docker Image:**
    
    Run the following command in the project directory to build the Docker image:
    
    ```bash
    
    docker-compose build
    
    ```
    
2. **Start Docker Containers:**
    
    To run the Docker containers in the background:
    
    ```bash
    
    docker-compose up -d
    
    ```
    
    This will start two containers:
    
    - **tor_service**: Provides a proxy service to access the Tor network, enabling the tool to reach forums on .onion sites.
    - **forum_monitoring_app**: Runs the main application for scraping forum data and taking screenshots.
3. **Verify Containers are Running:**
    
    To confirm the containers are running successfully, check their status with:
    
    ```bash
    
    docker ps
    
    ```
    
4. **Inspect Logs:**
    
    To ensure screenshots and data are being saved, check the logs of the `forum_monitoring_app` container:
    
    ```bash
    
    docker logs forum_monitoring_app
    
    ```
    

<h2>Usage</h2>

After completing the steps above with Docker, you can start the application by running the following command in the project directory:

```bash

go run main.go

```

<h3>Accessing the Interface</h3>

The server will start at [http://localhost:8081](http://localhost:8081/). You can access the ForumScraperCLI interface by navigating to this address in your browser, where you can begin monitoring forum posts and taking screenshots.

<h2>Scraping Posts</h2>
<img src="menu.png" />

**1.** Select a forum from the dropdown menu.<br>
**2.** Click the **Take Screenshot** button to capture a screenshot and fetch links from the selected forum.<br>

The tool will navigate to the selected forum, retrieve the latest post, and save its link in the **links.txt** file.
