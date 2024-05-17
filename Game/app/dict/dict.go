package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["ShowSport"] = "<b>Вид спорта:</b> %s"
	dict["writedate"] = "Введите дату проведения игры в формате ДДММГГГГ используя любой разделитель"
	dict["first"] = "Просмотр расписания"
	dict["second"] = "Регистрация на игру"
	dict["third"] = "Наши фото и видео"
	dict["fourth"] = "Настройки | Мои игры"
	dict["MainMenu"] = "Главное Меню"
	dict["ShowDate"] = "<b>Дата:</b> %s"
	dict["writetime"] = "Введите время проведения игры в формате ЧЧММ используя любой разделитель"
	dict["ShowTime"] = "<b>Время:</b> %s"
	dict["writeseats"] = "Введите количество свободных мест на эту игру"
	dict["ShowSeats"] = "<b>Всего свободных мест:</b> %d"
	dict["writeprice"] = "Введите цену за одно место в формате цифры на эту игру"
	dict["writecurrency"] = "Введите имя валюты. Я не никак не контралирую то название, которое вы введете, так что советую вводить так, чтобы все понимали. Пример: USD EURO TL и тд"
	dict["ShowTotalprice"] = "<b>Цена на одно место:</b> %d %s"
	dict["writelink"] = `Пршлите ссылку с Google Maps с тем местом, где будет проходить игра. Очень важно, чтоб в ссылке были координаты. 
	Если у вас не будет получатся коректная ссылка, то могу предложить скопировать пример ссылки и вписать на место пропусков координаты вручную. 
	https://www.google.com/maps?q=<i>Тут место для широты</i>,<i>Тут место для долготы</i>`
	dict["ShowLink"] = "<b>Ссылка на место проведения:</b> https://www.google.com/maps?q=%s,%s"
	dict["writeaddress"] = "Введите название адреса"
	dict["ShowAddress"] = "<b>Название адреса:</b> %s"
	dict["clarification"] = "Вы закончили заполнять информацию для создания новой игры. Сохраните эту игру, если все данные верны"
	dict["save"] = "Сохранить"
	dict["change"] = "Изменить"
	dict["gameWasSaved"] = "Игра сохранена и теперь доступна вашим клиентам для регистрации"
	dict["chooseChangable"] = "Выберите что хотите изменить"
	dict["sport"] = "Спорт"
	dict["date"] = "Дата"
	dict["time"] = "Время"
	dict["seats"] = "Места"
	dict["price"] = "Цена (цифра)"
	dict["currency"] = "Валюта"
	dict["link"] = "Ссылка"
	dict["nameaddress"] = "Название адреса"
	dict["volleyball"] = "Волейбол"
	dict["football"] = "Футбол"
	dict["writesport"] = "Выберите вид спорта"
	dict["dataWasChanged"] = "Данные изменены. Желаете изменить что то еще?"
	dict["yes"] = "Да"
	dict["no"] = "Нет"
	dict["gameWasDeleted"] = "Игра удалена. Желаете сделать что-то еще?"
	dict["createGame"] = "Создать игру"
	dict["changeGame"] = "Изменить игру"
	dict["deleteGame"] = "Удалить игру"
	dict["gameDirections"] = "Выберите направление связаное с играми"
	dict["chooseGame"] = "Выберите игру"
}

func en(dict map[string]string) {
	dict["ShowSport"] = "<b>Sport type:</b> %s"
	dict["writedate"] = "Enter the game date in the format DDMMYYYY using any separator"
	dict["first"] = "View schedule"
	dict["second"] = "Game registration"
	dict["third"] = "Our photos and videos"
	dict["fourth"] = "Settings | My games"
	dict["MainMenu"] = "Main Menu"
	dict["ShowDate"] = "<b>Date:</b> %s"
	dict["writetime"] = "Enter the game time in the format HHMM using any separator"
	dict["ShowTime"] = "<b>Time:</b> %s"
	dict["writeseats"] = "Enter the number of available seats for this game"
	dict["ShowSeats"] = "<b>Total available seats:</b> %d"
	dict["writeprice"] = "Enter the price for one seat in numeric format for this game"
	dict["writecurrency"] = "Enter the currency name. I have no control over the name you enter, so I advise entering it in a way that everyone understands. Example: USD, EURO, TL, etc"
	dict["ShowTotalprice"] = "<b>Price for one seat:</b> %d %s"
	dict["writelink"] = `Send a link from Google Maps with the location of the game. It is very important that the link contains coordinates. 
	If you cannot generate a correct link, I can suggest copying the example link and manually entering the coordinates in the placeholders. 
	https://www.google.com/maps?q=<i>Here goes latitude</i>,<i>Here goes longitude</i>`
	dict["ShowLink"] = "<b>Location link:</b> https://www.google.com/maps?q=%s,%s"
	dict["writeaddress"] = "Enter the name of the address"
	dict["ShowAddress"] = "<b>Address name:</b> %s"
	dict["clarification"] = "You have finished filling out the information for creating a new game. Save this game if all the data is correct"
	dict["save"] = "Save"
	dict["change"] = "Modify"
	dict["gameWasSaved"] = "The game has been saved and is now available for registration by your clients"
	dict["chooseChangable"] = "Choose what you want to modify"
	dict["sport"] = "Sport"
	dict["date"] = "Date"
	dict["time"] = "Time"
	dict["seats"] = "Seats"
	dict["price"] = "Price (number)"
	dict["currency"] = "Currency"
	dict["link"] = "Link"
	dict["nameaddress"] = "Address Name"
	dict["volleyball"] = "Volleyball"
	dict["football"] = "Football"
	dict["writesport"] = "Choose a sport"
	dict["dataWasChanged"] = "The data has been modified. Do you want to change anything else?"
	dict["yes"] = "Yes"
	dict["no"] = "No"
	dict["gameWasDeleted"] = "The game has been deleted. Do you want to do anything else?"
	dict["createGame"] = "Create game"
	dict["changeGame"] = "Modify game"
	dict["deleteGame"] = "Delete game"
	dict["gameDirections"] = "Choose the option related to games"
	dict["chooseGame"] = "Choose a game"
}

func tur(dict map[string]string) {
	dict["ShowSport"] = "<b>Spor türü:</b> %s"
	dict["writedate"] = "Herhangi bir ayırıcı kullanarak oyun tarihini DDMMYYYY formatında girin"
	dict["first"] = "Programı görüntüle"
	dict["second"] = "Oyun kaydı"
	dict["third"] = "Bizim fotoğraflar ve videolar"
	dict["fourth"] = "Ayarlar | Oyunlarım"
	dict["MainMenu"] = "Ana Menü"
	dict["ShowDate"] = "<b>Tarih:</b> %s"
	dict["writetime"] = "Herhangi bir ayırıcı kullanarak oyun saati HHMM formatında girin"
	dict["ShowTime"] = "<b>Saat:</b> %s"
	dict["writeseats"] = "Bu oyun için mevcut koltuk sayısını girin"
	dict["ShowSeats"] = "<b>Toplam boş koltuk sayısı:</b> %d"
	dict["writeprice"] = "Bu oyun için bir koltuk için fiyatı sayısal formatta girin"
	dict["writecurrency"] = "Para birimi adını girin. Girdiğiniz isim üzerinde kontrolüm yok, bu yüzden herkesin anlayabileceği bir şekilde girmenizi öneririm. Örnek: USD, EURO, TL, vb"
	dict["ShowTotalprice"] = "<b>Bir koltuk için fiyat:</b> %d %s"
	dict["writelink"] = `Oyunun gerçekleşeceği yerin Google Haritalar'dan bir bağlantısını gönderin. Bağlantının koordinatlar içermesi çok önemlidir. 
	Eğer doğru bir bağlantı oluşturamazsanız, örnek bağlantıyı kopyalayıp yer tutuculara koordinatları el ile girmeyi önerebilirim. 
	https://www.google.com/maps?q=<i>Buraya enlem</i>,<i>Buraya boylam</i>`
	dict["ShowLink"] = "<b>Yer bağlantısı:</b> https://www.google.com/maps?q=%s,%s"
	dict["writeaddress"] = "Adresin adını girin"
	dict["ShowAddress"] = "<b>Adres adı:</b> %s"
	dict["clarification"] = "Yeni bir oyun oluşturmak için bilgileri doldurmayı tamamladınız. Eğer tüm veriler doğru ise, bu oyunu kaydedin"
	dict["save"] = "Kaydet"
	dict["change"] = "Değiştir"
	dict["gameWasSaved"] = "Oyun kaydedildi ve şimdi müşterileriniz tarafından kayıt için kullanılabilir"
	dict["chooseChangable"] = "Değiştirmek istediğiniz şeyi seçin"
	dict["sport"] = "Spor"
	dict["date"] = "Tarih"
	dict["time"] = "Saat"
	dict["seats"] = "Koltuklar"
	dict["price"] = "Fiyat (sayı)"
	dict["currency"] = "Para Birimi"
	dict["link"] = "Bağlantı"
	dict["nameaddress"] = "Adres Adı"
	dict["volleyball"] = "Voleybol"
	dict["football"] = "Futbol"
	dict["writesport"] = "Bir spor seçin"
	dict["dataWasChanged"] = "Veriler değiştirildi. Başka bir şeyi değiştirmek ister misiniz?"
	dict["yes"] = "Evet"
	dict["no"] = "Hayır"
	dict["gameWasDeleted"] = "Oyun silindi. Başka bir şey yapmak ister misiniz?"
	dict["createGame"] = "Oyun oluştur"
	dict["changeGame"] = "Oyunu değiştir"
	dict["deleteGame"] = "Oyunu sil"
	dict["gameDirections"] = "Oyunlarla ilgili bir seçenek seçin"
	dict["chooseGame"] = "Bir oyun seçin"
}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)
	Dictionary["tur"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
	tur(Dictionary["tur"])
}
