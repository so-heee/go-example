USE sample;

INSERT INTO `users` (`id`, `name`, `created_at`, `updated_at`)
    VALUES (1, '太郎', '2022-01-01 01:01:01', '2022-01-01 01:01:01'), (2, '次郎', '2022-02-02 02:02:02', '2022-02-02 02:02:02'), (3, '三郎', '2022-02-03 03:03:03', '2022-02-03 03:03:03');

INSERT INTO `tweets` (`id`, `text`, `user_id`, `created_at`, `updated_at`)
    VALUES (1, '私は太郎です', 1, '2022-01-01 01:01:01', '2022-01-01 01:01:01'), (2, '私も太郎です', 1, '2022-01-01 01:01:02', '2022-01-01 01:01:02'), (3, '私は次郎です', 2, '2022-02-02 02:02:02', '2022-02-02 02:02:02');
