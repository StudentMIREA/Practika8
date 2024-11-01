Практическая работа 8.
=================================
### ФИО: Рябова Екатерина Сергеевна | Группа: ЭФБО-01-22
### Дисциплина: Программирование корпоративных систем

***

#### 1. Создание сервера.
На языке go пишем программу которая будет обрабатывать запросы.

![Снимок экрана 2024-10-28 192357](https://github.com/user-attachments/assets/de9e8b97-4466-49ab-bd19-c88ad31f2286)

После этого в командной строке запускаем программу

![Снимок экрана 2024-10-28 192455](https://github.com/user-attachments/assets/b4afac2e-953c-4861-8eec-5e2674c28c57)

![Снимок экрана 2024-10-29 091627](https://github.com/user-attachments/assets/64f36373-c78d-4a1a-8112-29b000ba3e75)

#### 2. Создание главной страницы с товарами. 
На главной странице находятся карточки товаров. Внизу находится панель навигации. При нажатии на карточку товара осуществляется переход на страницу товара. 

При нажатии на сердечко товар отправляется в избранное. Если нажать на активированное сердечко, то товар удалится из избранного. 

При нажатии на кнопку "В корзину", товар отправляется в корзину и появляет возможность выбрать количество товаров. При нулевом значении количества товаров, товар удаляется из корзины.

В правом верхнем углу располагается кнопка, при нажатии на которую осуществляется переход на страницу создания нового товара.

<img src="https://github.com/user-attachments/assets/a40e9555-237a-4c5b-86e9-0a5ab92dd33e" width="300">

#### 3. Создание страницы с товаром.
При нажатии на товар осуществляется переход на страницу товара. Можно вернуться на главную страницу, нажав стрелочку в левом верхнем углу. 

При нажатии на иконку сердечка под товаром, товар добавляется в избранное или удаляется из избранного.

При нажатии на иконку корзины под товаром, товар добавляется в корзину и появляет возможность выбрать количество товаров. При нулевом значении количества товаров, товар удаляется из корзины. Также появляется кнопка, которая позволяет перейти в корзину с товарами. Если пролистать страницу вниз, то можно обнаружить кнопку для удаления товара из каталога. После удаления появляется соответствующее сообщение.

<img src="https://github.com/user-attachments/assets/af107373-4482-4012-bfbb-5dcdd08250bd" width="300">
<img src="https://github.com/user-attachments/assets/4f33fb5f-dc80-4af2-85cb-d33a2c1af569" width="300">

#### 4. Создание страницы добавления товара.
На данной странице вводится информация о товаре. Если есть незаполненные поля, сохранить товар будет не получится. При вводе URL картинки справа от поля ввода появится выбранная картинка при корректном вводе.

<img src="https://github.com/user-attachments/assets/8f9ea032-7813-438c-9fb5-7fecb375d0a2" width="300">
<img src="https://github.com/user-attachments/assets/0b06f6a0-66b5-499b-9590-22adb100219a" width="300">

#### 5. Создание страницы избранных товаров.
На этой странице находятся карточки товаров, которые были добавлены пользователем в избранное. Внизу находится панель навигации. При нажатии на карточку товара осуществляется переход на страницу товара. 

При нажатии на сердечко товар удаляется из избранного. Если избранных товаров нет, появляется соответствующая надпись.

При нажатии на кнопку "В корзину", товар отправляется в корзину и появляет возможность выбрать количество товаров. При нулевом значении количества товаров, товар удаляется из корзины.

<img src="https://github.com/user-attachments/assets/9c90f371-a91e-40b4-8cb8-9162bf0a46bf" width="300">
<img src="https://github.com/user-attachments/assets/65bfd629-3659-46a7-84e3-e5e0da4b6e8f" width="300">

#### 6. Создание страницы корзины.

На этой странице находятся карточки товаров, которые были добавлены пользователем в корзину. Внизу находится панель навигации. При нажатии на карточку товара осуществляется переход на страницу товара. 

У каждого товара можно выбрать количество товаров. Для удаления необходимо свайпнуть товар влево.

<img src="https://github.com/user-attachments/assets/c2720e00-4ddf-4d1c-ac91-23a9e078de27" width="300">
<img src="https://github.com/user-attachments/assets/6c5be591-d1b0-4eb8-a8d0-b2dfed21cd8d" width="300">

#### 7. Создание страницы профиля.
На этой странице находится информация о пользователе. Внизу находится панель навигации. При нажатии на иконку редактирования осуществляется переход на страницу редактирования профиля.

<img src="https://github.com/user-attachments/assets/532cb022-979b-434d-ac9a-40ada74f0a38" width="300">

#### 8. Создание страницы редактирования профиля.
На данной странице вводится информация о пользователе. Если есть незаполненные поля, сохранить информацию не получится. При вводе URL картинки сверху от поля ввода появится выбранная картинка при корректном вводе.

<img src="https://github.com/user-attachments/assets/8e130a09-060a-4f62-945c-249d3c942bb7" width="300">

#### Конечный результат

https://github.com/user-attachments/assets/db780198-5c73-4c44-94ab-201fc17b18c8
