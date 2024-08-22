# Uy vazifasi: GitHub Webhook for "repository" Events Only

## Maqsad
Foydalanuvchidan faqat `repository` voqealarini tinglaydigan va yuborilgan ma'lumotlarni to'g'ri qayta ishlaydigan `Go webhook` serverini yaratish talab qilinadi.

## Talablar
1. **GitHub Webhookni sozlash:**
    - Repozitoriyangiz uchun `GitHub` da webhookni ro'yxatdan o'tkazing
    - Faqat Repository voqealarini tanlang (`repository.created`, `repository.deleted` kabi).
    - Webhookni Go serveringizga yuborish uchun sozlang.

2. **Webhook Server:**
    - `Go` serverini qurib, `/webhook` endpointini e’lon qiling.
    - Faqat `repository` voqealarini tinglayotganingizga ishonch hosil qiling.
    - `GitHub` tomonidan yuborilgan `JSON` ma'lumotlarini qayta ishlang va loglarga kiriting.
    
3. **Ma'lumotlarni qayta ishlash:**
    - Voqealar tafsilotlarini ma'lumotlar bazasiga (masalan, PostgreSQL yoki MongoDB) saqlash.