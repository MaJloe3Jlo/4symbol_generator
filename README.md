`EN`**Key generator API v.0.0.1 by m3**

App to give an 4-symbol key.
Statuses of key: "not issued", "issued" anf "off".

Requires:
1) Redis on port :6379 with no password.
2) Any browser.
3) curl or Postman for POST method.

Server starts on :7100 port and if keys do not exist in Redis - waiting for input in console.

How to use:
1) GET methods:
    * path "/" - is a main information of app.
    * path "/key" - is gave you 4-symbol key with status "not issued" and ID of it.
    * path "/key/{id}" - an information of key by ID(you get by "key" path).
    * path "/statistics" - is given information about keys. How many: "issued", "not issued" or "off"
2) POST method:
    * path "/key/{id}" - is gave status "off" to key, only if it's had status "issued".
    
`RU`**Генератор ключей v.0.0.1 от m3**

Приложение выдает 4символьные ключи. Статусы ключей: "не выдан", "выдан", "погашен".

Требования:
1) Redis на порте :6379 без пароля.
2) Любой браузер.
3) curl или Postman для POST метода.

Приложение стартует на :7100 порте и если ключи не существуют в Redis - ожидает ввода в консоли.

Как использовать:
1) GET методы:
    * "/" - основная информация о приложении.
    * "/key" - выдает 4символьный ключ у которого статус "не выдан" и его ID.
    * "/key/{id}" - инфомация о ключе по его ID (полученного в методе "key").
    * "/statistics" - общая информация о ключах. Сколько из них имеют статусы: "не выдан", "выдан" и "погашен".
2) POST методы:
    * "/key/{id}" - перевод ключ в статус "погашен, только если ключ был в статусе "выдан".            
