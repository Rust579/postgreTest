select first_name||'_', age * 2 twice_age,
 case WHEN age > 20 THEN 'взрослый' ELSE 'ребенок' end age_cat 
 from persons
--where person.first_name = 'Rustem' and age > 25
order by age asc;

select * from persons; -- Вытащить все строки и столбцы

select first_name, age from persons; -- Вытащить все из столбцов first_name, age

select * from persons 
where age > 30
order by age desc; -- Выбрать всех с age > 30 c сортировкой по age по убыванию

select * from persons 
where age > 30 and first_name != 'Роман'

select *, first_name||' '||last_name FIO,
 case WHEN age  <= 33 THEN 'Родился в РФ' ELSE 'Родился в СССР' end born_in
 from persons;

select max(age) max_age, min(age) min_age, sum(age) sum, count(age) count_age  
from persons;

select first_name, max(birth_date) max_age, min(birth_date) min_age, count(birth_date) count_age
from persons
group by first_name;

select distinct first_name from persons; -- уникальность

select count(distinct first_name) from persons;

select distinct first_name, last_name from persons;

select * from
(select *, first_name||' '||last_name FIO,
 case WHEN age  <= 33 THEN 'Родился в РФ' ELSE 'Родился в СССР' end born_in
 from persons) borns
where born_in = 'Родился в СССР'; -- Обращение к виртуальной таблице borns

select s.stud_number, p.first_name ||' '|| p.last_name FIO, p.passport  
from students s 
join persons p on s.id_person = p.id_person;

select s.stud_number, p.first_name ||' '|| p.last_name FIO, p.passport  
from students s, persons p 
WHERE s.id_person = p.id_person;

select s.stud_number, p.first_name ||' '|| p.last_name FIO, p.passport, g.group_number 
from students s 
join persons p on s.id_person = p.id_person
join groups g on s.id_group  = g.id_group;

-- Добавить в группы и в мувементс номер курса!
-- Создать таблицу учебный план
-- Написать запрос который вытащит перечень групп с количеством студентов в каждой и курса
-- Вернуть группы в мувементс!
-- В адреса добавить тип адреса (место жительства и регистрации)!
-- В персонс добавить поле гражданство!
-- Иностранные таблицы удалить!