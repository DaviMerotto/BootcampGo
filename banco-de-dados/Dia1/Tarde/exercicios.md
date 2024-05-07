### Exercicio 2

<img width="1023" alt="ex1day1tarde" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/49eb7e9b-b9b2-404d-b73e-2652168afca8">

### Exercicio 3

<img width="1023" alt="ex2Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/aaffef9b-3478-4c57-be1e-732c973fa75d">

### Exercicio 4

<img width="1023" alt="ex4Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/7a660c75-8c33-4f37-a111-284c8fd2cd6d">

### Exercicio 5 

<img width="1023" alt="ex5Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/8590b30e-9743-40ee-8bd3-3f96e358dd83">

### Exercicio 6

<img width="1023" alt="ex6Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/bb2bb59d-4cac-468e-8d2d-78c52120065d">

### Exercicio 7

<img width="1023" alt="ex7Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/de7a499a-3a6b-4fc0-94a7-749a7f6bf61e">

### Exercicio 8

<img width="1023" alt="ex8Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/dd29d46c-8ee2-4d88-89df-c36607cdc46e">

### Exercicio 9

<img width="1023" alt="ex9Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/b972ee3a-6668-4b32-b1b6-afa5ab022ba3">

### Exercicio 10

<img width="1023" alt="ex10Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/a115423d-3bd3-4a5d-b0a3-c2e46cefb0f7">

### Exercicio 11

<img width="1023" alt="ex11Day1Afternoon" src="https://github.com/DaviMerotto/BootcampGo/assets/48769725/207dcb9b-7d7f-4cfe-8e16-bd6970bb2fed">

### Exercicio 12

![ex12Day1Afternoon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/2f9c8fad-3b48-4fb7-a483-a5676b2e3dea)

### Exercicio 13

![ex13Day1Afternoon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/222f1e63-86c7-44ff-96dc-22a5cf0bbe37)

### Exercicio 14

![ex14Day1Afternoon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/05b1c1fe-b018-42a8-8d9c-43ddbeb430f3)

### Exercicio 15

![ex15Day1Afternoon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/381d8575-f714-4463-b17b-ac762cf27f5d)

### Exercicio 16

![ex16Day1Afternoon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/7fe0acb9-b4bd-48bc-a909-9fb487b78dc9)

### Exercicio 17

![ex17Day1Afterroon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/ef978ef3-6826-4903-b762-9e5418055d40)

### Exercicio 18

![ex18Day1Afternoon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/45fb8cc2-b7cf-4338-87b1-0c9f2d7be915)

### Codigo SQL

```SQL
SELECT 
  a.first_name,
  a.last_name 
FROM actors a
WHERE a.rating > 7.5;

SELECT 
  m.title,
  m.rating,
  m.awards
FROM movies m
WHERE m.awards > 2 AND m.rating > 7.5;

SELECT
  m.title,
  m.rating	
FROM movies m
ORDER BY m.rating;

SELECT * FROM movies
ORDER BY rating desc
LIMIT 5
OFFSET 5;

SELECT * FROM actors
LIMIT 10;

SELECT * FROM actors
LIMIT 10
OFFSET 20;

SELECT * FROM actors
LIMIT 10
OFFSET 40;

SELECT m.title, m.rating FROM movies m
WHERE m.title like 'Toy Story%';

SELECT * FROM actors a
WHERE a.first_name like 'Sam';

SELECT m.title FROM movies m
WHERE m.release_date BETWEEN '2004-01-01' AND '2008-12-31';

SELECT m.title FROM movies m
WHERE m.rating > 3
  AND m.awards > 1 
  AND m.release_date BETWEEN '1988-01-01' AND '2009-12-31'
  LIMIT 3
  OFFSET 10;
```
