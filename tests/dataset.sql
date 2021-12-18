INSERT INTO users (username, password, public_key, private_key, quota) VALUES 
('robin', 'toto', 'pub'::bytea, 'priv'::bytea, 128500),
('este', 'toto', 'pub'::bytea, 'priv'::bytea, 0);

INSERT INTO files (id, name, size, chunk, f_owner, last_edit) VALUES
('58b58f40-0503-4244-b22c-6d7b0715687c', 'toto.txt', 128000, 4096, 'robin', NOW() AT TIME ZONE 'utc'),
('aebbcb2f-f9c2-41a3-82a3-eb9c508ceb8a', 'ficher.test', 500, 100, 'robin', NOW() AT TIME ZONE 'utc');

INSERT INTO file_access (id, access_state, f_shared_by, f_shared_to, f_file, encryption_key) VALUES
('6af874fe-4fc2-4d6b-848d-d74c1c010960', 'PRIVATE', 'robin', NULL, '58b58f40-0503-4244-b22c-6d7b0715687c', 'enc'::bytea),
('ae994406-69da-468e-805c-11917a8e9162', 'PRIVATE', 'robin', NULL, 'aebbcb2f-f9c2-41a3-82a3-eb9c508ceb8a', 'enc'::bytea),
('48a2a7e6-da47-4328-9a7f-fb3e923f1d4d', 'SHARED', 'robin', 'este', 'aebbcb2f-f9c2-41a3-82a3-eb9c508ceb8a', 'enc'::bytea);

-- SELECT ALL FILE THAT BELONG TO THE USER (to get shard files, replace "PRIVATE" by "SHARED")
SELECT
	t1.id AS access_id,
	t1.access_state,
  t1.f_shared_to AS shared_to,
  t1.f_shared_by AS shared_by,
	t2.id AS file_id,
	t2.name AS file_name,
  t2.size AS file_size,
  t2.last_edit,
  t1.favorite,
	t2.encryption_key,
FROM file_access AS t1
  INNER JOIN files AS t2 ON t1.f_file = t2.id
WHERE t1.f_shared_by = 'robin' AND t1.access_state = 'PRIVATE';



SELECT t1.id AS access_id, t1.access_state, t1.f_shared_to AS shared_to, t1.f_shared_by AS shared_by, t2.id AS file_id, t2.name AS file_name, t2.size AS file_size, t2.last_edit, t1.favorite, t1.encryption_key FROM file_access AS t1 INNER JOIN files AS t2 ON t1.f_file = t2.id WHERE t1.f_shared_by = 'monoko' AND t1.access_state = 'PRIVATE';