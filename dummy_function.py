#!/usr/bin/env python3
"""
Dummy function file - Test dosyası
"""

def dummy_function():
    """Basit bir test fonksiyonu"""
    print("Merhaba! Bu bir dummy fonksiyondur.")
    result = 42 * 2
    print(f"Hesaplama sonucu: {result}")
    return result

def another_dummy_function(name="Kullanıcı"):
    """Parametreli bir test fonksiyonu"""
    message = f"Selam {name}, bu fonksiyon çalışıyor!"
    print(message)
    return message

if __name__ == "__main__":
    print("=== Dummy fonksiyonlar çalıştırılıyor ===")
    print()

    # İlk fonksiyonu çağır
    result1 = dummy_function()
    print()

    # İkinci fonksiyonu çağır
    result2 = another_dummy_function("Test Kullanıcısı")
    print()

    print(f"Tüm fonksiyonlar başarıyla çalıştırıldı!")
    print(f"Sonuç 1: {result1}")
    print(f"Sonuç 2: {result2}")
