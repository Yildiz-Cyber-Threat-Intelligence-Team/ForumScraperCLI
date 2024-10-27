<h1> ForumScraperCLI </h1>
 **ForumScraperCLI** , otomatik forum izlemeye yönelik bir komut satırı aracıdır. En az 10 forumdaki gönderileri sıyırıp tartışmaları, konuları ve anahtar kelimeleri gerçek zamanlı olarak izliyor. Bu araç, forum gönderilerini toplama sürecini basitleştirir; istihbarat toplamak ve forum etkinlikleriyle ilgili güncel bilgileri korumak için idealdir.

<h2>Özellikler</h2>

- **Otomatik Kazıma**: Çeşitli forumlardan veri toplar.
- **Gerçek Zamanlı İzleme**: Tartışmaları ve konuları anlık olarak izler.
- **Esnek Konfigürasyon**: Forum URL'lerini ve HTML öğesi seçicilerini kolayca ekleyebilir veya değiştirebilirsiniz.

<h2>Kullanım</h2>
Uygulamayı çalıştırmak için proje dizinine gidin ve şu komutu çalıştırın: <br>
```bash
 go run main.go 
Sunucu **'http://localhost:8080'** adresinde başlayacaktır. Tarayıcınızı açın ve bu adrese giderek forum izleme aracına erişin.
<img src="terminal.png" />


<h2>Gönderileri Kazıma</h2>
<img src="menu.png" />
1.Açılır menüden bir forum seçin. <br>
2.Seçilen forumdan ekran görüntüsü almak ve bağlantıları çekmek için "Ekran Görüntüsü Al" butonuna tıklayın. <br>
Uygulama, ilgili foruma giderek son gönderiyi alacak ve bağlantısını **'links.txt'** dosyasına ekleyecektir.
