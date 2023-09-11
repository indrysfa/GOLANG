select concat(first_name, ' ', last_name) as nama_lengkap
from sakila.customer c
where active = 1
order by first_name asc 

-- huruf depan U
select * from sakila.film f where left(title, 1) = 'U'

-- select * from sakila.film f where like title = 'U%'

-- Tampilkan kota yang berada di negara India. Urutkan secara ascending.
select * from sakila.city c where country_id = 44 
-- order by asc 

-- Tampilkan customer_id dan nilai rata-rata amount dari tabel payment dengan kondisi customer yang memiliki rata-rata amount lebih besar dari 4.
-- select customer_id avg(amount) from sakila.payment p
-- group by customer_id 
-- having avg(amount) > 4 

-- Hapus data pada tabel rental yang dikelola oleh staff dengan id 1
delete from sakila.rental where staff_id = 1

-- 6.Ubah nama depan staff menjadi "Mikhael" untuk staff dengan id = 1
UPDATE sakila.staff 
   SET first_name  = 'Mikhael'
 WHERE staff_id = 1

 -- 7.Tambahkan kategori film dengan nama kategori "Romance"
INSERT INTO sakila.category(name)
VALUES ('Romance')

-- 8.Tampilkan nama lengkap customer dan alamatnya, urutkan berdasarkan nama (Clue : menggunakan JOIN)
SELECT concat(first_name, ' ', last_name) nama_lengkap, address
  FROM sakila.address c 
  JOIN sakila.customer a ON c.address_id = a.address_id
  
-- 9.Tampilkan semua film yang memiliki rental rate lebih tinggi dari rental rate rata-rata semua film (clue : menggunakan subquery)
SELECT title, rental_rate 
from sakila.film
where rental_rate > (
select avg(rental_rate)
from sakila.film f)
  
-- 10.Tampilkan nama lengkap semua aktor dalam film ANGELS LIFE. (clue : JOIN)
select DISTINCT concat(first_name, ' ', last_name) as actor from sakila.actor a
join sakila.film_actor fa on a.actor_id = fa.actor_id

-- select concat(first_name, ' ', last_name) as nama_lengkap
-- from actor as a 
-- join sakila.film_actor as fa on a.actor_id fa.actor_id 
-- join sakila.film as f on fa.film_id = f.film_id 
-- where f.title  = 'ANGELS LIFE'


-- 11.Tampilkan nama kota dan jumlah customer dalam kota tersebut.
-- select city (select count(*) from sakila.address a where a.city_id = c.city_id) as jumlah_pelanggan
-- from city c

-- 12.Lakukan penambahan durasi untuk film yang memiliki kategori Action.
select * from sakila.film_category fc
join sakila.film f on c.category_id = fc.category_id 
where fc.

-- 13.Tampilkan judul film yang dipinjam oleh customer bernama GERTRUDE
select film.title
from film
join inventory on film.film_id = inventory.film_id 
join customer on rental.customer_id = customer.customer_id
where customer

-- 14.

-- 15. Lakukan update biaya rental menjadi 1.2 kali dari harga aslinya untuk film dengan kategori Documentary.
update sakila.film
set rental_rate = rental_rate * 1.2
where film_id in (select film_id 
from sakila.film_category 
where category_id = 6)

-- INSERT INTO sakila.film(title, description , release_year, language_id, original_language_id, rental_duration, rental_rate, `length`, replacement_cost, rating, special_features)
-- VALUES ('Title film', 'desc 123', '2007', 1, null, 6, 4.99, 34, 'PG', 'Trailers')
-- where 











