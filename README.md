# Kripto Fiyat Bildirimi Uygulaması

Bu uygulama, belirli bir kripto para biriminin en yüksek ve en düşük fiyatlarını izleyerek bu değerleri günceller ve bu değerleri bir API aracılığıyla sunar. Ayrıca, Telegram botu aracılığıyla bu fiyatları belirli bir kullanıcıya bildirim olarak gönderir.

## Kurulum

1. Bu kodu çalıştırmak için öncelikle Go programlama dilinin yüklü olması gerekmektedir.
2. Telegram botu oluşturmak için Telegram BotFather'dan bir bot tokenı alın.
3. Gerekli kütüphaneleri yüklemek için terminalde `go get gopkg.in/tucnak/telebot.v2` komutunu çalıştırın.
4. Kodu çalıştırarak uygulamayı başlatın.

## Kullanım

- Uygulama, `/price` endpoint'i üzerinden en yüksek ve en düşük fiyatları sunar.
- Telegram botunu başlatmak için `/botToken` endpoint'ına bot tokenınızı POST isteğiyle gönderin.
- Telegram'da botunuzla etkileşim kurarak kripto fiyatlarını görüntüleyebilirsiniz.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
