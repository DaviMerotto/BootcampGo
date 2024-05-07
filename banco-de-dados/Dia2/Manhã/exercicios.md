### Parte 1

### Exercicio 1

- é uma forma de relacionar duas tabelas e gerar uma tabela resultante com as relações explícitas entre elas

### Exercicio 2

- Inner join: captura os dados que estão presentes na intersecção das duas tabelas, onde o campo usado para linkar os dados, está presente de ambos os lados
- Left join: captura os dados que estão presentes na intersecção das duas tabelas, mas também pega todos os dados da tabela a esquerda que não estão na intersecção.

### Exercicio 3

-  combina registros com valores idênticos da lista de campos especificada em um único registro

### Exercicio 4

-  usada especificamente para filtrar registros resultantes de uma consulta agregada

### Exercicio 5

- O diagrama a esquerda representa o inner join, onde os dados resultantes estão linkados entre as duas tabelas
- O diagrama a direita representa o left join, onde os dados resultantes estão linkados entre as duas tabelas, mas também temos os dados da tabela a esquerda que não estão linkados

### Exercicio 6

- RIGHT JOIN
```SQL
  SELECT m.title, g.name FROM movies m
  RIGHT JOIN genres g ON m.genre_id = g.id
```

- FULL JOIN
```SQL
  SELECT m.title, g.name FROM movies m
  LEFT OUTER JOIN genres g ON g.id = m.genre_id
  UNION
  SELECT m.title, g.name FROM movies m
  RIGHT OUTER JOIN genres g ON g.id = m.genre_id;
```

### Parte 2

### Exercicio 1

```SQL
  SELECT s.title, g.name FROM series s
  LEFT JOIN genres g ON s.genre_id = g.id;
```
![ex1Day2Morning](https://github.com/DaviMerotto/BootcampGo/assets/48769725/a4a523b7-a83f-43f5-9cfa-1bfa97e300fc)

### Exercicio 2

```SQL
  SELECT e.title, a.first_name, a.last_name FROM episodes e
  LEFT JOIN actor_episode ae ON ae.episode_id = e.id
  LEFT JOIN actors a ON ae.actor_id = a.id;
```
![ex2Day2Morning](https://github.com/DaviMerotto/BootcampGo/assets/48769725/ae8c9af1-42ca-4fd0-a415-e40c001f441a)

### Exercicio 3

```SQL
  SELECT e.title, a.first_name, a.last_name FROM episodes e
  LEFT JOIN actor_episode ae ON ae.episode_id = e.id
  LEFT JOIN actors a ON ae.actor_id = a.id;
```
![ex3Day2Morning](https://github.com/DaviMerotto/BootcampGo/assets/48769725/61eb9f17-df91-4a82-95fa-1e53cbdc49e0)

### Exercicio 4

```SQL
  SELECT g.name, COUNT(m.id) 'movies' FROM genres g
  LEFT JOIN movies m ON m.genre_id = g.id
  GROUP BY g.name;
```
![ex4Day2Morning](https://github.com/DaviMerotto/BootcampGo/assets/48769725/a07584fc-d411-493a-ab2b-f3864e7ccf24)

### Exercicio 5

```SQL
  SELECT DISTINCT a.first_name, a.last_name FROM actors a
  LEFT JOIN actor_movie am ON am.actor_id = a.id
  LEFT JOIN movies m ON m.id = am.movie_id
  WHERE m.title like 'Star Wars%'
```
![ex5Day2Morning](https://github.com/DaviMerotto/BootcampGo/assets/48769725/5a0d01ef-f372-4d14-b0d2-2970b278e13c)

