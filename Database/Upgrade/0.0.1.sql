CREATE TABLE stats (id integer primary key autoincrement,
created_at datetime,
updated_at datetime,
deleted_at datetime,
ip_address nvarchar(30),
page nvarchar(50));

CREATE TABLE pages (id integer primary key autoincrement,
created_at datetime,
updated_at datetime,
deleted_at datetime,
title nvarchar(30),
description nvarchar(70),
controller nvarchar(20),
url nvarchar(30),
prev_id integer,
next_id integer);

Insert into pages(title, description, controller, url, prev_id, next_id)
Values ('Termites', 'Cellular automata painting a grid', 'projects', 'termites', null, 2),
('Particles', 'Unity particle systems drawing in 3D space','projects', 'particles', 1, 3),
('Markov Chains', 'Probablistic model for picking colors','projects', 'markovchains', 2, 4),
('Seir Model - Color', 'SEIR epidemiology model generating colors', 'projects', 'seircolor', 3, null);