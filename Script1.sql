select first_name||'_', age * 2 twice_age,
 case WHEN age > 20 THEN 'взрослый' ELSE 'ребенок' end age_cat 
 from person 
--where person.first_name = 'Rustem' and age > 25
order by age asc 

select * from person; -- Вытащить все строки и столбцы

select first_name, age from person; -- Вытащить все из столбцов first_name, age

select * from person 
where age > 30
order by age desc; -- Выбрать всех с age > 30 c сортировкой по age по убыванию

select * from person 
where age > 30 and first_name != 'Роман'

select *, first_name||' '||last_name FIO,
 case WHEN age  <= 33 THEN 'Родился в РФ' ELSE 'Родился в СССР' end born_in
 from person;

select max(age) max_age, min(age) min_age, sum(age) sum, count(age) count_age  
from person;

select first_name, max(age) max_age, min(age) min_age, sum(age) sum, count(age) count_age
from person
group by first_name;

select distinct first_name from person; -- уникальность

select count(distinct first_name) from person;

select distinct first_name, last_name from person;

select * from
(select *, first_name||' '||last_name FIO,
 case WHEN age  <= 33 THEN 'Родился в РФ' ELSE 'Родился в СССР' end born_in
 from person) borns
where born_in = 'Родился в СССР'; -- Обращение к виртуальной таблице borns
 

-- Проверить структуру БД на применимость в жизни. 
-- Скорректировать структуру БД так, чтобы можно было вести учет для студентов, которые учатся на двух направлениях одновременно
-- Повторить скрипты SQL
-- Добавить строк в таблицы (по 10 шт)