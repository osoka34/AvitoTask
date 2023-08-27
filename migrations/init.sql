CREATE TABLE public.user
(
    user_id int NOT NULL PRIMARY KEY,
    user_name VARCHAR(100) NOT NULL
);
CREATE TABLE public.segment
(
    segment_id SERIAL PRIMARY KEY,
    segment_name VARCHAR(100) NOT NULL
);

CREATE TABLE public.users_in_segm
(
    user_id int ,
    segment_id int ,
    ttl DATE,
    CONSTRAINT id PRIMARY KEY (user_id, segment_id),
    FOREIGN KEY (user_id)  REFERENCES public.user(user_id),
    FOREIGN KEY (segment_id)  REFERENCES public.segment(segment_id)
);

-- Вставка значений в таблицу public.user
INSERT INTO public.user (user_id, user_name)
VALUES
    (1, 'User A'),
    (2, 'User B'),
    (3, 'User C'),
    (4, 'User D'),
    (5, 'User E'),
    (6, 'User F'),
    (7, 'User G');


INSERT INTO public.segment  (segment_name)
VALUES
    ( 'Segment X'),
    ( 'Segment Y'),
    ( 'Segment Z');

-- -- Вставляем только тех пользователей, которые не принадлежат данному сегменту
-- INSERT INTO public.users_in_segm (user_id, segment_id)
-- SELECT u.user_id, 101 -- Нужный сегмент
-- FROM public.user u
-- LEFT JOIN public.users_in_segm us ON u.user_id = us.user_id
-- WHERE random() < 0.3 -- 30% пользователей
--   AND us.user_id IS NULL;

-- UPDATE
--     public.segment
-- SET
--     segment_name = 'Segment F'
-- WHERE
--     segment_id = 100;

--
-- INSERT INTO public.segment(segment_id, segment_name)
-- VALUES (105, 'Segment T')
-- 			RETURNING
-- 				segment_id;
--
-- DELETE FROM public.segment
-- WHERE segment_id = 104;
--
-- DELETE FROM public.segment
-- WHERE segment_name = 'Segment T';

-- SELECT segment_id, segment_name FROM public.segment
-- WHERE segment_name = 'Segment F';
--
-- UPDATE
--     public.segment
-- SET
--     segment_name = 'seg F'
-- WHERE
--         segment_id = 1;
--
--
-- UPDATE
--     public.segment
-- SET
--     segment_name = 'seg F'
-- WHERE
--         segment_id = 1000;
--
-- DELETE FROM public.segment
-- WHERE segment_name = 'seg F';
--
-- SELECT ROW_COUNT();
--
-- SELECT COALESCE(count(*),0) AS count FROM public.segment WHERE 'segment_name'='Segment F';