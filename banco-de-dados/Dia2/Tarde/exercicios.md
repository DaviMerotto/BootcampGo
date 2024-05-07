### Exercicio 1

-  A normalização é um processo de padronização e validação de dados que consiste em eliminar redundâncias e inconsistências, completando os dados por meio de uma série de regras que atualizam as informações, protegendo sua integridade e favorecendo a interpretação, para que seja mais fácil consultar e mais eficiente para quem a administra.

### Exercicio 2

```SQL
  INSERT INTO movies
    (created_at,
     updated_at,
     title,
     rating,
     awards,
     release_date,
     length)
  VALUES 
    ('2024-05-07',
      null,
      'Star Wars Episodio IV - Uma nova esperança',
      8.6,
      6,
      '1977-11-18',
      121
    );
```

### Exercicio 3

```SQL
INSERT INTO genres
  (created_at,
  updated_at,
  name,
  ranking,
  active)
VALUES
  (
    null,
    null,
    'Ópera espacial',
    13,
    1
  )
```

### Exercicio 4

```SQL
  UPDATE movies m
  SET m.genre_id = 13
  WHERE m.id = 22
```

### Exercicio 5

```SQL
  UPDATE actors a
  SET a.favorite_movie_id = 22
  WHERE a.id = 47
```

### Exercicio 6

```SQL
 CREATE TEMPORARY TABLE temp_movies SELECT * FROM movies
```

### Exercicio 7

```SQL
  DELETE FROM temp_movies tm WHERE tm.awards < 5
```

### Exercicio 8

```SQL
  SELECT g.id, g.created_at, g.updated_at, g.name, g.ranking, g.active 
  FROM genres g
  INNER JOIN movies m ON m.genre_id = g.id
```

### Exercicio 9

```SQL
  SELECT a.id,a.first_name,a.last_name,a.rating,m.title FROM actors a
  INNER JOIN movies m ON a.favorite_movie_id = m.id
  WHERE m.awards > 3
```

### Exercicio 11

- Indices são mecanismos para acelerar consultas SQL, eles fornecem melhores caminhos para chegar aos dados, evitando realizar consultas completas ou lineares de dados de uma tabela

### Exercicio 12

```SQL
  CREATE INDEX movie_title_index
  ON movies (title);
```

### Exercicio 13

![ex13Day2Afternoon](https://github.com/DaviMerotto/BootcampGo/assets/48769725/14895e8c-7bb2-4071-9e0e-5a1d6b8d44f3)
