CREATE TABLE stats (id integer primary key autoincrement, created_at datetime, updated_at datetime, deleted_at datetime, ip_address nvarchar(30), page nvarchar(50));
CREATE TABLE pages (id integer primary key autoincrement, created_at datetime, updated_at datetime, deleted_at datetime, title nvarchar(30), controller nvarchar(25), url nvarchar(25), prev_id integer, next_id integer);

Insert into pages(title, controller, url, prev_id, next_id)
Values ('Termites', 'projects', 'termites', null, 2),
('Particles', 'projects', 'particles', 1, 3),
('Markov Chains', 'projects', 'markovchains', 2, 4),
('Seir Model - Color', 'projects', 'seircolor', 3, null);