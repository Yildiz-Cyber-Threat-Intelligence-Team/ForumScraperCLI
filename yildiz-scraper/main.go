package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

// Web sitelerini tanımla
var forums = []struct {
	name        string
	url         string
	elementPath string // HTML seçicileri
}{
	{"Breaking Bad", "http://bbzzzsvqcrqtki6umym6itiixfhni37ybtt7mkbjyxn2pgllzxf2qgyd.onion/th-all-threads/newest", "div.structItem:nth-child(1)"},
	{"Darkzone", "http://darkobds5j7xpsncsexzwhzaotyc4sshuiby3wtxslq5jy2mhrulnzad.onion/darkzone-forum/forum/darkzone-forum-community/", "div.content-element:nth-child(1) > div:nth-child(2)"},
	{"DeepWeb Question and Answers", "http://b7ehf7dabxevdsm5szkn2jecnliwzoxlsn4lijxqxikrlykbbsfrqfad.onion/", "div#q606.qa-q-list-item.qa-q-list-item-featured"},
	{"Respostas Ocultas", "http://xh6liiypqffzwnu5734ucwps37tn2g6npthvugz3gdoqpikujju525yd.onion/", "#q297120"},
	{"Out3r Space", "https://reycdxyc24gf7jrnwutzdn3smmweizedy7uojsa7ols6sflwu25ijoyd.onion/archives/", "li.post-item:nth-child(2)"},
	{"The Tor Forum", "http://thetorfkljnxqhu4u3fpdbda4g3vv4njslayyxxhfbk32nffh5sn3wad.onion/", "div.qa-q-list-item:first-child"},
	{"DarkWeb Forums", "http://forums56xf3ix34sooaio4x5n275h4i7ktliy4yphhxohuemjpqovrad.onion/forums/general-discussion.9/", "div.structItem:nth-child(1) > div:nth-child(2)"},
	{"Suprbay", "http://suprbaydvdcaynfo4dgdzgxb4zuso7rftlil5yg5kqjefnw4wq4ulcad.onion/", "table.tborder:nth-child(1)"},
	{"Hidden Answers", "http://7eoz4h2nvw4zlr7gvlbutinqqpm546f5egswax54az6lt2u7e3t6d7yd.onion/", "#q13246"},
	{"FrenchPool", "http://jkie5viyrmymttownlksylz5vipyxxvs6qgy2yybgbssoiuf7a7klpqd.onion/viewforum.php?id=1", "tr.rowodd:nth-of-type(3)"},
	{"Wall of Shame", "http://mblogci3rudehaagbryjznltdp33ojwzkq6hn2pckvjq33rycmzczpid.onion/", "div.col-lg-4:nth-child(1)"},
}

var currentLanguage = "tr" // Varsayılan dil Türkçe

// Element ekran görüntüsü alma ve link çekme işlevi
func takeElementScreenshotAndLink(url, elementPath string, screenshotIndex int) error {
	allocCtx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.ProxyServer("socks5://127.0.0.1:9150"),
		)...,
	)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var buf []byte
	var postLink string

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(elementPath, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			time.Sleep(30 * time.Second)
			return nil
		}),
		chromedp.Screenshot(elementPath, &buf, chromedp.ByQuery),
		chromedp.AttributeValue(elementPath+" a", "href", &postLink, nil),
	})
	if err != nil {
		return fmt.Errorf("element ekran görüntüsü veya link alma hatası: %w", err)
	}

	fileName := fmt.Sprintf("screenshot_%d_%d.png", time.Now().Unix(), screenshotIndex)
	filePath := filepath.Join("screenshots", fileName)
	if err := os.MkdirAll("screenshots", 0755); err != nil {
		return fmt.Errorf("klasör oluşturulamadı: %w", err)
	}
	if err := os.WriteFile(filePath, buf, 0644); err != nil {
		return fmt.Errorf("dosya kaydedilemedi: %w", err)
	}

	fullLink := url
	if postLink != "" {
		fullLink = url + "\n" + postLink
	}

	linkFileName := "links.txt"
	f, err := os.OpenFile(linkFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("link dosyası açma hatası: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(fullLink + "\n\n"); err != nil {
		return fmt.Errorf("link yazma hatası: %w", err)
	}

	fmt.Printf("Element ekran görüntüsü ve link kaydedildi: %s\n", filePath)
	return nil
}
func renderIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    // Dil seçimine göre başlık ve buton metinleri
    title := "Forum Görüntüleme ve Ekran Görüntüsü Alma"
    buttonText := "Ekran Görüntüsü Al"
    footerText := "Gizlilik için ekran görüntüsü alma uygulaması"
    selectPrompt := "Bir Forum Seçin:"

    if currentLanguage == "en" {
        title = "Forum Viewer and Screenshot Taker"
        buttonText = "Take Screenshot"
        footerText = "Screenshot application for privacy"
        selectPrompt = "Select a Forum:"
    }

    fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="%s">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <div class="loader" id="loader"></div>
    <div class="container">
        <h1>%s</h1>

        <div class="language-selector">
            <button onclick="switchLanguage('tr')">Türkçe</button>
            <button onclick="switchLanguage('en')">English</button>
        </div>

        <form action="/screenshot" method="POST" onsubmit="showLoader();">
            <div style="text-align: center;">
                <label for="forum-select" style="display: block; font-weight: bold; margin-bottom: 5px;">%s</label>
                <select id="forum-select" name="forum" style="margin-bottom: 15px;">
                    `, currentLanguage, title, title, selectPrompt)

    for i, forum := range forums {
        fmt.Fprintf(w, `<option value="%d">%s</option>`, i, forum.name)
    }
    fmt.Fprintf(w, `
                </select>
                <br>
                <button type="submit" style="padding: 10px 20px; font-size: 1em; background-color: #9c27b0; color: #fff; border: none; border-radius: 5px; cursor: pointer;">%s</button>
            </div>
        </form>

        <footer>
            <img src="/static/cti.jpeg" alt="Logo" class="logo">
            <span id="footer-text">%s</span>
        </footer>
    </div>

    <div id="modal" class="modal">
        <div class="modal-content">
            <span id="closeBtn" class="close-btn">&times;</span>
            <p id="modal-message"></p>
        </div>
    </div>

    <script>
        document.addEventListener("DOMContentLoaded", function() {
            document.getElementById("modal").style.display = "none";
        });

        let isLanguageSwitching = false; // Dil değişikliği sırasında modal'ı engellemek için bayrak

        // Yükleniyor animasyonunu başlat
        function showLoader() {
            document.getElementById("loader").style.display = "block";
        }

        // Yükleniyor animasyonunu durdur
        function hideLoader() {
            document.getElementById("loader").style.display = "none";
        }

        // Modal'ı göster, yalnızca isLanguageSwitching olmadığında çalışır
        function showModal(message) {
            if (isLanguageSwitching) {
                isLanguageSwitching = false; // Dil değişikliği sırasında modal'ı göstermeyi engelle
                return;
            }
            document.getElementById("modal").style.display = "flex";
            document.getElementById("modal-message").textContent = message;
        }

        // Modal'ı kapat
        document.getElementById("closeBtn").onclick = function() {
            document.getElementById("modal").style.display = "none";
        }

        // Dil değiştirme işlevi
        function switchLanguage(lang) {
            isLanguageSwitching = true; // Dil değişikliği sırasında modal'ın tetiklenmesini engelle
            fetch("/switch-language?lang=" + lang)
                .then(() => location.reload()); // Sayfayı yeniden yükle
        }

        // Ekran görüntüsü alma formunun gönderimi
        document.querySelector("form").onsubmit = function(event) {
            event.preventDefault(); // Formun normal gönderimini engelle

            showLoader(); // Yükleniyor simgesi gösteriliyor
            const formData = new FormData(this);

            fetch("/screenshot", {
                method: "POST",
                body: formData
            })
            .then(response => response.json())
            .then(data => {
                // Başarı veya hata mesajını modal'da göster
                showModal(data.message); // Sadece ekran görüntüsü başarılı veya hata durumunda göster
            })
            .catch(error => {
                showModal("Bir hata oluştu: " + error.message);
            })
            .finally(() => {
                hideLoader(); // İşlem tamamlandığında yükleniyor simgesini gizle
            });
        };
    </script>
</body>
</html>`, buttonText, footerText)
}

func screenshotHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Yalnızca POST desteklenir", http.StatusMethodNotAllowed)
		return
	}

	forumIndexStr := r.FormValue("forum")
	forumIndex, err := strconv.Atoi(forumIndexStr)
	if err != nil || forumIndex < 0 || forumIndex >= len(forums) {
		http.Error(w, "Geçersiz forum seçimi", http.StatusBadRequest)
		return
	}

	selectedForum := forums[forumIndex]
	err = takeElementScreenshotAndLink(selectedForum.url, selectedForum.elementPath, 1)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		message := "Ekran görüntüsü alınamadı: " + err.Error()
		if currentLanguage == "en" {
			message = "Screenshot could not be taken: " + err.Error()
		}
		fmt.Fprintf(w, `{"status": "error", "message": "%s"}`, message)
		return
	}

	message := "Ekran görüntüsü başarıyla alındı ve kaydedildi."
	if currentLanguage == "en" {
		message = "Screenshot successfully taken and saved."
	}
	fmt.Fprintf(w, `{"status": "success", "message": "%s"}`, message)
}

func switchLanguageHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "en" || lang == "tr" {
		currentLanguage = lang
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", renderIndex)
	http.HandleFunc("/screenshot", screenshotHandler)
	http.HandleFunc("/switch-language", switchLanguageHandler)

	fmt.Println("Sunucu http://localhost:8080 adresinde başlatıldı.")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
