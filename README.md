# basitSunucuistemci

Aşağıdaki adresteki uygulamadaki process kapanmama sorununun çözülmesi ile oluşturulmuş yeni depo
https://github.com/pricheal/python-client-server/

## Kullanımı:

Ağ üzerinde veya aynı makine üzerinde 1 sunucu ve birden fazla istemci uygulamalarını ayrı ayrı çalıştırarak kullanılabilecek yapıdır.

### Aynı makine üzerinde kullanımı 

Öncelikle aşağıdaki komut ile sunucu ayağa kaldırılır. Host ve port isteği yerine aşağıdaki gibi 127.0.0.1 ve örnek bir port olarak 8989 yazılır:

> python3 sunucu.py
> 
> Host: 127.0.0.1
> 
> Port: 8989
 
Bu şekilde sistem sunucu görevini yapmaya başlamıştır. Başka bir terminalde aşağıdaki komut ve değerler yazılır:

> python3 istemci.py
> 
> Host: 127.0.0.1
> 
> Port: 8989

Bu şekilde sisteme bağlantı kurulmuş olur ve sunucuda aşağıdaki gibi bir log oluşur:

> Yeni Bağlantı: ID 0 ('127.0.0.1', 32780)

İstemci üzerinde bir şeyler yazarsanız bu loglar sunucuya gelir. 

> ID 0: Örnek içerik

Başka bir terminal açıp yeni bir istemci ayağa kaldırabilirsiniz:

> python3 istemci.py
> 
> Host: 127.0.0.1
> 
> Port: 8989

Bu işlemden sonra sunucu'daki log'da aşağıdaki gibi çıktı oluşur.

> Yeni Bağlantı: ID 1 ('127.0.0.1', 32798)

Şimdi bu iki istemci birbiriyle konuşabilir durumdadır.
